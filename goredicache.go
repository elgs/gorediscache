package gorediscache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	redisPool *redis.Pool
	Data      map[string]any
}

func (this *Cache) Set(key string, value any) {
	this.Data[key] = value
}

func (this *Cache) Get(key string) any {
	if val, ok := this.Data[key]; ok {
		return val
	}
	return nil
}

func (this *Cache) Empty() {
	for k := range this.Data {
		delete(this.Data, k)
	}
}

func (this *Cache) Close() error {
	return this.redisPool.Close()
}

func NewCache(redisUrl string) *Cache {

	redisPool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisUrl)
		},
	}

	return &Cache{
		redisPool: redisPool,
	}
}
