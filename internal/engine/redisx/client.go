package redisx

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
}

type redisClient struct {
	client *redis.Client
}

func Get(
	redisUrl string,
) RedisClient {
	return &redisClient{redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})}
}

func (r *redisClient) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *redisClient) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.client.Set(key, value, expiration).Result()
}
