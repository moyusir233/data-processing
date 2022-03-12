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

func NewWarningDetectService(wu *biz.WarningDetectUsecase, logger log.Logger) *WarningDetectService {
	return &WarningDetectService{
		warningDetectUsecase: wu,
		upgrader:             &websocket.Upgrader{},
		logger:               log.NewHelper(logger),
	}
}

func (s *WarningDetectService) BatchGetDeviceState(ctx context.Context, req *pb.BatchGetDeviceStateRequest) (*pb.BatchGetDeviceStateReply, error) {
	// 构造查询选项
	option := &biz.QueryOption{
		Begin:  req.Start.AsTime().Unix(),
		End:    req.End.AsTime().Unix(),
		Offset: (req.Page - 1) * req.Count,
		Count:  req.Count,
	}

	// 调用biz层的查询函数
	states, err := s.warningDetectUsecase.BatchGetDeviceStateInfo(int(req.DeviceClassId), option, new(utilApi.TestedDeviceState))
	if err != nil {
		return nil, err
	}

	// 将接口类型向下转型
	reply := new(pb.BatchGetDeviceStateReply)
	for _, s := range states {
		reply.States = append(reply.States, s.(*utilApi.TestedDeviceState))
	}

	return reply, nil
}

func (s *WarningDetectService) BatchGetWarning(ctx context.Context, req *pb.BatchGetWarningRequest) (*pb.BatchGetWarningReply, error) {
	// 构造查询选项
	option := &biz.QueryOption{
		Begin:  req.Start.AsTime().Unix(),
		End:    req.End.AsTime().Unix(),
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
