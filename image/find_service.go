package image

import (
	"github.com/nmarsollier/imagego/db"
	"github.com/nmarsollier/imagego/tools/errs"
)

// ErrSize the size is incorrect
var ErrSize = errs.NewValidation().Add("size", "invalid")

// Find searches for an image of a particular size
func Find(imageID string, size int, deps ...interface{}) (*db.Image, error) {
	if size <= 0 {
		return find(imageID, deps...)
	}

	sizedID := buildSizeID(imageID, size)

	// Search for the exact image size
	image, err := find(sizedID, deps...)
	if err != nil && err != errs.NotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	return findAndResize(imageID, size, deps...)
}

func findAndResize(imageID string, size int, deps ...interface{}) (*db.Image, error) {
	// The desired size is not found, search for the original,
	// resize it, save...
	image, err := find(imageID, deps...)
	if err != nil {
		return nil, err
	}

	image, err = resize(image, size, deps...)
	if err != nil {
		return nil, err
	}

	_, err = Insert(image, deps...)
	if err != nil {
		return nil, err
	}

	return image, nil
}
