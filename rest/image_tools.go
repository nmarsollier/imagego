package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/db"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/rest/server"
)

// GetImage returns a saved image in base64 format
func getImage(c *gin.Context) (*db.Image, error) {
	imageID := c.Param("imageID")
	size := getSizeParam(c)

	deps := server.GinDeps(c)
	data, err := image.Find(imageID, size, deps...)
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

// normalizeParamSize returns the size from the header
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
