package rest

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/internal/image"
)

//	@Summary		Get jpeg
//	@Description	Gets an image from the server in jpeg format.
//	@Tags			Image
//	@Accept			json
//	@Produce		image/jpeg
//	@Param			correlation_id	header		string				true	"Logging Correlation Id"
//	@Param			Size			path		string				true	"[160|320|640|800|1024|1200]"
//	@Param			imageID			path		string				true	"Image ID"
//	@Success		200				{file}		jpeg				"Image"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/images/:imageID/jpeg [get]
//
// Gets an image from the server in jpeg format.
func initGetImageIdJpeg(engine *gin.Engine) {
	engine.GET("/images/:imageID/jpeg", sendJpegImage)
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
