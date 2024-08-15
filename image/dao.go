package image

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/redisx"
)

// Insert agrega una imagen a la db
func Insert(image *Image, ctx ...interface{}) (string, error) {
	if err := image.validateSchema(); err != nil {
		glog.Error(err)
		return "", err
	}

	client := redisx.Get(ctx...)
	err := client.Set(image.ID, image.Image, 0).Err()
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return image.ID, nil
}

// Find encuentra y devuelve una imagen desde la base de datos
func find(imageID string, ctx ...interface{}) (*Image, error) {
	client := redisx.Get(ctx...)
	data, err := client.Get(imageID).Result()
	if err != nil {
		glog.Error(err)
		return nil, errs.NotFound
	}

	result := Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}
