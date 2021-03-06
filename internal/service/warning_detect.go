package service

import (
	"context"
	"gitee.com/moyusir/data-processing/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"net/http"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
)

type WarningDetectService struct {
	pb.UnimplementedWarningDetectServer
	warningDetectUsecase *biz.WarningDetectUsecase
	// 用于将前端发送进行ws连接的http请求转换为ws连接
	upgrader *websocket.Upgrader
	logger   *log.Helper
}

func NewWarningDetectService(wu *biz.WarningDetectUsecase, logger log.Logger) (
	*WarningDetectService, func(), error) {
	// 开启预警检测
	err := wu.StartDetection()
	if err != nil {
		return nil, nil, err
	}

	// 返回关闭预警检测的清理函数
	return &WarningDetectService{
			warningDetectUsecase: wu,
			upgrader: &websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				// 解决跨域问题
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			logger: log.NewHelper(logger),
		}, func() {
			wu.CloseDetection()
		}, nil
}

func (s *WarningDetectService) BatchGetDeviceStateInfo(ctx context.Context, req *pb.BatchGetDeviceStateRequest) (*pb.BatchGetDeviceStateReply, error) {
	option := new(biz.QueryOption)

	if req.Past != nil && req.Past.AsDuration() != 0 {
		option.Past = req.Past.AsDuration()
	} else if req.Start != nil {
		start := req.Start.AsTime()
		option.Start = &start
		if req.End != nil {
			stop := req.End.AsTime()
			option.Stop = &stop
		}
	}

	if req.Filter != nil {
		option.Filter = req.Filter
	} else {
		option.Filter = make(map[string]string)
	}

	option.Limit = int(req.Limit)
	option.Offset = int(req.Offset)

	states, total, err := s.warningDetectUsecase.BatchGetDeviceStateInfo(int(req.DeviceClassId), option)
	if err != nil {
		return nil, err
	}

	return &pb.BatchGetDeviceStateReply{States: states, Total: int32(total)}, nil
}

func (s *WarningDetectService) DeleteDeviceStateInfo(ctx context.Context, request *pb.DeleteDeviceStateRequest) (*pb.DeleteDeviceStateReply, error) {
	err := s.warningDetectUsecase.DeleteDeviceState(request)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteDeviceStateReply{Success: true}, nil
}

func (s *WarningDetectService) BatchGetWarning(ctx context.Context, req *pb.BatchGetWarningRequest) (*pb.BatchGetWarningReply, error) {
	option := new(biz.QueryOption)
	if req.Past != nil && req.Past.AsDuration() != 0 {
		option.Past = req.Past.AsDuration()
	} else if req.Start != nil {
		start := req.Start.AsTime()
		option.Start = &start
		if req.End != nil {
			stop := req.End.AsTime()
			option.Stop = &stop
		}
	}

	if req.Filter != nil {
		option.Filter = req.Filter
	} else {
		option.Filter = make(map[string]string)
	}

	option.Limit = int(req.Limit)
	option.Offset = int(req.Offset)

	warnings, total, err := s.warningDetectUsecase.BatchGetWarning(option)
	if err != nil {
		return nil, err
	}

	return &pb.BatchGetWarningReply{Warnings: warnings, Total: int32(total)}, nil
}

func (s *WarningDetectService) DeleteWarning(ctx context.Context, request *pb.DeleteWarningRequest) (*pb.DeleteWarningReply, error) {
	err := s.warningDetectUsecase.DeleteWarningMessage(request)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteWarningReply{Success: true}, nil
}

func (s *WarningDetectService) UpdateWarning(ctx context.Context, request *pb.UpdateWarningRequest) (*pb.UpdateWarningReply, error) {
	err := s.warningDetectUsecase.UpdateWarningProcessedState(request)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateWarningReply{Success: true}, nil
}

func (s *WarningDetectService) ServeWebsocketConnection(w http.ResponseWriter, r *http.Request) error {
	// 利用前端发送的http请求建立ws连接
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Errorf("建立ws连接时发生了错误:%v", err)
		return err
	}

	// 将建立的连接交给biz层的预警检测用例处理
	s.warningDetectUsecase.AddWarningPushConnection(conn)
	return nil
}
