package middlewares

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/test"
	"github.com/nmarsollier/imagego/tools/custerror"
)

func TestNoHeader(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)

	AuthValidator(context)

	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}

func TestInvalidHeader(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)
	context.Request.Header.Add("Authorization", "b")

	securityValidate = func(token string) (*security.User, error) {
		return nil, custerror.Unauthorized
	}

	AuthValidator(context)

	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}

func TestUnautorized(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)
	context.Request.Header.Add("Authorization", "bearer 123")

	defer func(restore func(token string) (*security.User, error)) {
		securityValidate = restore
	}(securityValidate)

	securityValidate = func(token string) (*security.User, error) {
		return nil, custerror.Unauthorized
	}

	AuthValidator(context)

	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}
