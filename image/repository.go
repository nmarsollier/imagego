package image

import (
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
	"github.com/nmarsollier/imagego/tools/redisx"
)

// Insert adds an image to the db
func Insert(image *Image, deps ...interface{}) (string, error) {
	if err := image.validateSchema(deps...); err != nil {
		log.Get(deps...).Error(err)
		return "", err
	}

	client := redisx.Get(deps...)
	err := client.Set(image.ID, image.Image, 0).Err()
	if err != nil {
		log.Get(deps...).Error(err)
		return "", err
	}

	return image.ID, nil
}

// Find finds and returns an image from the database
func find(imageID string, deps ...interface{}) (*Image, error) {
	client := redisx.Get(deps...)
	data, err := client.Get(imageID).Result()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, errs.NotFound
	}

	result := Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}
