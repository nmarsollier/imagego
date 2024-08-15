package redisx

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/nmarsollier/imagego/tools/env"
)

func Get(ctx ...interface{}) RedisClient {
	for _, o := range ctx {
		if client, ok := o.(RedisClient); ok {
			return client
		}
	}

	return redisClient{
		redis.NewClient(&redis.Options{
			Addr:     env.Get().RedisURL,
			Password: "",
			DB:       0,
		}),
	}
}

type RedisClient interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type redisClient struct {
	client *redis.Client
}

func (c redisClient) Get(key string) *redis.StringCmd {
	return c.client.Get(key)
}

func (c redisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(key, value, expiration)
}
