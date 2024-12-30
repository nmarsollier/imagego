package graph

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/imagego/internal/env"
	"github.com/nmarsollier/imagego/internal/graph/model"
	"github.com/nmarsollier/imagego/internal/graph/schema"
)

func Start() {
	logger := log.Get(env.Get().FluentURL, env.Get().ServerName)
	port := env.Get().GqlPort
	srv := handler.NewDefaultServer(model.NewExecutableSchema(model.Config{Resolvers: &schema.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("GraphQL playground on port : ", port)
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", env.Get().GqlPort), nil))
}
