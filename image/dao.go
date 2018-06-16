package image

import (
	"github.com/nmarsollier/imagego/tools/db"
)

// Insert agrega una imagen a la db
func Insert(image *Image) (string, error) {
	if err := image.validateSchema(); err != nil {
		return "", err
	}

	client := db.Client()
	err := client.Set(image.ID, image.Image, 0).Err()
	if err != nil {
		return "", err
	}

	return image.ID, nil
}

// Find encuentra y devuelve una imagen desde la base de datos
func Find(imageID string) (*Image, error) {
	client := db.Client()
	data, err := client.Get(imageID).Result()
	if err != nil {
		return nil, db.ErrNotFound
	}

	result := Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}
