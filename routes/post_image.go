package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/middlewares"
)

/**
 * @api {post} /v1/image Crear Imagen
 * @apiName Crear Imagen
 * @apiGroup Imagen
 *
 * @apiDescription Agrega una nueva imagen al servidor.
 *
 * @apiExample {json} Body
 *    {
 *      "image" : "{Imagen en formato Base 64}"
 *    }
 *
 * @apiSuccessExample {json} Respuesta
 *     HTTP/1.1 200 OK
 *     {
 *       "id": "{Id de imagen}"
 *     }
 *
 * @apiUse AuthHeader
 * @apiUse ParamValidationErrors
 * @apiUse OtherErrors
 */
func init() {
	router().POST(
		"/v1/image",
		middlewares.ValidateAuthentication,
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

	c.JSON(200, gin.H{
		"id": id,
	})
}

func getBodyImage(c *gin.Context) (string, error) {
	type NewRequest struct {
		Image string `json:"image" binding:"required"`
	}
	body := NewRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		return "", err
	}

	return body.Image, nil
}
