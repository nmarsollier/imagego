package main

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/tools/errors"
)

/**
 * @apiDefine SizeHeader
 *
 * @apiParamExample {String} Header Size
 *    Size=[160|320|640|800|1024|1200]
 */

// NewImage Crea una imagen nueva
/**
 * @api {post} /image Crear Imagen
 * @apiName CreateImage
 * @apiGroup Imagen
 *
 * @apiDescription Agrega una nueva imagen al servidor.
 *
 * @apiParamExample {json} Body
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
func NewImage(c *gin.Context) {
	if err := validateAuthentication(c); err != nil {
		errors.Handle(c, err)
		return
	}

	type NewRequest struct {
		Image string `json:"image" binding:"required"`
	}
	body := NewRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		errors.Handle(c, err)
		return
	}

	img := image.New()
	img.Image = body.Image

	id, err := image.Insert(img)

	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

// GetImage devuelve una imagen guardada
func getImage(c *gin.Context) (*image.Image, error) {
	if err := validateAuthentication(c); err != nil {
		return nil, err
	}

	imageID := c.Param("imageID")
	size := image.Size(c.GetHeader("Size"))

	var data *image.Image
	var err error

	if size > 0 {
		data, err = image.FindSize(imageID, size)
	} else {
		data, err = image.Find(imageID)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetImage devuelve una imagen guardada en formato base64
/**
 * @api {get} /image/:id Obtener Imagen
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
func GetImage(c *gin.Context) {
	data, err := getImage(c)

	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(200, data)
}

// GetImageJpeg obtiene la imagen en formato jpeg
/**
 * @api {get} /image/:id/jpeg Obtener Imagen Jpeg
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
func GetImageJpeg(c *gin.Context) {
	data, err := getImage(c)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	str := data.Image[strings.Index(data.Image, ",")+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	decodedData, err := ioutil.ReadAll(reader)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.Data(200, "image/jpeg", decodedData)
}
