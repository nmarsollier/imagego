package routes

import (
	"github.com/gin-gonic/gin"
)

// Obtiene una imagen del servidor en formato base64
//
//	@Summary		Obtener imagen
//	@Description	Obtiene una imagen del servidor en formato base64
//	@Tags			Imagen
//	@Accept			json
//	@Produce		json
//	@Param			Size	path		string					true	"[160|320|640|800|1024|1200]"
//	@Param			imageID	path		string					true	"ID de la imagen"
//	@Success		200		{object}	image.Image				"Informacion de la Imagen"
//	@Failure		400		{object}	apperr.ErrValidation	"Bad Request"
//	@Failure		404		{object}	apperr.ErrCustom		"Not Found"
//	@Failure		500		{object}	apperr.ErrCustom		"Internal Server Error"
//	@Router			/v1/image/:imageID [get]
//
// Init inicializa la ruta
func init() {
	router().GET("/v1/image/:imageID", sendImage)
}

func sendImage(c *gin.Context) {
	data, err := getImage(c)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, data)
}
