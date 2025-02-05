package handler

import (
	"github.com/ArdiSasongko/Ecommerce-order/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-order/internal/external"
	"github.com/ArdiSasongko/Ecommerce-order/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
	Order interface {
		CreateOrder(*fiber.Ctx) error
		UpdateStatus(*fiber.Ctx) error
	}
	Middleware interface {
		AuthMiddleware() fiber.Handler
	}
}

func NewHandler(db *pgxpool.Pool, auth auth.JWTAuth) Handler {
	service := service.NewService(db, auth)
	external := external.NewExternal()
	return Handler{
		Health: &HealthHandler{},
		Order: &OrderHandler{
			service: service,
		},
		Middleware: &MiddlewareHandler{
			external: external,
		},
	}
}
