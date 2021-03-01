package image

import (
	"github.com/go-redis/redis/v7"
	"github.com/nmarsollier/imagego/tools/custerror"
	"github.com/nmarsollier/imagego/tools/env"
)

// Insert agrega una imagen a la db
func Insert(image *Image) (string, error) {
	if err := image.validateSchema(); err != nil {
		return "", err
	}

	client := client()
	err := client.Set(image.ID, image.Image, 0).Err()
	if err != nil {
		return "", err
	}

	return image.ID, nil
}

// Find encuentra y devuelve una imagen desde la base de datos
func find(imageID string) (*Image, error) {
	client := client()
	data, err := client.Get(imageID).Result()
	if err != nil {
		return nil, custerror.NotFound
	}

	result := Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}

func client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     env.Get().RedisURL,
		Password: "",
		DB:       0,
	})
}
