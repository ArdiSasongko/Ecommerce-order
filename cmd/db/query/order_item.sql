-- name: InsertOrderItem :one
INSERT INTO order_items (order_id, product_id, variant_id, quantity, price) values ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetOrderItemsByOrderID :many
SELECT * FROM order_items WHERE order_id = $1;