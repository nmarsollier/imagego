package routes

import (
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /v1/image/:id Obtener Imagen
 * @apiName Obtener Imagen
 * @apiGroup Imagen
 *
 * @apiDescription Obtiene una imagen del servidor en formato base64
 *
 * @apiUse SizeHeader
 *
 * @apiSuccessExample {json} Respuesta
 *    {
 *      "id": "{Id de imagen}",
 *      "image" : "{Imagen en formato Base 64}"
 *    }
 *
 * @apiUse AuthHeader
 * @apiUse ParamValidationErrors
 * @apiUse OtherErrors
 */
func init() {
	getRouter().GET("/v1/image/:imageID", handleGetImageID)
}

func handleGetImageID(c *gin.Context) {
	data, err := getImage(c)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, data)
}
