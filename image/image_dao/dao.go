package image_dao

import "time"

/**
 * @title ImageDao
 * @desc ImageDao is the interface that defines the methods to interact with the image storage
 */
type ImageDao interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
}
