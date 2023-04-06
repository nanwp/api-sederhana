package payments

import "time"

type Payment struct {
	ID        int
	Name      string
	Type      string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Payment) TableName() string {
	return "tbl_payment"
}
