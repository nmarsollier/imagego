package image

import (
	"github.com/nmarsollier/imagego/internal/engine/errs"
	uuid "github.com/satori/go.uuid"
)

// Image estructura de la imagen
type Image struct {
	ID    string `json:"id"  validate:"required"`
	Image string `json:"image"  validate:"required"`
}

// New crea una nueva imagen
func New(img string) *Image {
	return &Image{
		ID:    uuid.NewV4().String(),
		Image: img,
	}
}

// ErrData the image does not seem valid
var ErrData = errs.NewValidation().Add("image", "invalid")
