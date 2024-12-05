package db

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/nmarsollier/imagego/tools/env"
)

var (
	redis_once     sync.Once
	redis_instance *RedisDao
)

func getRedisDb() ImageDao {
	redis_once.Do(func() {
		redis_instance = &RedisDao{redis.NewClient(&redis.Options{
			Addr:     env.Get().RedisURL,
			Password: "",
			DB:       0,
		})}
	})
	return redis_instance
}

type RedisDao struct {
	client *redis.Client
}

func (r *RedisDao) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisDao) Set(key string, value string, expiration time.Duration) (string, error) {
	return r.client.Set(key, value, expiration).Result()
}
