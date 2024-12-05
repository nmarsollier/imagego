package image_dao

import "time"

type ImageDao interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
}
