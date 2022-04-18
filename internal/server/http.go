package server

import (
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/conf"
	"gitee.com/moyusir/data-processing/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	h "net/http"
	"strings"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, cs *service.ConfigService, ws *service.WarningDetectService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(
				recovery.WithLogger(logger),
			),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterWarningDetectHTTPServer(srv, ws)
	v1.RegisterConfigHTTPServer(srv, cs)
	// 额外添加处理ws连接请求的http路由
	srv.HandleFunc(
		"/warnings/push/"+strings.ReplaceAll(conf.Username, "_", "-"),
		func(writer h.ResponseWriter, request *h.Request) {
			ws.ServeWebsocketConnection(writer, request)
		},
	)
	return srv
}
