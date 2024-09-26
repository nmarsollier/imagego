package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/imagego/tools/errs"
)

// ErrorHandler a middleware to handle errors
func ErrorHandler(c *gin.Context) {
	c.Next()

	handleErrorIfNeeded(c)
}

func handleErrorIfNeeded(c *gin.Context) {
	lastErr := c.Errors.Last()
	if lastErr == nil {
		return
	}

	err := lastErr.Err
	if err == nil {
		return
	}

	handleError(c, err)
}

// handleError handles any error to serialize it as JSON to the client
func handleError(c *gin.Context, err interface{}) {
	// Check for known error types
	switch value := err.(type) {
	case errs.RestError:
		// These are validations made with NewCustom
		handleCustom(c, value)
	case errs.Validation:
		// These are validations made with NewValidation
		c.JSON(400, err)
	case validator.ValidationErrors:
		// These are validator validations used in structure validations
		handleValidationError(c, value)
	case error:
		// Other errors
		c.JSON(500, ErrorData{
			Error: value.Error(),
		})
	default:
		// Unknown error type, return internal error
		handleCustom(c, errs.Internal)
	}
}

func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	err := errs.NewValidation()

	for _, e := range validationErrors {
		err.Add(strings.ToLower(e.Field()), e.Tag())
	}

	c.JSON(400, err)
}

func handleCustom(c *gin.Context, err errs.RestError) {
	c.JSON(err.Status(), err)
}

type ErrorData struct {
	Error string `json:"error"`
}
