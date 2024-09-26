package redisx

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/nmarsollier/imagego/tools/env"
)

var (
	once     sync.Once
	instance *redis.Client
)

func Get(ctx ...interface{}) RedisClient {
	for _, o := range ctx {
		if client, ok := o.(RedisClient); ok {
			return client
		}
	}

	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:     env.Get().RedisURL,
			Password: "",
			DB:       0,
		})
	})
	return instance
}

type RedisClient interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
