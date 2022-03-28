package biz

import (
	"container/list"
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/conf"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"gitee.com/moyusir/util/parser"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	// BatchGetDeviceStateInfo 批量查询某一类设备的状态信息
	BatchGetDeviceStateInfo(deviceClassID int, option *QueryOption) ([]*v1.DeviceState, error)
	// BatchGetDeviceWarningDetectField 批量查询某一类设备某个字段的信息，用于预警检测
	BatchGetDeviceWarningDetectField(deviceClassID int, fieldName string, option *QueryOption) (*api.QueryTableResult, error)
	// GetWarningMessage 查询当前存储在数据库中的警告信息
	GetWarningMessage(option *QueryOption) ([]*utilApi.Warning, error)
	// SaveWarningMessage 保存警告信息
	SaveWarningMessage(bucket string, warnings ...*utilApi.Warning) error
	// RunWarningDetectTask 依据预警字段注册的预警规则，创建并运行下采样设备状态信息数据的task
	RunWarningDetectTask(config *WarningDetectTaskConfig) (*domain.Run, error)
	// StopWarningDetectTask 关闭指定task的运行
	StopWarningDetectTask(run *domain.Run) error
}

// WarningDetectTaskConfig 创建负责下采样设备状态信息的task的配置信息
type WarningDetectTaskConfig struct {
	// task的唯一标识名
	Name string
	// 设备的类别号
	DeviceClassID int
	// 设备的字段名，通过设备号和字段名确定需要进行下采样的设备字段
	FieldName string
	// 下采样的时间窗口大小
	Every time.Duration
	// 数据聚合类型
	AggregateType utilApi.DeviceStateRegisterInfo_AggregationOperation
	// 下采样的数据需要写入的目标bucket
	TargetBucket string
}

// QueryOption 查询influxdb时可以附加的参数
type QueryOption struct {
	// 查询的桶
	Bucket string
	// 查询的时间范围，绝对的时间值
	Start, Stop *time.Time
	// 利用相对于now()的相对时间值进行查询，比如查询过去5m到现在的所有时间序列数据，
	// 当绝对的时间查询和相对的时间查询参数都存在时优先使用相对的时间查询
	Past time.Duration
	// 查询时的过滤条件，可以指定tag、_measurement、_field
	Filter map[string]string
	// 依据指定的列名进行groupBy操作，用于reshape查询结果
	GroupBy []string
}

// 链表节点，保存mutex、channel以及状态标志
type warningPushNode struct {
	mutex          *sync.Mutex
	isActive       bool
	warningChannel chan *utilApi.Warning
}

func NewWarningDetectUsecase(repo UnionRepo, registerInfo []utilApi.DeviceStateRegisterInfo, logger log.Logger) (*WarningDetectUsecase, error) {
	ctx, cancel := context.WithCancel(context.Background())

	infoParser, err := parser.NewRegisterInfoParser(registerInfo)
	if err != nil {
		return nil, err
	}

	w := &WarningDetectUsecase{
		repo:               repo,
		logger:             log.NewHelper(logger),
		parser:             infoParser,
		warningDetectGroup: new(errgroup.Group),
		warningPushGroup:   new(errgroup.Group),
		// TODO 考虑容量
		warningChannel:        make(chan *utilApi.Warning, 10),
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

	// 返回关闭预警预测的清理函数
	return w, nil
}

// 负责对某个预警字段进行预警检测,start用于标志成功启动
func (u *WarningDetectUsecase) warningDetect(deviceClassID int, field *parser.WarningDetectField, start chan<- struct{}) error {
	// 启动协程相应的下采样task
	// 启动下采样的task
	// 每个用户拥有三个桶:<username>、<username-warnings>、<username-warning_detect>
	// 分别保存用户设备状态信息、警告信息、下采样的设备状态信息
	every := field.Rule.Duration.AsDuration()
	taskConf := &WarningDetectTaskConfig{
		Name:          fmt.Sprintf("%s-%d-%s-task", conf.Username, deviceClassID, field.Name),
		DeviceClassID: deviceClassID,
		FieldName:     field.Name,
		Every:         every,
		AggregateType: field.Rule.AggregationOperation,
		TargetBucket:  fmt.Sprintf("%s-warning_detect", conf.Username),
	}
	run, err := u.repo.RunWarningDetectTask(taskConf)
	if err != nil {
		return err
	}
	defer u.repo.StopWarningDetectTask(run)

	// 发出任务启动成功的信号
	start <- struct{}{}

	// 控制查询间隔的定时器
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	// 避免访问nil map而实例化的空map
	m := make(map[string]string)
	// 查询用的option
	option := &QueryOption{
		Bucket: taskConf.TargetBucket,
		Past:   every,
		Filter: m,
	}
	for {
		select {
		case <-u.ctx.Done():
			return nil
		case <-ticker.C:
			// 调用repo层函数进行查询
			// TODO 考虑错误处理
			tableResult, err := u.repo.BatchGetDeviceWarningDetectField(deviceClassID, field.Name, option)
			if err != nil {
				u.logger.Error(err)
				continue
			}

			// 依据解析注册信息得到的预警规则进行预警检测
			for tableResult.Next() {
				record := tableResult.Record()
				value := record.Value().(float64)
				if field.Func(value) {
					deviceID := record.Measurement()
					u.warningChannel <- &utilApi.Warning{
						DeviceId:        deviceID,
						DeviceClassId:   int32(deviceClassID),
						DeviceFieldName: field.Name,
						WarningMessage:  fmt.Sprintf("%s %s warning", deviceID, field.Name),
						Start:           timestamppb.New(record.ValueByKey("_start").(time.Time)),
						End:             timestamppb.New(record.ValueByKey("_stop").(time.Time)),
					}
				}
			}
			tableResult.Close()
		}
	}
}

// 负责保存警告信息到数据库以及将channel的警告消息扇出到warningFanOutChannels的channel中
// 并负责检测链表的状态，及时将非活跃状态的节点放回池中
func (u *WarningDetectUsecase) warningFanOut() {
	bucket := fmt.Sprintf("%s-warnings", conf.Username)
	for {
		select {
		case <-u.ctx.Done():
			return
		case warning := <-u.warningChannel:
			// 首先保存警告消息
			// TODO 考虑是否需要推送序列化失败或者保存失败的警告消息
			err := u.repo.SaveWarningMessage(bucket, warning)
			if err != nil {
				u.logger.Error(err)
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
func (u *WarningDetectUsecase) StartDetection() error {
	// 依据注册信息，为每一类设备的每一个预警字段，开启预警检测的协程
	// 两个chan用于确定协程正确启动
	startDone := make(chan struct{}, len(u.parser.Info))
	startError := make(chan error, len(u.parser.Info))

	for i := 0; i < len(u.parser.Info); i++ {
		fields := u.parser.GetWarningDetectFields(i)
		for j := 0; j < len(fields); j++ {
			// 由于启动协程和循环的进行速度不一致，因此不能直接以闭包的形式将i与j传入
			// 协程中使用，否则可能会造成数组越界的panic
			deviceClassID, fieldIndex := i, j
			u.warningDetectGroup.Go(func() error {
				err := u.warningDetect(deviceClassID, &(fields[fieldIndex]), startDone)
				if err != nil {
					startError <- err
				}
				return nil
			})
		}
	}

	// 等待预警检测的协程都启动完毕
	for i := 0; i < len(u.parser.Info); i++ {
		select {
		case <-startDone:
		case err := <-startError:
			return err
		}
	}

	// 开启预警扇出的协程
	u.warningDetectGroup.Go(func() error {
		u.warningFanOut()
		return nil
	})

	return nil
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

// BatchGetDeviceStateInfo 批量查询设备状态信息
func (u *WarningDetectUsecase) BatchGetDeviceStateInfo(
	deviceClassID int,
	option *QueryOption) ([]*v1.DeviceState, error) {
	option.Bucket = conf.Username
	return u.repo.BatchGetDeviceStateInfo(deviceClassID, option)
}

// BatchGetWarning 批量查询警告信息
func (u *WarningDetectUsecase) BatchGetWarning(option *QueryOption) ([]*utilApi.Warning, error) {
	// 以<用户名-warning>为名的bucket中保存着警告信息
	option.Bucket = fmt.Sprintf("%s-warnings", conf.Username)
	return u.repo.GetWarningMessage(option)
}

// GetDeviceStateRegisterInfo 查询关于设备状态，包括预警规则在内的注册信息
func (u *WarningDetectUsecase) GetDeviceStateRegisterInfo(deviceClassID int) *utilApi.DeviceStateRegisterInfo {
	return &u.parser.Info[deviceClassID]
}
