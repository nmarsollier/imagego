package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/log"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
)

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	user, err := validateToken(c)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	ctx := GinCtx(c)
	c.Set("logger", log.Get(ctx...).WithField(log.LOG_FIELD_USER_ID, user.ID))
}

func validateToken(c *gin.Context) (*security.User, error) {
	tokenString, err := getHeaderToken(c)
	if err != nil {
		return nil, errs.Unauthorized
	}

	ctx := GinCtx(c)
	user, err := security.Validate(tokenString, ctx...)
	if err != nil {
		return nil, errs.Unauthorized
	}

	return user, nil
}

// get token from Authorization header
func getHeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}
