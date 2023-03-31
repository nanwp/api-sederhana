package users

type UserResponse struct {
	Name     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
