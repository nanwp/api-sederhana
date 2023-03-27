package users

import "time"

type User struct {
	ID        string
	Email     string
	Username  string
	Password  string
	Phone     string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
