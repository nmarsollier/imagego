package image

import (
	"github.com/nmarsollier/imagego/db"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
)

// Insert adds an image to the db
func Insert(image *db.Image, deps ...interface{}) (string, error) {
	if err := image.ValidateSchema(deps...); err != nil {
		log.Get(deps...).Error(err)
		return "", err
	}

	client := db.Get(deps...)
	_, err := client.Set(image.ID, image.Image)
	if err != nil {
		log.Get(deps...).Error(err)
		return "", err
	}

	return image.ID, nil
}

// Find finds and returns an image from the database
func find(imageID string, deps ...interface{}) (*db.Image, error) {
	client := db.Get(deps...)
	data, err := client.Get(imageID)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, errs.NotFound
	}

	result := db.Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}
