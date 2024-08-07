package custerror

import (
	"fmt"
	"strconv"
)

// - Algunos errors comunes en el sistema -

// Unauthorized el usuario no esta autorizado al recurso
var Unauthorized = NewCustom(401, "Unauthorized")

// NotFound cuando un registro no se encuentra en la db
var NotFound = NewCustom(400, "Document not found")

// Internal esta aplicación no sabe como manejar el error
var Internal = NewCustom(500, "Internal server error")

// NewCustom creates a new errCustom
func NewCustom(status int, message string) Custom {
	return &ErrCustom{
		status:  status,
		Message: message,
	}
}

//  - Algunas definiciones necesarias -

// Custom es una interfaz para definir errores custom
// La necesitamos para poder castear correctamente en el handler
type Custom interface {
	Status() int
	Error() string
}

// ErrCustom es un error personalizado para http
type ErrCustom struct {
	status  int
	Message string `json:"error"`
}

func (e *ErrCustom) Error() string {
	return fmt.Sprintf(e.Message)
}

// Status http status code
func (e *ErrCustom) Status() int {
	return e.status
}

// Status http status code
func (e *ErrCustom) String() string {
	return strconv.Itoa(e.Status()) + e.Error()
}
