package middlewares

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/test"
	"github.com/nmarsollier/imagego/tools/custerror"
)

func TestCustomError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	handleError(context, custerror.NewCustom(400, "Custom Test"))

	response.Assert(400, "{\"error\":\"Custom Test\"}")
}

func TestConstantError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	handleError(context, custerror.Unauthorized)

	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}

func TestValidationError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	validation := custerror.NewValidation()
	validation.Add("abc", "abd wrong")
	context.Error(validation)

	handleError(context, validation)

	response.Assert(400, "{\"messages\":[{\"path\":\"abc\",\"message\":\"abd wrong\"}]}")
}

func TestDefaultError(t *testing.T) {
	response := test.ResponseWriter(t)
	context, _ := gin.CreateTestContext(response)

	handleError(context, errors.New("Custom Error"))

	response.Assert(500, "{\"error\":\"Custom Error\"}")
}
