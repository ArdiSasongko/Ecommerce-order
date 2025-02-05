package model

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

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

	UpdateStatus struct {
		Status  string `json:"status" validate:"required"`
		OrderID int32  `json:"order_id"`
	}
)

func (u UpdateStatus) Validate() error {
	return Validate.Struct(u)
}

type (
	OrderResponse struct {
		ID         int32                `json:"id"`
		UserID     int32                `json:"user_id"`
		TotalPrice float32              `json:"total_price"`
		Status     string               `json:"status"`
		OrderItems []OrderItemsResponse `json:"order_items"`
	}

	OrderItemsResponse struct {
		ID        int32   `json:"id"`
		OrderID   int32   `json:"order_id"`
		ProductID int32   `json:"product_id"`
		VariantID int32   `json:"variant_id"`
		Quantity  int32   `json:"quantity"`
		Price     float32 `json:"price"`
	}
)
