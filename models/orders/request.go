package orders

type OrderCreate struct {
	UserId      string `json:"user_id" binding:"required"`
	PaymentId   int    `json:"payment_id:" binding:"required"`
	TotalPrice  int    `json:"total_price" binding:"required"`
	TotalPaid   int    `json:"total_paid" binding:"required"`
	TotalReturn int    `json:"total_return" binding:"required"`
}

