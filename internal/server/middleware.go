package server

import (
	"agent/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
)

func RouteListMiddleware(apis *conf.Apis) middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			q, ok := req.(*http.Request)
			if !ok {
				return nil, stdhttp.ErrNotSupported
			}

			for _, e := range apis.Metadata.RestEndpoints {
				if q.URL.Path == e.Url {
					return next(ctx, req)
				}
			}

			return nil, stdhttp.ErrNotSupported
		}
	}
}
