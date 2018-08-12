package request

// NewUserRequest - Request Body for New User
type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
