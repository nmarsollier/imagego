package image

import (
	"github.com/nmarsollier/imagego/tools/errs"
)

// ErrSize the size is incorrect
var ErrSize = errs.NewValidation().Add("size", "invalid")

// Find searches for an image of a particular size
func Find(imageID string, size int, ctx ...interface{}) (*Image, error) {
	if size <= 0 {
		return find(imageID, ctx...)
	}

	sizedID := buildSizeID(imageID, size)

	// Search for the exact image size
	image, err := find(sizedID, ctx...)
	if err != nil && err != errs.NotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	return findAndResize(imageID, size, ctx...)
}

func findAndResize(imageID string, size int, ctx ...interface{}) (*Image, error) {
	// The desired size is not found, search for the original,
	// resize it, save...
	image, err := find(imageID, ctx...)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size, ctx...)
	if err != nil {
		return nil, err
	}

	_, err = Insert(image, ctx...)
	if err != nil {
		return nil, err
	}

	return image, nil
}
