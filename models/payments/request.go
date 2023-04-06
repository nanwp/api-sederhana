package payments

type PaymentCreate struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Logo string `json:"logo" binding:"required"`
}
type PaymentUpdate struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}
