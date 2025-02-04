package model

type (
	OrderPayload struct {
		UserID     int32               `json:"user_id"`
		TotalPrice float32             `json:"total_price"`
		Status     string              `json:"status"`
		OrderItems []OrderItemsPayload `json:"order_items"`
	}

	OrderItemsPayload struct {
		OrderID   int32   `json:"order_id"`
		ProductID int32   `json:"product_id"`
		VariantID int32   `json:"variant_id"`
		Quantity  int32   `json:"quantity"`
		Price     float32 `json:"price"`
	}
)

type (
	OrderResponse struct {
		UserID     int32                `json:"user_id"`
		TotalPrice float32              `json:"total_price"`
		Status     string               `json:"status"`
		OrderItems []OrderItemsResponse `json:"order_items"`
	}

	OrderItemsResponse struct {
		OrderID   int32   `json:"order_id"`
		ProductID int32   `json:"product_id"`
		VariantID int32   `json:"variant_id"`
		Quantity  int32   `json:"quantity"`
		Price     float32 `json:"price"`
	}
)
