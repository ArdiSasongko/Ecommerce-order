CREATE TABLE IF NOT EXISTS order_items (
    id serial PRIMARY KEY,
    order_id int NOT NULL,
    product_id int NOT NULL,
    variant_id int NOT NULL,
    quantity int NOT NULL,
    price numeric(10, 2) NOT NULL,
    created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);