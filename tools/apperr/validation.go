package apperr

import (
	"encoding/json"
)

// NewValidationField crea un error de validación para un solo campo
func NewValidationField(field string, err string) Validation {
	return &ErrValidation{
		Messages: []ErrField{
			{
				Path:    field,
				Message: err,
			},
		},
	}
}

// NewValidation crea un error de validación para un solo campo
func NewValidation() Validation {
	return &ErrValidation{
		Messages: []ErrField{},
	}
}

// Validation es una interfaz para definir errores custom
// Validation es un error de validaciones de parameteros o de campos
type Validation interface {
	Add(path string, message string) Validation
	Size() int
	Error() string
}

// ErrField define un campo inválido. path y mensaje de error
type ErrField struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

// ErrValidation es un error de validaciones de parameteros o de campos
type ErrValidation struct {
	Messages []ErrField `json:"messages"`
}

func (e *ErrValidation) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return "ErrValidation que no se puede pasar a json."
	}
	return string(body)
}

// Add agrega errores a un validation error
func (e *ErrValidation) Add(path string, message string) Validation {
	err := ErrField{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

// Size devuelve la cantidad de errores
func (e *ErrValidation) Size() int {
	return len(e.Messages)
}
