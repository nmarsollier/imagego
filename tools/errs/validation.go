package errs

import (
	"encoding/json"
)

// Validation is an interface to define custom errors
// Validation is an error for parameter or field validations
type Validation interface {
	Add(path string, message string) Validation
	Error() string
}

func NewValidation() Validation {
	return &ValidationErr{
		Messages: []errField{},
	}
}

type ValidationErr struct {
	Messages []errField `json:"messages"`
}

func (e *ValidationErr) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return "ErrValidation invalid."
	}
	return string(body)
}

// Add adds errors to a validation error
func (e *ValidationErr) Add(path string, message string) Validation {
	err := errField{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

// errField defines an invalid field. path and error message
type errField struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}
