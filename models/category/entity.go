package category

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "tbl_category"
}
