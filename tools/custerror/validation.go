package custerror

import (
	"encoding/json"
	"fmt"
)

// NewValidationField crea un error de validación para un solo campo
func NewValidationField(field string, err string) Validation {
	return &errValidation{
		Messages: []errField{
			{
				Path:    field,
				Message: err,
			},
		},
	}
}

// NewValidation crea un error de validación para un solo campo
func NewValidation() Validation {
	return &errValidation{
		Messages: []errField{},
	}
}

// Validation es una interfaz para definir errores custom
// Validation es un error de validaciones de parameteros o de campos
type Validation interface {
	Add(path string, message string) Validation
	Size() int
	Error() string
}

// errField define un campo inválido. path y mensaje de error
type errField struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

// ErrValidation es un error de validaciones de parameteros o de campos
type errValidation struct {
	Messages []errField `json:"messages"`
}

func (e *errValidation) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ErrValidation que no se puede pasar a json.")
	}
	return fmt.Sprintf(string(body))
}

// Add agrega errores a un validation error
func (e *errValidation) Add(path string, message string) Validation {
	err := errField{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

// Size devuelve la cantidad de errores
func (e *errValidation) Size() int {
	return len(e.Messages)
}
