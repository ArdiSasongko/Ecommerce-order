package handler

import (
	"github.com/ArdiSasongko/Ecommerce-order/internal/config/logger"
	"github.com/ArdiSasongko/Ecommerce-order/internal/external"
	"github.com/ArdiSasongko/Ecommerce-order/internal/model"
	"github.com/ArdiSasongko/Ecommerce-order/internal/service"
	"github.com/gofiber/fiber/v2"
)

var log = logger.NewLogger()

type OrderHandler struct {
	service service.Service
}

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*external.Response)
	payload := new(model.OrderPayload)
	payload.UserID = user.Data.ID

	if err := ctx.BodyParser(payload); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("body parse error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := h.service.Order.CreateOrder(ctx.Context(), payload)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}
