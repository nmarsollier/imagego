package errs

// Unauthorized the user is not authorized to access the resource
var Unauthorized = NewRestError(401, "Unauthorized")

// NotFound when a record is not found in the db
var NotFound = NewRestError(404, "Document not found")

// AlreadyExist when a record cannot be inserted into the db
var AlreadyExist = NewRestError(400, "Already exist")

// Internal this application does not know how to handle the error
var Internal = NewRestError(500, "Internal server error")

// - Creation of errors -
// NewRestError creates a new errCustom
func NewRestError(status int, message string) RestError {
	return &restError{
		status:  status,
		Message: message,
	}
}

//  - Some necessary definitions -

// RestError is an interface to define custom errors
type RestError interface {
	Status() int
	Error() string
}

// restError is a custom error for http
type restError struct {
	status  int
	Message string `json:"error"`
}

func (e *restError) Error() string {
	return e.Message
}

// Status http status code
func (e *restError) Status() int {
	return e.status
}
