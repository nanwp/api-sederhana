package users

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "tbl_user"
}
