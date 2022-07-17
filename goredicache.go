package gorediscache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	redisPool  *redis.Pool
	DefaultTTL time.Duration
}

func (this *Cache) SetString(key string, value any, ttl time.Duration) error {
	conn := this.redisPool.Get()
	defer conn.Close()
	if ttl <= 0 {
		ttl = this.DefaultTTL
	}
	_, err := conn.Do("SET", key, value, "EX", ttl.Seconds())
	return err
}

func (this *Cache) GetString(key string) (string, error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	// TODO: increase ttl if found
	reply, err := conn.Do("GET", key)
	if err != nil || reply == nil {
		return "", err
	}
	return redis.String(reply, err)
}

func (this *Cache) SetMap(key string, value map[string]any, ttl time.Duration) error {
	conn := this.redisPool.Get()
	defer conn.Close()
	if ttl <= 0 {
		ttl = this.DefaultTTL
	}
	a := []any{key}
	for k, v := range value {
		a = append(a, k, v)
	}
	conn.Send("HMSET", a...)
	conn.Send("EXPIRE", key, ttl.Seconds())
	return conn.Flush()
}

func (this *Cache) GetMap(key string) (map[string]string, error) {
	conn := this.redisPool.Get()
	defer conn.Close()
	// TODO: increase ttl if found
	reply, err := conn.Do("HGETALL", key)
	if err != nil || reply == nil {
		return nil, err
	}
	return redis.StringMap(reply, err)
}

func (this *Cache) Delete(keys ...any) error {
	conn := this.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", keys...)
	return err
}

func (this *Cache) Close() error {
	return this.redisPool.Close()
}

func NewCache(redisUrl string, defaultTTL time.Duration) *Cache {

	redisPool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisUrl)
		},
	}

	if defaultTTL <= 0 {
		defaultTTL = 3600 * 24 * time.Second // 24 hours
	}

	return &Cache{
		redisPool:  redisPool,
		DefaultTTL: defaultTTL,
	}
}
