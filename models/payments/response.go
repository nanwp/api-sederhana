package payments

type PaymentResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}
