package service

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/ArdiSasongko/Ecommerce-order/internal/external"
	"github.com/ArdiSasongko/Ecommerce-order/internal/model"
	"github.com/ArdiSasongko/Ecommerce-order/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var StatusOrder = map[string]bool{
	"pending":   true,
	"cancelled": true,
	"refunded":  true,
	"completed": true,
}

var FlowUpdate = map[string][]string{
	"pending":   {"cancelled", "refunded", "completed"},
	"refunded":  {"cancelled", "completed"},
	"completed": {"refunded"},
}

type OrderService struct {
	q        *sqlc.Queries
	db       *pgxpool.Pool
	external external.External
}

func (s *OrderService) insertOrder(ctx context.Context, req model.OrderPayload) (int32, string, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, "", err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	priceStr := fmt.Sprintf("%.2f", req.TotalPrice)
	priceNumeric := pgtype.Numeric{}
	if err := priceNumeric.Scan(priceStr); err != nil {
		return 0, "", err
	}

	req.Status = string(sqlc.OrderStatusPending)
	id, err := qtx.InsertOrder(ctx, sqlc.InsertOrderParams{
		UserID:     req.UserID,
		TotalPrice: priceNumeric,
		Status:     sqlc.OrderStatus(req.Status),
	})
	if err != nil {
		return 0, "", err
	}

	var ids []int32

	for _, item := range req.OrderItems {
		priceStr := fmt.Sprintf("%.2f", item.Price)
		priceNumeric := pgtype.Numeric{}
		if err := priceNumeric.Scan(priceStr); err != nil {
			return 0, "", err
		}
		idItems, err := qtx.InsertOrderItem(ctx, sqlc.InsertOrderItemParams{
			OrderID:   id,
			ProductID: item.ProductID,
			VariantID: item.VariantID,
			Quantity:  item.Quantity,
			Price:     priceNumeric,
		})

		if err != nil {
			return 0, "", err
		}

		ids = append(ids, idItems)
	}

	if err := qtx.InsertOrderOrderItem(ctx, sqlc.InsertOrderOrderItemParams{
		OrdersItems: ids,
		ID:          id,
	}); err != nil {
		return 0, "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, "", err
	}

	return id, req.Status, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *model.OrderPayload) (*model.OrderResponse, error) {
	id, status, err := s.insertOrder(ctx, *req)
	if err != nil {
		return nil, err
	}

	kafkaPayload := model.PaymentInitialPayload{
		OrderID:    id,
		TotalPrice: req.TotalPrice,
	}

	jsonPayload, err := json.Marshal(kafkaPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal :%w", err)
	}

	if err := s.external.Kafka.ProduceKafkaMessage(ctx, jsonPayload); err != nil {
		if err := s.q.UpdateStatus(ctx, sqlc.UpdateStatusParams{
			ID:     id,
			Status: sqlc.OrderStatusCancelled,
		}); err != nil {
			return nil, err
		}

		return nil, err
	}

	return &model.OrderResponse{
		TotalPrice: req.TotalPrice,
		Status:     status,
		UserID:     req.UserID,
	}, nil
}

func checkValidFlow(from, to string) bool {
	allowed, ok := FlowUpdate[from]
	if !ok {
		return false
	}
	return slices.Contains(allowed, to)
}

func (s *OrderService) UpdateStatusOrder(ctx context.Context, req model.UpdateStatus) error {
	if !StatusOrder[req.Status] {
		return fmt.Errorf("invalid status! only 'pending', 'cancelled', 'completed', 'refunded'")
	}

	order, err := s.q.GetOrderByID(ctx, req.OrderID)
	if err != nil {
		return err
	}

	if !checkValidFlow(string(order.Status), req.Status) {
		return fmt.Errorf("invalid flow update status")
	}

	if err := s.q.UpdateStatus(ctx, sqlc.UpdateStatusParams{
		Status: sqlc.OrderStatus(req.Status),
		ID:     req.OrderID,
	}); err != nil {
		return err
	}

	return nil
}
