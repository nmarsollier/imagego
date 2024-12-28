package image

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/commongo/errs"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/redisx"
)

type ImageRepository interface {
	Insert(image *Image) (string, error)
	Find(imageID string) (*Image, error)
	ValidateSchema(image *Image) error
}

func NewImageRepository(
	log log.LogRusEntry,
	redisClient redisx.RedisClient,
) ImageRepository {
	return &imageRepository{
		log:         log,
		redisClient: redisClient,
	}
}

type imageRepository struct {
	log         log.LogRusEntry
	redisClient redisx.RedisClient
}

// Insert adds an image to the db
func (r *imageRepository) Insert(image *Image) (string, error) {
	if err := r.ValidateSchema(image); err != nil {

		r.log.Error(err)
		return "", err
	}

	_, err := r.redisClient.Set(image.ID, image.Image, 0)
	if err != nil {
		r.log.Error(err)
		return "", err
	}

	return image.ID, nil
}

// Find finds and returns an image from the database
func (r *imageRepository) Find(imageID string) (*Image, error) {
	data, err := r.redisClient.Get(imageID)
	if err != nil {
		r.log.Error(err)
		return nil, errs.NotFound
	}

	result := Image{
		ID:    imageID,
		Image: data,
	}
	return &result, nil
}

func (s *imageRepository) ValidateSchema(image *Image) error {
	validate := validator.New()
	if err := validate.Struct(image); err != nil {
		s.log.Error(err)
		return err
	}
	if !strings.Contains(image.Image, "data:image/") {
		s.log.Error(ErrData)
		return ErrData
	}
	return nil
}
