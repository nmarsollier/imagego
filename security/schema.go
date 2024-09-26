package security

// User is the logged-in user
type User struct {
	ID          string   `json:"id"  validate:"required"`
	Name        string   `json:"name"  validate:"required"`
	Permissions []string `json:"permissions"`
	Login       string   `json:"login"  validate:"required"`
}
