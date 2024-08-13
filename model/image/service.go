package image

import (
	"github.com/nmarsollier/imagego/tools/apperr"
)

// ErrSize el tamaño es incorrecto
var ErrSize = apperr.NewValidation().Add("size", "invalid")

// Find busca una imagen para un tamaño en particular
func Find(imageID string, size int) (*Image, error) {
	if size <= 0 {
		return find(imageID)
	}

	sizedID := buildSizeID(imageID, size)

	// Busco el tamaño justo de imagen
	image, err := find(sizedID)
	if err != nil && err != apperr.NotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	return findAndResize(imageID, size)
}

func findAndResize(imageID string, size int) (*Image, error) {
	// No se encuentra el tamaño buscado, buscamos la original,
	// y le ajustamos el tamaño, guardamos...
	image, err := find(imageID)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size)
	if err != nil {
		return nil, err
	}

	_, err = Insert(image)
	if err != nil {
		return nil, err
	}

	return image, nil
}
