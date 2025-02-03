package handler

import "github.com/gofiber/fiber/v2"

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
}

func NewHandler() Handler {
	return Handler{
		Health: &HealthHandler{},
	}
}
