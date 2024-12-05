package db

import (
	"time"

	"github.com/nmarsollier/imagego/tools/env"
)

/**
 * @title ImageDao
 * @desc ImageDao is the interface that defines the methods to interact with the image storage
 */
type ImageDao interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) (string, error)
}

func Get(deps ...interface{}) ImageDao {
	for _, o := range deps {
		if client, ok := o.(ImageDao); ok {
			return client
		}
	}

	if env.Get().Source == "redis" {
		return getRedisDb()
	} else {
		return getDynamoDb()
	}
}
