package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"

	"net/http"
)

type CacheResponseWriter struct {
	key string
	rw  http.ResponseWriter

	buf []byte
	ctx context.Context
	rdb *redis.Client
}

func NewCacheResponseWriter(key string, rw http.ResponseWriter, rdb *redis.Client) *CacheResponseWriter {
	return &CacheResponseWriter{
		key: key,
		rw:  rw,
		buf: make([]byte, 0, 5120),
		ctx: context.Background(),
		rdb: rdb,
	}
}

func (w *CacheResponseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *CacheResponseWriter) WriteHeader(statusCode int) {
	w.rw.WriteHeader(statusCode)
}

func (w *CacheResponseWriter) PushCache() {
	go func(data []byte) {
		if err := w.rdb.Set(w.ctx, w.key, data, 10*time.Second).Err(); err != nil {
			log.Debugf("writing to redis error: %v", err)
		}

		w.buf = make([]byte, 0, 5120)
		log.Debugf("writing to redis succeed, len: %d", len(data))
	}(w.buf)
}

func (w *CacheResponseWriter) Write(p []byte) (n int, err error) {
	n, err = w.rw.Write(p)
	if err != nil {
		return n, err
	}
	w.buf = append(w.buf, p...)

	return n, nil
}
