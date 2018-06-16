package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

func resize(image *Image, size int) (*Image, error) {
	str := image.Image[strings.Index(image.Image, ",")+1:]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))

	img, err := imaging.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()
	if bounds.Size().X <= size {
		return image, nil
	}

	// Resize srcImage to width = 800px preserving the aspect ratio.
	dstImage := imaging.Resize(img, size, 0, imaging.Lanczos)

	var buffer bytes.Buffer
	writer := base64.NewEncoder(base64.StdEncoding, &buffer)
	imaging.Encode(writer, dstImage, imaging.JPEG, imaging.JPEGQuality(70))
	writer.Close()

	result := Image{
		ID:    fmt.Sprintf("%s_%d", image.ID, size),
		Image: "data:image/jpeg;base64," + string(buffer.Bytes()),
	}
	return &result, nil
}
