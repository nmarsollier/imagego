package db

import (
	"github.com/go-redis/redis"
	"github.com/nmarsollier/imagego/tools/env"
	"github.com/nmarsollier/imagego/tools/errors"
)

// ErrData la imagen no parece valida
var ErrNotFound = errors.NewValidationField("id", "invalid")

// Client retorna un cliente a redis
func Client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     env.Get().RedisURL,
		Password: "",
		DB:       0,
	})
}
