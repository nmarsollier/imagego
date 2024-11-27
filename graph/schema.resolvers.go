package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/nmarsollier/imagego/graph/model"
	"github.com/nmarsollier/imagego/graph/resolvers"
)

// CreateImage is the resolver for the createImage field.
func (r *mutationResolver) CreateImage(ctx context.Context, imageBase64 string) (string, error) {
	return resolvers.CreateImage(ctx, imageBase64)
}

// GetImage is the resolver for the getImage field.
func (r *queryResolver) GetImage(ctx context.Context, id string, size model.ImageSize) (*model.Image, error) {
	return resolvers.FindImageByID(ctx, id, size)
}

// Mutation returns model.MutationResolver implementation.
func (r *Resolver) Mutation() model.MutationResolver { return &mutationResolver{r} }

// Query returns model.QueryResolver implementation.
func (r *Resolver) Query() model.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }