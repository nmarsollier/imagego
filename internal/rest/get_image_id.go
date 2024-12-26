package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/internal/rest/server"
)

//	@Summary		Get image
//	@Description	Gets an image from the server in base64 format
//	@Tags			Image
//	@Accept			json
//	@Produce		json
//	@Param			correlation_id	header		string				true	"Logging Correlation Id"
//	@Param			Size			path		string				true	"[160|320|640|800|1024|1200]"
//	@Param			imageID			path		string				true	"Image ID"
//	@Success		200				{object}	image.Image			"Image Information"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/image/:imageID [get]
//
// Gets an image from the server in base64 format
func initGetImageId() {
	server.Router().GET("/v1/image/:imageID", sendImage)
}

func sendImage(c *gin.Context) {
	data, err := getImage(c)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, data)
}
