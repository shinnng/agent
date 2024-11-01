package server

import (
	"agent/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Config) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	// set agent options
	if c.Agent.Network != "" {
		opts = append(opts, http.Network(c.Agent.Network))
	}
	if c.Agent.Addr != "" {
		opts = append(opts, http.Address(c.Agent.Addr))
	}
	if c.Agent.Timeout != nil {
		opts = append(opts, http.Timeout(c.Agent.Timeout.AsDuration()))
	}
	if c.Data != nil {
		opts = append(opts, CachedProxyOption(c.Data))
	}

	srv := http.NewServer(opts...)
	// TODO: add path white list
	circuitbreaker.Client()
	return srv
}
