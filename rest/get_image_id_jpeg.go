package rest

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/rest/server"
)

//	@Summary		Obtener jpeg
//	@Description	Obtiene una imagen del servidor en formato jpeg.
//	@Tags			Imagen
//	@Accept			json
//	@Produce		image/jpeg
//	@Param			correlation_id	header		string				true	"Logging Correlation Id"
//	@Param			Size			path		string				true	"[160|320|640|800|1024|1200]"
//	@Param			imageID			path		string				true	"ID de la imagen"
//	@Success		200				{file}		jpeg				"Imagen"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/image/:imageID/jpeg [get]
//
// Obtiene una imagen del servidor en formato jpeg.
func initGetImageIdJpeg() {
	server.Router().GET("/v1/image/:imageID/jpeg", sendJpegImage)
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
