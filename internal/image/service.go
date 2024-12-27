package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/nmarsollier/imagego/internal/engine/errs"
	"github.com/nmarsollier/imagego/internal/engine/log"
)

// ErrSize the size is incorrect
var ErrSize = errs.NewValidation().Add("size", "invalid")

type ImageService interface {
	Insert(image *Image) (string, error)
	Find(imageID string, size int) (*Image, error)
}

func NewImageService(
	log log.LogRusEntry,
	repo ImageRepository,
) ImageService {
	return &imageService{
		log:  log,
		repo: repo,
	}
}

type imageService struct {
	log  log.LogRusEntry
	repo ImageRepository
}

func (s *imageService) Insert(image *Image) (string, error) {
	return s.repo.Insert(image)
}

// Find searches for an image of a particular size
func (s *imageService) Find(imageID string, size int) (*Image, error) {
	if size <= 0 {
		return s.repo.Find(imageID)
	}

	sizedID := buildSizeID(imageID, size)

	// Search for the exact image size
	image, err := s.repo.Find(sizedID)
	if err != nil && err != errs.NotFound {
		return nil, err
	}
	if err == nil {
		return image, nil
	}

	return s.findAndResize(imageID, size)
}

func (s *imageService) findAndResize(imageID string, size int) (*Image, error) {
	// The desired size is not found, search for the original,
	// resize it, save...
	image, err := s.repo.Find(imageID)
	if err != nil {
		return nil, err
	}

	image, err = s.resize(image, size)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.Insert(image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (s *imageService) resize(image *Image, size int) (*Image, error) {
	str := image.Image[strings.Index(image.Image, ",")+1:]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))

	img, err := imaging.Decode(reader)
	if err != nil {
		s.log.Error(err)
	}
	bounds := img.Bounds()
	if bounds.Size().X <= size {
		return &Image{
			ID:    buildSizeID(image.ID, size),
			Image: image.Image,
		}, nil
	}

	// Resize srcImage to width = 800px preserving the aspect ratio.
	dstImage := imaging.Resize(img, size, 0, imaging.Lanczos)

	var buffer bytes.Buffer
	writer := base64.NewEncoder(base64.StdEncoding, &buffer)
	imaging.Encode(writer, dstImage, imaging.JPEG, imaging.JPEGQuality(70))
	writer.Close()

	result := Image{
		ID:    buildSizeID(image.ID, size),
		Image: "data:image/jpeg;base64," + buffer.String(),
	}
	return &result, nil
}

func buildSizeID(imageID string, size int) string {
	return fmt.Sprintf("%s_%d", imageID, size)
}
