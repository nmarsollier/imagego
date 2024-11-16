package tools

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
)

func ValidateLoggedIn(ctx context.Context) (*security.User, error) {
	env := GqlCtx(ctx)

	tokenString, err := TokenString(ctx)
	if err != nil {
		return nil, err
	}

	user, err := security.Validate(tokenString, env...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ValidateAdmin(ctx context.Context) (*security.User, error) {
	user, err := ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	hasAdminPermission := false
	for _, permission := range user.Permissions {
		if permission == "admin" {
			hasAdminPermission = true
			break
		}
	}

	if !hasAdminPermission {
		return nil, errs.Unauthorized
	}

	return user, nil
}

// HeaderToken Token data from Authorization header
func TokenString(ctx context.Context) (string, error) {
	operationContext := graphql.GetOperationContext(ctx)
	tokenString := operationContext.Headers.Get("Authorization")

	if strings.Index(strings.ToUpper(tokenString), "BEARER ") == 0 {
		tokenString = tokenString[7:]
	} else {
		return "", errs.Unauthorized
	}

	return tokenString, nil
}
