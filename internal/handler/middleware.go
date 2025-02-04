package handler

import (
	"strings"

	"github.com/ArdiSasongko/Ecommerce-order/internal/external"
	"github.com/gofiber/fiber/v2"
)

type MiddlewareHandler struct {
	external external.External
}

func (m *MiddlewareHandler) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing token header authorization",
			})
		}

		rContext := ctx.Context()
		parts := strings.Split(authToken, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "header are malformed",
			})
		}

		token := parts[1]

		resp, err := m.external.User.Profile(rContext, token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		ctx.Locals("user", resp)
		return ctx.Next()
	}
}
