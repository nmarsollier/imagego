package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/rest/server"
)

// GetImage devuelve una imagen guardada en formato base64
func getImage(c *gin.Context) (*image.Image, error) {
	imageID := c.Param("imageID")
	size := getSizeParam(c)

	ctx := server.GinCtx(c)
	data, err := image.Find(imageID, size, ctx...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getSizeParam(c *gin.Context) int {
	headerSize, ok := c.GetQuery("Size")
	if !ok {
		headerSize = c.GetHeader("Size")
	}

	return normalizeParamSize(headerSize)
}

// normalizeParamSize retorna el tamaño a partir del header
func normalizeParamSize(sizeHeader string) int {
	switch sizeHeader {
	case "160":
		return 160
	case "320":
		return 320
	case "640":
		return 640
	case "800":
		return 800
	case "1024":
		return 1024
	case "1200":
		return 1200
	}
	return 0
}
