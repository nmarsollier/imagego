package image

import (
	"github.com/nmarsollier/imagego/tools/errs"
)

// ErrSize el tamaño es incorrecto
var ErrSize = errs.NewValidation().Add("size", "invalid")

// Find busca una imagen para un tamaño en particular
func Find(imageID string, size int, ctx ...interface{}) (*Image, error) {
	if size <= 0 {
		return find(imageID, ctx...)
	}

	sizedID := buildSizeID(imageID, size)

	// Busco el tamaño justo de imagen
	image, err := find(sizedID, ctx...)
	if err != nil && err != errs.NotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	return findAndResize(imageID, size, ctx...)
}

func findAndResize(imageID string, size int, ctx ...interface{}) (*Image, error) {
	// No se encuentra el tamaño buscado, buscamos la original,
	// y le ajustamos el tamaño, guardamos...
	image, err := find(imageID, ctx...)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size, ctx...)
	if err != nil {
		return nil, err
	}

	_, err = Insert(image, ctx...)
	if err != nil {
		return nil, err
	}

	return image, nil
}
