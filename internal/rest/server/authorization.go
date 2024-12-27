package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/internal/engine/errs"
	"github.com/nmarsollier/imagego/internal/engine/log"
	"github.com/nmarsollier/imagego/internal/security"
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

	deps := GinDi(c)
	c.Set("logger", deps.Logger().WithField(log.LOG_FIELD_USER_ID, user.ID))
}

func validateToken(c *gin.Context) (*security.User, error) {
	tokenString, err := getHeaderToken(c)
	if err != nil {
		return nil, errs.Unauthorized
	}

	deps := GinDi(c)
	user, err := deps.SecurityService().Validate(tokenString)
	if err != nil {
		return nil, errs.Unauthorized
	}

	return user, nil
}

// get token from Authorization header
func getHeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(strings.ToUpper(tokenString), "BEARER ") != 0 {
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}
