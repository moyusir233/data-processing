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

func NewWarningDetectUsecase(repo UnionRepo, registerInfo []utilApi.DeviceStateRegisterInfo, logger log.Logger) *WarningDetectUsecase {
	ctx, cancel := context.WithCancel(context.Background())
	return &WarningDetectUsecase{
		repo:                  repo,
		logger:                log.NewHelper(logger),
		parser:                parser.NewRegisterInfoParser(registerInfo),
		warningDetectGroup:    new(errgroup.Group),
		warningPushGroup:      new(errgroup.Group),
		warningChannel:        make(chan *utilApi.Warning),
		warningFanOutChannels: list.New(),
		ctx:                   ctx,
		cancel:                cancel,
		mutex:                 new(sync.Mutex),
	}
}

// 负责对某个预警字段进行预警检测
func (u *WarningDetectUsecase) warningDetect(info *DeviceGeneralInfo, field *parser.WarningDetectField) {
	// 获得预警字段对应的标签
	_, label := GetDeviceStateFieldKeyAndLabel(info, field.Name)

	// 配置查询的默认选项
	option := &TSQueryOption{
		Begin:           0,
		End:             0,
		AggregationType: field.Aggregation,
		TimeBucket:      field.Interval,
	}

	// 控制查询间隔的定时器
	ticker := time.NewTicker(field.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-u.ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			option.Begin = now.Unix()
			option.End = now.Add(field.Interval).Unix()

			// TODO 考虑错误处理
			_, err := u.repo.BatchGetDeviceWarningDetectField(label, option)
			if err != nil {
				continue
			}
		}
	}
}

// 负责保存警告信息到数据库以及将channel的警告消息扇出到warningFanOutChannels的channel中
func (u *WarningDetectUsecase) warningFanOut() {

}

// 负责向前端推送预警信息
func (u *WarningDetectUsecase) warningPush(conn *websocket.Conn, warningChannel chan *utilApi.Warning) {

}

// BatchGetDeviceStateInfo 批量查询设备状态信息，并依据传入的proto message模板对二进制信息进行反序列化
// 注意函数执行后stateTemplate中会填充有查询得到的数据
func (u *WarningDetectUsecase) BatchGetDeviceStateInfo(
	info *DeviceGeneralInfo,
	option *QueryOption,
	stateTemplate proto.Message) ([]proto.Message, error) {
	// 以<用户id>:device_state:<设备类别号>为键，在zset中保存
	// 以timestamp为score，以设备状态二进制protobuf信息为value的键值对

	// 获得redis key
	key := GetDeviceStateKey(info)

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

// GetDeviceStateRegisterInfo 查询关于设备状态，包括预警规则在内的注册信息
func (u *WarningDetectUsecase) GetDeviceStateRegisterInfo(deviceClassID int) *utilApi.DeviceStateRegisterInfo {
	return &u.parser.Info[deviceClassID]
}
