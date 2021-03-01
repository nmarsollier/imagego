package middlewares

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/model/security"
	"github.com/nmarsollier/imagego/tools/custerror"
	"github.com/nmarsollier/imagego/tools/test"
	"gopkg.in/go-playground/assert.v1"
)

func TestNoHeader(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)

	ValidateAuthentication(context)

	response.Assert(0, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}

func TestInvalidHeader(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest("GET", "/", nil)
	context.Request.Header.Add("Authorization", "b")

	securityValidate = func(token string) (*security.User, error) {
		return nil, custerror.Unauthorized
	}

	ValidateAuthentication(context)

	response.Assert(0, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
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

	ValidateAuthentication(context)

	response.Assert(0, "")
	assert.Equal(t, context.Errors.Last().Error(), custerror.Unauthorized.Error())
}
