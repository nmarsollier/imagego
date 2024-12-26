package resolvers

import (
	"context"

	"github.com/nmarsollier/imagego/internal/graph/tools"
	"github.com/nmarsollier/imagego/internal/image"
)

func CreateImage(ctx context.Context, imageBase64 string) (string, error) {

	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return "", err
	}
	env := tools.GqlDi(ctx)
	id, err := env.ImageService().Insert(image.New(imageBase64))
	if err != nil {
		return "", err
	}

	return id, nil
}
