package test

import (
	err "errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/authgo/tools/test"
	"github.com/nmarsollier/imagego/tools/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestHandleUnauthorized(t *testing.T) {
	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, errors.Unauthorized)
	response.Assert(401, "{\"error\":\"Unauthorized\"}")
}

func TestHandleNotFound(t *testing.T) {
	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, errors.NotFound)
	response.Assert(400, "{\"error\":\"Document not found\"}")
}

func TestHandleInternal(t *testing.T) {
	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, errors.Internal)
	response.Assert(500, "{\"error\":\"Internal server error\"}")
}

func TestHandleNewValidation(t *testing.T) {
	validation := errors.NewValidation()
	validation.Add("f1", "Ef1")
	validation.Add("f2", "Ef2")

	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, validation)
	response.Assert(400, "{\"messages\":[{\"path\":\"f1\",\"message\":\"Ef1\"},{\"path\":\"f2\",\"message\":\"Ef2\"}]}")
}

func TestHandleValidationError(t *testing.T) {
	type validStruct struct {
		OkField  string
		Required string `validate:"required"`
		Min      string `validate:"min=5"`
		Max      string `validate:"max=1"`
	}

	e := &validStruct{
		Min: "a",
		Max: "ab",
	}

	err := validator.New().Struct(e)

	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, err)
	response.Assert(400, "{\"messages\":[{\"path\":\"required\",\"message\":\"required\"},{\"path\":\"min\",\"message\":\"min\"},{\"path\":\"max\",\"message\":\"max\"}]}")
}

func TestHandleError(t *testing.T) {
	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, err.New("Test"))
	response.Assert(500, "{\"error\":\"Test\"}")
}

func TestHandleNotError(t *testing.T) {
	response := test.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	errors.Handle(context, "Test")
	response.Assert(500, "{\"error\":\"Internal server error\"}")
}
