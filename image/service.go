package image

import (
	"github.com/nmarsollier/imagego/tools/custerror"
)

// ErrSize el tamaño es incorrecto
var ErrSize = custerror.NewValidationField("size", "invalid")

var daoFind func(imageID string) (*Image, error) = find
var daoInsert func(image *Image) (string, error) = Insert

// Find busca una imagen para un tamaño en particular
func Find(imageID string, size int) (*Image, error) {
	if size <= 0 {
		return daoFind(imageID)
	}

	sizedID := sizeID(imageID, size)

	// Busco el tamaño justo de imagen
	image, err := daoFind(sizedID)
	if err != nil && err != custerror.NotFound {
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
	image, err := daoFind(imageID)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size)
	if err != nil {
		return nil, err
	}

	_, err = daoInsert(image)
	if err != nil {
		return nil, err
	}

	return image, nil
}
