-- name: InsertOrder :one
INSERT INTO orders (user_id, total_price, status) values ($1, $2, $3) RETURNING id;

-- name: InsertOrderOrderItem :exec
UPDATE orders SET orders_items = $1 WHERE id = $2;

-- name: UpdateStatus :exec
UPDATE orders SET status = $1 WHERE id = $2;

-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;