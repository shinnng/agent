package server

import (
	"agent/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Config, apis *conf.Apis) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Agent.Network != "" {
		opts = append(opts, http.Network(c.Agent.Network))
	}
	if c.Agent.Addr != "" {
		opts = append(opts, http.Address(c.Agent.Addr))
	}
	if c.Agent.Timeout != nil {
		opts = append(opts, http.Timeout(c.Agent.Timeout.AsDuration()))
	}
	// set proxy option
	if c.Data != nil {
		opts = append(opts, CachedProxyOption(c.Data))
	}
	// set route white list
	if apis != nil {
		opts = append(opts, http.Middleware(RouteListMiddleware(apis)))
	}

	srv := http.NewServer(opts...)
	circuitbreaker.Client()
	return srv
}
