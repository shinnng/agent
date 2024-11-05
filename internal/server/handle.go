package server

import (
	"agent/internal/conf"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	stdhttp "net/http"
	"net/http/httputil"
	"net/url"
)

type CachedReverseProxy struct {
	conf  *conf.Data
	rdb   *redis.Client
	ctx   context.Context
	proxy *httputil.ReverseProxy // reverse proxy to DAL(hasura)
}

func (p CachedReverseProxy) PerCheck() {
	// todo: check rdb and proxy
}

func (p CachedReverseProxy) GetCacheKey(url string) string {
	sha2 := sha256.New()
	sha2.Write([]byte(url))
	return hex.EncodeToString(sha2.Sum(nil))
}

func (p *CachedReverseProxy) GetCacheData(req *http.Request) ([]byte, error) {
	key := p.GetCacheKey(req.URL.String())
	data, err := p.rdb.Get(p.ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Debugf("get cache data miss %s", key)
		} else {
			log.Errorf("get cache data error %s", key)
		}
		return nil, err
	}

	log.Debugf("get cache data succeed, len: %d", len(data))
	return data, nil
}

func (p *CachedReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := p.ServeHTTPByCache(rw, req)
	if err != nil {
		p.ServeHTTPByProxy(rw, req)
	}
}

func (p *CachedReverseProxy) ServeHTTPByCache(rw http.ResponseWriter, req *http.Request) error {
	data, err := p.GetCacheData(req)
	if err != nil {
		return err
	}

	// write data to responseWrite
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	//rw.Header().Del("transfer-encoding")
	rw.WriteHeader(stdhttp.StatusOK)

	// batch rw
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

func (p *CachedReverseProxy) ServeHTTPByProxy(rw http.ResponseWriter, req *http.Request) {
	crw := NewCacheResponseWriter(p.GetCacheKey(req.URL.String()), rw, p.rdb)
	p.proxy.ServeHTTP(crw, req)

	crw.PushCache()
}

func NewCachedReverseProxy(cfg *conf.Data) *CachedReverseProxy {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		DB:       int(cfg.Redis.Db),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})
	ctx := context.Background()

	// set proxy property
	target, err := url.Parse(cfg.Hasura.Addr)
	if err != nil {
		log.Fatalf("Error parsing target URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.BufferPool = newBufferPool()

	p := &CachedReverseProxy{
		conf:  cfg,
		rdb:   rdb,
		ctx:   ctx,
		proxy: proxy,
	}
	p.PerCheck()
	return p
}

type bufferPool struct {
	bytes []byte
}

func newBufferPool() *bufferPool {
	return &bufferPool{make([]byte, 102400)}
}
func (bp bufferPool) Get() []byte  { return bp.bytes }
func (bp bufferPool) Put(v []byte) { bp.bytes = v }
