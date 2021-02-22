package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nmarsollier/imagego/tools/custerror"
)

// ErrorHandler a middleware to handle errors
func ErrorHandler(c *gin.Context) {
	c.Next()

	handleErrorIfNeeded(c)
}

func handleErrorIfNeeded(c *gin.Context) {
	err := c.Errors.Last().Err
	if err == nil {
		return
	}

	HandleError(c, err)
}

// HandleError maneja cualquier error para serializarlo como JSON al cliente
func HandleError(c *gin.Context, err interface{}) {
	// Compruebo tipos de errores conocidos
	switch value := err.(type) {
	case custerror.Custom:
		// Son validaciones hechas con NewCustom
		handleCustom(c, value)
	case custerror.Validation:
		// Son validaciones hechas con NewValidation
		c.JSON(400, err)
	case validator.ValidationErrors:
		// Son las validaciones de validator usadas en validaciones de estructuras
		handleValidationError(c, value)
	case error:
		// Otros errores
		c.JSON(500, gin.H{
			"error": value.Error(),
		})
	default:
		// No se sabe que es, devolvemos internal
		handleCustom(c, custerror.Internal)
	}
}

/**
 * @apiDefine ParamValidationErrors
 *
 * @apiErrorExample 400 Bad Request
 *     HTTP/1.1 400 Bad Request
 *     {
 *        "messages" : [
 *          {
 *            "path" : "{Nombre de la propiedad}",
 *            "message" : "{Motivo del error}"
 *          },
 *          ...
 *       ]
 *     }
 */
func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	err := custerror.NewValidation()

	for _, e := range validationErrors {
		err.Add(strings.ToLower(e.Field()), e.Tag())
	}

	c.JSON(400, err)
}

/**
 * @apiDefine OtherErrors
 *
 * @apiErrorExample 500 Server Error
 *     HTTP/1.1 500 Internal Server Error
 *     {
 *        "error" : "Not Found"
 *     }
 *
 */
func handleCustom(c *gin.Context, err custerror.Custom) {
	c.JSON(err.Status(), err)
}
