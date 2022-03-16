package service

import (
	"context"
	"gitee.com/moyusir/data-processing/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"net/http"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	utilApi "gitee.com/moyusir/util/api/util/v1"
)

type WarningDetectService struct {
	pb.UnimplementedWarningDetectServer
	warningDetectUsecase *biz.WarningDetectUsecase
	// 用于将前端发送进行ws连接的http请求转换为ws连接
	upgrader *websocket.Upgrader
	logger   *log.Helper
}

func NewWarningDetectService(wu *biz.WarningDetectUsecase, logger log.Logger) (
	*WarningDetectService, func()) {
	// 开启预警检测
	wu.StartDetection()

	// 返回关闭预警检测的清理函数
	return &WarningDetectService{
			warningDetectUsecase: wu,
			upgrader:             &websocket.Upgrader{},
			logger:               log.NewHelper(logger),
		}, func() {
			wu.CloseDetection()
		}
}

func (s *WarningDetectService) BatchGetDeviceState0(ctx context.Context, req *pb.BatchGetDeviceStateRequest) (*pb.BatchGetDeviceStateReply0, error) {
	// 代码注入deviceClassID
	deviceClassID := 0

	// 构造查询选项
	var begin, end int64
	if req.Start != nil {
		begin = req.Start.AsTime().Unix()
	} else {
		begin = 0
	}
	if req.End != nil {
		end = req.End.AsTime().Unix()
	} else {
		end = 0
	}

	option := &biz.QueryOption{
		Begin:  begin,
		End:    end,
		Offset: (req.Page - 1) * req.Count,
		Count:  req.Count,
	}

	// 调用biz层的查询函数
	states, err := s.warningDetectUsecase.BatchGetDeviceStateInfo(deviceClassID, option, new(pb.DeviceState0))
	if err != nil {
		return nil, err
	}

	reply := new(pb.BatchGetDeviceStateReply0)
	for _, s := range states {
		// 将接口类型向下转型
		reply.States = append(reply.States, s.(*pb.DeviceState0))
	}

	return reply, nil
}

func (s *WarningDetectService) BatchGetDeviceState1(ctx context.Context, req *pb.BatchGetDeviceStateRequest) (*pb.BatchGetDeviceStateReply1, error) {
	// 代码注入deviceClassID
	deviceClassID := 1

	// 构造查询选项
	var begin, end int64
	if req.Start != nil {
		begin = req.Start.AsTime().Unix()
	} else {
		begin = 0
	}
	if req.End != nil {
		end = req.End.AsTime().Unix()
	} else {
		end = 0
	}

	option := &biz.QueryOption{
		Begin:  begin,
		End:    end,
		Offset: (req.Page - 1) * req.Count,
		Count:  req.Count,
	}

	// 调用biz层的查询函数
	states, err := s.warningDetectUsecase.BatchGetDeviceStateInfo(deviceClassID, option, new(pb.DeviceState1))
	if err != nil {
		return nil, err
	}

	reply := new(pb.BatchGetDeviceStateReply1)
	for _, s := range states {
		// 将接口类型向下转型
		reply.States = append(reply.States, s.(*pb.DeviceState1))
	}

	return reply, nil
}

func (s *WarningDetectService) BatchGetWarning(ctx context.Context, req *pb.BatchGetWarningRequest) (*pb.BatchGetWarningReply, error) {
	// 构造查询选项
	var begin, end int64
	if req.Start != nil {
		begin = req.Start.AsTime().Unix()
	} else {
		begin = 0
	}
	if req.End != nil {
		end = req.End.AsTime().Unix()
	} else {
		end = 0
	}

	option := &biz.QueryOption{
		Begin:  begin,
		End:    end,
		Offset: (req.Page - 1) * req.Count,
		Count:  req.Count,
	}

	// 调用biz层的查询函数
	warnings, err := s.warningDetectUsecase.BatchGetWarning(option)
	if err != nil {
		return nil, err
	}

	return &pb.BatchGetWarningReply{Warnings: warnings}, nil
}

func (s *WarningDetectService) GetDeviceStateRegisterInfo(ctx context.Context, req *pb.GetDeviceStateRegisterInfoRequest) (*utilApi.DeviceStateRegisterInfo, error) {
	registerInfo := s.warningDetectUsecase.GetDeviceStateRegisterInfo(int(req.DeviceClassId))
	return registerInfo, nil
}

func (s *WarningDetectService) ServeWebsocketConnection(w http.ResponseWriter, r *http.Request) error {
	// 利用前端发送的http请求建立ws连接
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	// 将建立的连接交给biz层的预警检测用例处理
	s.warningDetectUsecase.AddWarningPushConnection(conn)
	return nil
}
