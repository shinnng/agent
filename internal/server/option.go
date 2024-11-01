package server

import (
	"agent/internal/conf"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func CachedProxyOption(c *conf.Data) http.ServerOption {
	return func(s *http.Server) {
		cachedProxy := NewCachedReverseProxy(c)
		s.HandlePrefix("/", cachedProxy)
	}
}
