// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_item.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getOrderItemsByOrderID = `-- name: GetOrderItemsByOrderID :many
SELECT id, order_id, product_id, variant_id, quantity, price, created_at, updated_at FROM order_items WHERE order_id = $1
`

func (q *Queries) GetOrderItemsByOrderID(ctx context.Context, orderID int32) ([]OrderItem, error) {
	rows, err := q.db.Query(ctx, getOrderItemsByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.VariantID,
			&i.Quantity,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrderItem = `-- name: InsertOrderItem :one
INSERT INTO order_items (order_id, product_id, variant_id, quantity, price) values ($1, $2, $3, $4, $5) RETURNING id
`

type InsertOrderItemParams struct {
	OrderID   int32
	ProductID int32
	VariantID int32
	Quantity  int32
	Price     pgtype.Numeric
}

func (q *Queries) InsertOrderItem(ctx context.Context, arg InsertOrderItemParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertOrderItem,
		arg.OrderID,
		arg.ProductID,
		arg.VariantID,
		arg.Quantity,
		arg.Price,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
