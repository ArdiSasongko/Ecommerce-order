package handler

import (
	"strconv"

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

func (h *OrderHandler) UpdateStatus(ctx *fiber.Ctx) error {
	payload := new(model.UpdateStatus)
	id := ctx.Params("orderID")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		log.WithError(fiber.ErrBadRequest).Error("convert error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	payload.OrderID = int32(orderID)

	if err := ctx.BodyParser(payload); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("body parse error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("validate error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Order.UpdateStatusOrder(ctx.Context(), *payload); err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("internal server :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func (h *OrderHandler) GetOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("orderID")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		log.WithError(fiber.ErrBadRequest).Error("convert error :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := h.service.Order.GetOrder(ctx.Context(), int32(orderID))
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("internal server :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}

func (h *OrderHandler) GetOrders(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*external.Response)
	resp, err := h.service.Order.GetOrders(ctx.Context(), user.Data.ID)
	if err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("internal server :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    resp,
	})
}
