package image

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
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

func (e *Image) validateSchema(ctx ...interface{}) error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		log.Get(ctx...).Error(err)
		return err
	}
	if !strings.Contains(e.Image, "data:image/") {
		log.Get(ctx...).Error(ErrData)
		return ErrData
	}
	return nil
}
