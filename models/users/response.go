package users

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Alamat   string `json:"alamat"`
}
