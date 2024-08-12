package routes

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/model/image"
)

// Obtiene una imagen del servidor en formato jpeg.
//	@Summary		Obtener jpeg
//	@Description	Obtiene una imagen del servidor en formato jpeg.
//	@Tags			Imagen
//	@Accept			json
//	@Produce		image/jpeg
//	@Param			Size	path		string					true	"[160|320|640|800|1024|1200]"
//	@Param			imageID	path		string					true	"ID de la imagen"
//	@Success		200		{file}		jpeg					"Imagen"
//	@Failure		400		{object}	apperr.ErrValidation	"Bad Request"
//	@Failure		404		{object}	apperr.ErrCustom		"Not Found"
//	@Failure		500		{object}	apperr.ErrCustom		"Internal Server Error"
//	@Router			/v1/image/:imageID/jpeg [get]
//
// init Inicializa la ruta
func init() {
	router().GET("/v1/image/:imageID/jpeg", sendJpegImage)
}

func sendJpegImage(c *gin.Context) {
	image, err := getImage(c)
	if err != nil {
		c.Error(err)
		return
	}

	decodedData, err := toJpeg(image)
	if err != nil {
		c.Error(err)
		return
	}

	c.Data(200, "image/jpeg", decodedData)
}

func toJpeg(data *image.Image) ([]byte, error) {
	str := data.Image[strings.Index(data.Image, ",")+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	return ioutil.ReadAll(reader)
}
