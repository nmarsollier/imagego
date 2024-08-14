package engine

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

// handleError maneja cualquier error para serializarlo como JSON al cliente
func handleError(c *gin.Context, err interface{}) {
	// Compruebo tipos de errores conocidos
	switch value := err.(type) {
	case errs.RestError:
		// Son validaciones hechas con NewCustom
		handleCustom(c, value)
	case errs.Validation:
		// Son validaciones hechas con NewValidation
		c.JSON(400, err)
	case validator.ValidationErrors:
		// Son las validaciones de validator usadas en validaciones de estructuras
		handleValidationError(c, value)
	case error:
		// Otros errores
		c.JSON(500, ErrorData{
			Error: value.Error(),
		})
	default:
		// No se sabe que es, devolvemos internal
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
