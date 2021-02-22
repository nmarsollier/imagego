package custerror

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/test"
)

func TestCustomError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	HandleError(context, NewCustom(400, "Custom Test"))

	response.Assert(400, "{\"error\":\"Custom Test\"}")
}

func TestConstantError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	HandleError(context, Unauthorized)

	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}

func TestValidationError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	validation := NewValidation()
	validation.Add("abc", "abd wrong")
	context.Error(validation)

	HandleError(context, validation)

	response.Assert(400, "{\"messages\":[{\"path\":\"abc\",\"message\":\"abd wrong\"}]}")
}

func TestDefaultError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	HandleError(context, errors.New("Custom Error"))

	response.Assert(500, "{\"error\":\"Custom Error\"}")
}
