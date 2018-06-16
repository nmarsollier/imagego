package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/authgo/tools/errors"
	"github.com/nmarsollier/imagego/security"
)

// get token from Authorization header
func getTokenHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		return "", errors.Unauthorized
	}
	return tokenString[7:], nil
}

func validateAuthentication(c *gin.Context) error {
	tokenString, err := getTokenHeader(c)
	if err != nil {
		return errors.Unauthorized
	}

	if _, err = security.Validate(tokenString); err != nil {
		return errors.Unauthorized
	}

	return nil
}
