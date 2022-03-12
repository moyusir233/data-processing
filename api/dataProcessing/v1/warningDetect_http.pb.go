// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	v1 "gitee.com/moyusir/util/api/util/v1"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WarningDetectHTTPServer interface {
	BatchGetDeviceState(context.Context, *BatchGetDeviceStateRequest) (*BatchGetDeviceStateReply, error)
	BatchGetWarning(context.Context, *BatchGetWarningRequest) (*BatchGetWarningReply, error)
	GetDeviceStateRegisterInfo(context.Context, *GetDeviceStateRegisterInfoRequest) (*v1.DeviceStateRegisterInfo, error)
}

func RegisterWarningDetectHTTPServer(s *http.Server, srv WarningDetectHTTPServer) {
	r := s.Route("/")
	r.GET("/states/{device_class_id}/{page}/{count}", _WarningDetect_BatchGetDeviceState0_HTTP_Handler(srv))
	r.GET("/warnings/{page}/{count}", _WarningDetect_BatchGetWarning0_HTTP_Handler(srv))
	r.GET("/register-info/states/{device_class_id}", _WarningDetect_GetDeviceStateRegisterInfo0_HTTP_Handler(srv))
}

func _WarningDetect_BatchGetDeviceState0_HTTP_Handler(srv WarningDetectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchGetDeviceStateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.dataProcessing.v1.WarningDetect/BatchGetDeviceState")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchGetDeviceState(ctx, req.(*BatchGetDeviceStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchGetDeviceStateReply)
		return ctx.Result(200, reply)
	}
}

func _WarningDetect_BatchGetWarning0_HTTP_Handler(srv WarningDetectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchGetWarningRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.dataProcessing.v1.WarningDetect/BatchGetWarning")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchGetWarning(ctx, req.(*BatchGetWarningRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchGetWarningReply)
		return ctx.Result(200, reply)
	}
}

func _WarningDetect_GetDeviceStateRegisterInfo0_HTTP_Handler(srv WarningDetectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeviceStateRegisterInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.dataProcessing.v1.WarningDetect/GetDeviceStateRegisterInfo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDeviceStateRegisterInfo(ctx, req.(*GetDeviceStateRegisterInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.DeviceStateRegisterInfo)
		return ctx.Result(200, reply)
	}
}

type WarningDetectHTTPClient interface {
	BatchGetDeviceState(ctx context.Context, req *BatchGetDeviceStateRequest, opts ...http.CallOption) (rsp *BatchGetDeviceStateReply, err error)
	BatchGetWarning(ctx context.Context, req *BatchGetWarningRequest, opts ...http.CallOption) (rsp *BatchGetWarningReply, err error)
	GetDeviceStateRegisterInfo(ctx context.Context, req *GetDeviceStateRegisterInfoRequest, opts ...http.CallOption) (rsp *v1.DeviceStateRegisterInfo, err error)
}

type WarningDetectHTTPClientImpl struct {
	cc *http.Client
}

func NewWarningDetectHTTPClient(client *http.Client) WarningDetectHTTPClient {
	return &WarningDetectHTTPClientImpl{client}
}

func (c *WarningDetectHTTPClientImpl) BatchGetDeviceState(ctx context.Context, in *BatchGetDeviceStateRequest, opts ...http.CallOption) (*BatchGetDeviceStateReply, error) {
	var out BatchGetDeviceStateReply
	pattern := "/states/{device_class_id}/{page}/{count}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.dataProcessing.v1.WarningDetect/BatchGetDeviceState"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarningDetectHTTPClientImpl) BatchGetWarning(ctx context.Context, in *BatchGetWarningRequest, opts ...http.CallOption) (*BatchGetWarningReply, error) {
	var out BatchGetWarningReply
	pattern := "/warnings/{page}/{count}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.dataProcessing.v1.WarningDetect/BatchGetWarning"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarningDetectHTTPClientImpl) GetDeviceStateRegisterInfo(ctx context.Context, in *GetDeviceStateRegisterInfoRequest, opts ...http.CallOption) (*v1.DeviceStateRegisterInfo, error) {
	var out v1.DeviceStateRegisterInfo
	pattern := "/register-info/states/{device_class_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.dataProcessing.v1.WarningDetect/GetDeviceStateRegisterInfo"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}