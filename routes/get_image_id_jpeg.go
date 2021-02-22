package routes

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/tools/custerror"
)

/**
 * @api {get} /v1/image/:id/jpeg Obtener Imagen Jpeg
 * @apiName Obtener Imagen Jpeg
 * @apiGroup Imagen
 *
 * @apiDescription Obtiene una imagen del servidor en formato jpeg.
 *
 * @apiUse SizeHeader
 *
 * @apiSuccessExample Respuesta
 *    Imagen en formato jpeg
 *
 * @apiUse AuthHeader
 * @apiUse ParamValidationErrors
 * @apiUse OtherErrors
 */
func init() {
	getRouter().GET("/v1/image/:imageID/jpeg", handleGetImageJpeg)
}

func handleGetImageJpeg(c *gin.Context) {
	image, err := getImage(c)
	if err != nil {
		custerror.HandleError(c, err)
		return
	}

	decodedData, err := toJpeg(image)
	if err != nil {
		custerror.HandleError(c, err)
		return
	}

	c.Data(200, "image/jpeg", decodedData)
}

func toJpeg(data *image.Image) ([]byte, error) {
	str := data.Image[strings.Index(data.Image, ",")+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	return ioutil.ReadAll(reader)
}
