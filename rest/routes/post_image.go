package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/model/image"
	"github.com/nmarsollier/imagego/rest/engine"
)

// Agrega una nueva imagen al servidor.
// @Summary		Guardar imagen
// @Description	Agrega una nueva imagen al servidor.
// @Tags			Imagen
// @Accept			json
// @Produce		json
// @Param			image			body		NewRequest				true	"Imagen en base64"
// @Param			Authorization	header		string					true	"bearer {token}"
// @Success		200				{object}	NewImageResponse		"Imagen"
// @Failure		400				{object}	apperr.ValidationErr	"Bad Request"
// @Failure		401				{object}	engine.ErrorData		"Unauthorized"
// @Failure		404				{object}	engine.ErrorData		"Not Found"
// @Failure		500				{object}	engine.ErrorData		"Internal Server Error"
// @Router			/v1/image [post]
//
// Init inicializa la ruta
func init() {
	engine.Router().POST(
		"/v1/image",
		engine.ValidateAuthentication,
		saveImage,
	)
}

func saveImage(c *gin.Context) {
	bodyImage, err := getBodyImage(c)
	if err != nil {
		c.Error(err)
		return
	}

	id, err := image.Insert(image.New(bodyImage))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, NewImageResponse{ID: id})
}

func getBodyImage(c *gin.Context) (string, error) {
	body := NewRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		return "", err
	}

	return body.Image, nil
}

type NewRequest struct {
	Image string `json:"image" binding:"required"`
}

type NewImageResponse struct {
	ID string `json:"id"  validate:"required"`
}
