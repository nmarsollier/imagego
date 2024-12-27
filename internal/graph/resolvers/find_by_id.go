package resolvers

import (
	"context"

	"github.com/nmarsollier/imagego/internal/graph/model"
	"github.com/nmarsollier/imagego/internal/graph/tools"
)

func FindImageByID(ctx context.Context, id string, size model.ImageSize) (*model.Image, error) {
	env := tools.GqlDi(ctx)

	imageSize := 0
	switch size {
	case model.ImageSizeSize160:
		imageSize = 160
	case model.ImageSizeSize320:
		imageSize = 320
	case model.ImageSizeSize640:
		imageSize = 640
	case model.ImageSizeSize800:
		imageSize = 800
	case model.ImageSizeSize1024:
		imageSize = 1024
	case model.ImageSizeSize1200:
		imageSize = 1200
	}

	data, err := env.ImageService().Find(id, imageSize)
	if err != nil {
		return nil, err
	}

	return &model.Image{
		ID:    data.ID,
		Image: data.Image,
	}, nil
}
