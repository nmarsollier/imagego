package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/custerror"
)

/**
 * @apiDefine AuthHeader
 *
 * @apiExample {String} Header Autorización
 *    Authorization=bearer {token}
 *
 * @apiErrorExample 401 Unauthorized
 *    HTTP/1.1 401 Unauthorized
 */

// AuthValidator validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func AuthValidator(c *gin.Context) {
	if err := validateToken(c); err != nil {
		custerror.HandleError(c, err)
		return
	}
}

var securityValidate func(token string) (*security.User, error) = security.Validate

func validateToken(c *gin.Context) error {
	tokenString, err := getTokenHeader(c)
	if err != nil {
		return custerror.Unauthorized
	}

	if _, err = securityValidate(tokenString); err != nil {
		return custerror.Unauthorized
	}

	return nil
}

// get token from Authorization header
func getTokenHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		return "", custerror.Unauthorized
	}
	return tokenString[7:], nil
}
