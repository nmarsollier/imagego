package main

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/tools/errors"
)

// NewImage Crea una imagen nueva
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
func GetImage(c *gin.Context) {
	if err := validateAuthentication(c); err != nil {
		errors.Handle(c, err)
		return
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
		errors.Handle(c, err)
		return
	}

	c.JSON(200, data)
}

func GetImageJpeg(c *gin.Context) {
	if err := validateAuthentication(c); err != nil {
		errors.Handle(c, err)
		return
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
