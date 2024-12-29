package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/internal/image"
	"github.com/nmarsollier/imagego/internal/rest/server"
)

// Adds a new image to the server.
//
//	@Summary		Save image
//	@Description	Adds a new image to the server.
//	@Tags			Image
//	@Accept			json
//	@Produce		json
//	@Param			image			body		NewRequest			true	"Image in base64"
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Param			correlation_id	header		string				true	"Logging Correlation Id"
//	@Success		200				{object}	NewImageResponse	"Image"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/images/create [post]
//
// Init initializes the route
func initPostImage(engine *gin.Engine) {
	engine.POST(
		"/images/create",
		server.ValidateAuthentication,
		saveImage,
	)
}

func saveImage(c *gin.Context) {
	bodyImage, err := getBodyImage(c)
	if err != nil {
		c.Error(err)
		return
	}

	deps := server.GinDi(c)
	id, err := deps.ImageService().Insert(image.New(bodyImage))
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
