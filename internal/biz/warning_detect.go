package biz

import (
	"container/list"
	"context"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"gitee.com/moyusir/util/parser"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"sync"
	"time"
)

const WarningDetectFieldLabelName = "field_id"

type WarningDetectUsecase struct {
	repo   UnionRepo
	logger *log.Helper
	// 用于解析注册信息，获得设备预警规则
	parser *parser.RegisterInfoParser
	// 负责预警检测的协程组
	warningDetectGroup *errgroup.Group
	// 管理警告消息推送的协程组
	warningPushGroup *errgroup.Group
	// 存放警告消息的channel
	warningChannel chan *utilApi.Warning
	// 保存所有需要推送警告消息的channel链表，这些channel与负责通过ws连接推送信息的协程
	// 相连，扇出协程负责将warningMessageChannel中的消息推送到这些channel中
	warningFanOutChannels *list.List
	// 存放链表节点的池，用于复用channel
	pool *sync.Pool
	// 多协程控制
	ctx    context.Context
	cancel func()
	mutex  *sync.Mutex
}

type WarningDetectRepo interface {
	// BatchGetDeviceStateInfo 批量查询设备状态信息，返回保存在数据库中设备状态的二进制信息
	BatchGetDeviceStateInfo(key string, option *QueryOption) ([][]byte, error)
	// BatchGetDeviceWarningDetectField 批量查询具有指定label的TS中保存的预警字段值信息
	BatchGetDeviceWarningDetectField(label string, option *TSQueryOption) ([]TsQueryResult, error)
	// GetWarningMessage 查询当前存储在数据库中的警告信息
	GetWarningMessage(key string, option *QueryOption) ([][]byte, error)
	// SaveWarningMessage 保存警告信息
	SaveWarningMessage(key string, timestamp int64, value []byte) error
}

// QueryOption 查询ZSet时可以附加的查询参数
type QueryOption struct {
	// 查询的时间范围，以时间戳表示
	Begin, End int64
	// 偏移量和记录数限制
	Offset, Count int64
}

// TSQueryOption 批量查询TS时可以附加的查询参数
type TSQueryOption struct {
	// 查询的时间范围，以时间戳表示
	Begin, End int64
	// 查询的聚合类型
	AggregationType string
	// 聚合查询时使用的时间间隔，以毫秒为单位
	TimeBucket time.Duration
}

// TsQueryResult Ts查询结果的封装
type TsQueryResult struct {
	Key       string
	Value     float64
	Timestamp int64
}

// 链表节点，保存mutex、channel以及状态标志
type warningPushNode struct {
	mutex          *sync.Mutex
	isActive       bool
	warningChannel chan *utilApi.Warning
}

func NewWarningDetectUsecase(repo UnionRepo, registerInfo []utilApi.DeviceStateRegisterInfo, logger log.Logger) (*WarningDetectUsecase, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())
	infoParser, err := parser.NewRegisterInfoParser(registerInfo)
	if err != nil {
		return nil, nil, err
	}

	w := &WarningDetectUsecase{
		repo:                  repo,
		logger:                log.NewHelper(logger),
		parser:                infoParser,
		warningDetectGroup:    new(errgroup.Group),
		warningPushGroup:      new(errgroup.Group),
		warningChannel:        make(chan *utilApi.Warning),
		warningFanOutChannels: list.New(),
		pool: &sync.Pool{New: func() interface{} {
			return &warningPushNode{
				mutex:    new(sync.Mutex),
				isActive: false,
				// TODO 考虑chan的容量
				warningChannel: make(chan *utilApi.Warning, 10),
			}
		}},
		ctx:    ctx,
		cancel: cancel,
		mutex:  new(sync.Mutex),
	}

	w.StartDetection()

	// 返回关闭预警预测的清理函数
	return w, func() { w.CloseDetection() }, nil
}

// 负责对某个预警字段进行预警检测
func (u *WarningDetectUsecase) warningDetect(deviceClassID int, field *parser.WarningDetectField) {
	// 获得预警字段对应的标签
	label := GetDeviceStateFieldLabel(deviceClassID, field.Name)

	// 配置查询的默认选项
	option := &TSQueryOption{
		Begin:           0,
		End:             0,
		AggregationType: strings.ToLower(field.Rule.AggregationOperation.String()),
		TimeBucket:      field.Rule.Duration.AsDuration(),
	}

	// 控制查询间隔的定时器
	ticker := time.NewTicker(option.TimeBucket)
	defer ticker.Stop()

	for {
		select {
		case <-u.ctx.Done():
			return
		case <-ticker.C:
			// 配置批量查询TS的参数
			now := time.Now()
			begin := now.Add(-1 * option.TimeBucket)
			end := now
			option.Begin = begin.Unix()
			option.End = end.Unix()

			// 调用repo层函数进行查询
			// TODO 考虑错误处理
			detectFields, err := u.repo.BatchGetDeviceWarningDetectField(label, option)
			if err != nil {
				continue
			}

			// 依据解析注册信息得到的预警规则进行预警检测
			for _, f := range detectFields {
				if !field.Func(f.Value) {
					// 将警告消息发送到存放警告信息的channel中
					u.warningChannel <- &utilApi.Warning{
						DeviceId:        GetDeviceIDFromDeviceStateFieldKey(f.Key),
						DeviceFieldName: field.Name,
						WarningRule:     field.Rule,
						Start:           timestamppb.New(begin),
						End:             timestamppb.New(end),
					}
				}
			}
		}
	}
}

// 负责保存警告信息到数据库以及将channel的警告消息扇出到warningFanOutChannels的channel中
// 并负责检测链表的状态，及时将非活跃状态的节点放回池中
func (u *WarningDetectUsecase) warningFanOut() {
	// 获得警告消息保存的redis key
	warningsSaveKey := GetWarningsSaveKey()

	for {
		select {
		case <-u.ctx.Done():
			return
		case warning := <-u.warningChannel:
			// 首先保存警告消息
			// TODO 考虑是否需要推送序列化失败或者保存失败的警告消息
			marshal, err := proto.Marshal(warning)
			if err == nil {
				u.repo.SaveWarningMessage(warningsSaveKey, time.Now().Unix(), marshal)
			}

			// 利用了信号量，避免同时进行链表添加节点和检索链表的行为
			ele := u.warningFanOutChannels.Front()
			for ele != nil {
				next := ele.Next()
				u.mutex.Lock()

				// 互斥访问链表节点
				node := ele.Value.(*warningPushNode)
				node.mutex.Lock()

				// 当节点仍活跃，说明对应的前端连接未断开，应该推送警告消息
				if node.isActive {
					node.warningChannel <- warning
				} else {
					// 若节点不活跃，则将其从链表中删除，并放回池中
					u.warningFanOutChannels.Remove(ele)
					u.pool.Put(node)
				}
				node.mutex.Unlock()

				u.mutex.Unlock()
				ele = next
			}
		}
	}
}

// 负责向前端推送预警信息
func (u *WarningDetectUsecase) warningPush(conn *websocket.Conn, node *warningPushNode) {
	// 协程结束后关闭与前端的连接，并标志node为不活跃状态
	defer func() {
		conn.Close()
		node.mutex.Lock()
		node.isActive = false
		node.mutex.Unlock()
	}()

	var warning *utilApi.Warning
	for {
		select {
		case <-u.ctx.Done():
			return
		case warning = <-node.warningChannel:
			err := conn.WriteJSON(warning)
			// TODO 考虑错误处理
			if err != nil {
				return
			}
		}
	}
}

// StartDetection 开启预警检测
func (u *WarningDetectUsecase) StartDetection() {
	// 依据注册信息，为每一类设备的每一个预警字段，开启预警检测的协程
	for i := 0; i < len(u.parser.Info); i++ {
		fields := u.parser.GetWarningDetectFields(i)
		for j := 0; j < len(fields); j++ {
			u.warningDetectGroup.Go(func() error {
				// 由于构造预警检测的label时不需要使用
				u.warningDetect(i, &(fields[j]))
				return nil
			})
		}
	}
}

// CloseDetection 关闭预警检测
func (u *WarningDetectUsecase) CloseDetection() error {
	// 利用context结束所有预警检测的协程，并利用errgroup进行等待
	u.cancel()
	if err := u.warningDetectGroup.Wait(); err != nil {
		return err
	}
	return u.warningPushGroup.Wait()
}

// AddWarningPushConnection 添加需要进行前端推送的连接
func (u *WarningDetectUsecase) AddWarningPushConnection(conn *websocket.Conn) {
	// 从池中获得节点，初始化节点后将其连接到链表上
	node := u.pool.Get().(*warningPushNode)
	node.isActive = true

	u.mutex.Lock()
	u.warningFanOutChannels.PushFront(node)
	u.mutex.Unlock()

	u.warningPushGroup.Go(func() error {
		u.warningPush(conn, node)
		return nil
	})
}

// BatchGetDeviceStateInfo 批量查询设备状态信息，并依据传入的proto message模板对二进制信息进行反序列化
// 注意函数执行后stateTemplate中会填充有查询得到的数据
func (u *WarningDetectUsecase) BatchGetDeviceStateInfo(
	deviceClassID int,
	option *QueryOption,
	stateTemplate proto.Message) ([]proto.Message, error) {
	// 以<用户id>:device_state:<设备类别号>为键，在zset中保存
	// 以timestamp为score，以设备状态二进制protobuf信息为value的键值对

	// 获得redis key
	key := GetDeviceStateKey(deviceClassID)

	// 调用repo层函数，查询设备状态信息
	stateInfo, err := u.repo.BatchGetDeviceStateInfo(key, option)
	if err != nil {
		return nil, err
	}

	states := make([]proto.Message, 0, len(stateInfo))
	for _, info := range stateInfo {
		err := proto.Unmarshal(info, stateTemplate)
		if err != nil {
			return nil, err
		}
		states = append(states, proto.Clone(stateTemplate))
	}

	return states, nil
}

// BatchGetWarning 批量查询警告信息
func (u *WarningDetectUsecase) BatchGetWarning(option *QueryOption) ([]*utilApi.Warning, error) {
	// 用户的警告消息存放在以<用户id>:warning为key的ZSet中，
	// 并以timestamp为field，以警告消息的序列化二进制数据为value存储警告消息

	// 调用repo层函数进行查询
	warnings, err := u.repo.GetWarningMessage(GetWarningsSaveKey(), option)
	if err != nil {
		return nil, err
	}

	// 将二进制数据反序列化
	result := make([]*utilApi.Warning, len(warnings))
	for i, warning := range warnings {
		result[i] = new(utilApi.Warning)
		err := proto.Unmarshal(warning, result[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetDeviceStateRegisterInfo 查询关于设备状态，包括预警规则在内的注册信息
func (u *WarningDetectUsecase) GetDeviceStateRegisterInfo(deviceClassID int) *utilApi.DeviceStateRegisterInfo {
	return &u.parser.Info[deviceClassID]
}
