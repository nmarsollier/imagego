package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/rest/engine"
)

//	@Summary		Obtener imagen
//	@Description	Obtiene una imagen del servidor en formato base64
//	@Tags			Imagen
//	@Accept			json
//	@Produce		json
//	@Param			Size	path		string				true	"[160|320|640|800|1024|1200]"
//	@Param			imageID	path		string				true	"ID de la imagen"
//	@Success		200		{object}	image.Image			"Informacion de la Imagen"
//	@Failure		400		{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401		{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404		{object}	engine.ErrorData	"Not Found"
//	@Failure		500		{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/v1/image/:imageID [get]
//
// Obtiene una imagen del servidor en formato base64
func init() {
	engine.Router().GET("/v1/image/:imageID", sendImage)
}

func sendImage(c *gin.Context) {
	data, err := getImage(c)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, data)
}
