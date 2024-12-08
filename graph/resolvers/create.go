package resolvers

import (
	"context"

	"github.com/nmarsollier/imagego/graph/tools"
	"github.com/nmarsollier/imagego/image"
)

func CreateImage(ctx context.Context, imageBase64 string) (string, error) {
	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return "", err
	}

	env := tools.GqlCtx(ctx)
	id, err := image.Insert(image.NewImage(imageBase64), env...)
	if err != nil {
		return "", err
	}

	return id, nil
}
