package service

import (
	"github.com/ArdiSasongko/Ecommerce-order/internal/config/auth"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct{}

func NewService(db *pgxpool.Pool, auth auth.JWTAuth) Service {
	return Service{}
}
