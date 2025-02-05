package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-order/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-order/internal/external"
	"github.com/ArdiSasongko/Ecommerce-order/internal/model"
	"github.com/ArdiSasongko/Ecommerce-order/internal/storage/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Order interface {
		CreateOrder(context.Context, *model.OrderPayload) (*model.OrderResponse, error)
		UpdateStatusOrder(context.Context, model.UpdateStatus) error
		GetOrder(context.Context, int32) (*model.OrderResponse, error)
		GetOrders(context.Context, int32) ([]model.OrderResponse, error)
	}
}

func NewService(db *pgxpool.Pool, auth auth.JWTAuth) Service {
	external := external.NewExternal()
	q := sqlc.New(db)
	return Service{
		Order: &OrderService{
			q:        q,
			db:       db,
			external: external,
		},
	}
}
