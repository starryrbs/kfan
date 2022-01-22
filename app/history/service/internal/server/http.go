package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/starryrbs/kfan/api/history/service/v1"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	"github.com/starryrbs/kfan/app/history/service/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, history *service.HistoryService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			tracing.Server(),
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
	v1.RegisterHistoryHTTPServer(srv, history)
	return srv
}
