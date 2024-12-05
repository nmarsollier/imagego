package redisx

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/nmarsollier/imagego/image/image_dao"
	"github.com/nmarsollier/imagego/tools/env"
)

var (
	once     sync.Once
	instance *RedisDao
)

func Get(deps ...interface{}) image_dao.ImageDao {
	for _, o := range deps {
		if client, ok := o.(image_dao.ImageDao); ok {
			return client
		}
	}

	once.Do(func() {
		instance = &RedisDao{redis.NewClient(&redis.Options{
			Addr:     env.Get().RedisURL,
			Password: "",
			DB:       0,
		})}
	})
	return instance
}

type RedisDao struct {
	client *redis.Client
}

func (r *RedisDao) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisDao) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.client.Set(key, value, expiration).Result()
}
