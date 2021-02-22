package image

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/imagego/tools/custerror"
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

// ErrData la imagen no parece valida
var ErrData = custerror.NewValidationField("image", "invalid")

func (e *Image) validateSchema() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	if strings.Index(e.Image, "data:image/") < 0 {
		return ErrData
	}
	return nil
}
