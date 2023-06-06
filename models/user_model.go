package models

import "time"

type User struct {
	ID        string
	Name      string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	Name     string `json:"nama"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserCreate struct {
	Name     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "tbl_user"
}
