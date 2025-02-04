package model

type (
	PaymentInitialPayload struct {
		OrderID    int32   `json:"order_id"`
		TotalPrice float32 `json:"total_price"`
	}
)
