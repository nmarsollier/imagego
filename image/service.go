package image

import (
	"fmt"

	"github.com/nmarsollier/imagego/tools/db"
	"github.com/nmarsollier/imagego/tools/errors"
)

// ErrSize el tamaño es incorrecto
var ErrSize = errors.NewValidationField("size", "invalid")

// FindSize busca una imagen para un tamaño en particular
func FindSize(imageID string, size int) (*Image, error) {
	if size <= 0 {
		return nil, ErrSize
	}

	id := fmt.Sprintf("%s_%d", imageID, size)

	// Busco el tamaño ajustado
	image, err := Find(id)
	if err != nil && err != db.ErrNotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	// No se encuentra, buscamos la original, le ajustamos el tamaño, guardamos...
	image, err = Find(imageID)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size)
	if err != nil {
		return nil, err
	}
	id, err = Insert(image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// Size retorna el tamaño a partir del header
func Size(sizeHeader string) int {
	switch sizeHeader {
	case "160":
		{
			return 160
		}
	case "320":
		{
			return 320
		}
	case "640":
		{
			return 640
		}
	case "800":
		{
			return 800
		}
	case "1024":
		{
			return 1024
		}
	case "1200":
		{
			return 1200
		}
	}
	return 0
}
