package engine

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
)

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	if err := validateToken(c); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
}

var securityValidate func(token string) (*security.User, error) = security.Validate

func validateToken(c *gin.Context) error {
	tokenString, err := getHeaderToken(c)
	if err != nil {
		return errs.Unauthorized
	}

	if _, err = securityValidate(tokenString); err != nil {
		return errs.Unauthorized
	}

	return nil
}

// get token from Authorization header
func getHeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}
