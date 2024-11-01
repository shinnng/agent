package server

import (
	"agent/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	stdhttp "net/http"
	"net/http/httputil"
	"net/url"
)

type CachedReverseProxy struct {
	conf                    *conf.Data
	rdb                     *redis.Client
	readCtx, writeCtx       context.Context        // rdb read/write context
	readCancel, writeCancel context.CancelFunc     // rdb read/write cancel
	proxy                   *httputil.ReverseProxy // reverse proxy to DAL(hasura)
}

func (c CachedReverseProxy) PerCheck() {
	// todo: check rdb and proxy
}

func (c CachedReverseProxy) GetCacheKey(req *http.Request) string {
	req.URL.Query().Get("key")
	return ""
}

func (c *CachedReverseProxy) GetCacheData(req *http.Request) ([]byte, error) {
	key := c.GetCacheKey(req)
	return c.rdb.Get(c.readCtx, key).Bytes()
}

func (c *CachedReverseProxy) SetCacheData(rw http.ResponseWriter, req *http.Request) error {
	key := c.GetCacheKey(req)
	data := make([]byte, 0)
	err := c.rdb.Set(c.writeCtx, key, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c CachedReverseProxy) WriteResponseFromCache(rw http.ResponseWriter, data []byte) error {
	// todo: should to use http/Response.Write?
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(stdhttp.StatusOK)

	// batch write
	for begin := 0; begin < len(data); {
		size, err := rw.Write(data[begin:])
		if err != nil {
			log.Error("Error writing response:", err)
			return err
		}
		begin += size
	}

	return nil
}

func (c *CachedReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := c.ServeHTTPByCache(rw, req)
	if err != nil {
		log.Info("serve http by cache failed", "err", err)

		c.ServeHTTPByProxy(rw, req)
	}
}

func (c *CachedReverseProxy) ServeHTTPByCache(rw http.ResponseWriter, req *http.Request) error {
	data, err := c.GetCacheData(req)
	if err != nil {
		return err
	}
	// build response and return
	if err := c.WriteResponseFromCache(rw, data); err != nil {
		return err
	}
	return nil
}

func (c *CachedReverseProxy) ServeHTTPByProxy(rw http.ResponseWriter, req *http.Request) {
	c.proxy.ServeHTTP(rw, req)
	// todo: write to cache
}

func NewCachedReverseProxy(cfg *conf.Data) *CachedReverseProxy {
	// set redis property
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		DB:       int(cfg.Redis.Db),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})
	ctx := context.Background()
	readCtx, readCancel := context.WithTimeout(ctx, cfg.Redis.ReadTimeout.AsDuration())
	writeCtx, writeCancel := context.WithTimeout(ctx, cfg.Redis.WriteTimeout.AsDuration())

	// set proxy property
	target, err := url.Parse(cfg.Hasura.Addr)
	if err != nil {
		log.Fatalf("Error parsing target URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	cp := &CachedReverseProxy{
		conf:        cfg,
		rdb:         rdb,
		readCtx:     readCtx,
		writeCtx:    writeCtx,
		readCancel:  readCancel,
		writeCancel: writeCancel,
		proxy:       proxy,
	}
	cp.PerCheck()
	return cp
}
