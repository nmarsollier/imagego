package image

import (
	"github.com/nmarsollier/imagego/tools/errs"
)

// ErrSize the size is incorrect
var ErrSize = errs.NewValidation().Add("size", "invalid")

// Find searches for an image of a particular size
func Find(imageID string, size int, deps ...interface{}) (image *Image, err error) {
	if size <= 0 {
		return find(imageID, deps...)
	}

	sizedID := buildSizeID(imageID, size)

	// Search for the exact image size
	image, err = find(sizedID, deps...)
	if err != nil && err != errs.NotFound {
		return
	}

	if err != nil {
		image, err = findAndResize(imageID, size, deps...)
	}

	return
}

func findAndResize(imageID string, size int, deps ...interface{}) (image *Image, err error) {
	// The desired size is not found, search for the original,
	// resize it, save...
	image, err = find(imageID, deps...)
	if err != nil {
		return
	}

	image, err = resize(image, size, deps...)
	if err != nil {
		return
	}

	_, err = Insert(image, deps...)
	return
}
