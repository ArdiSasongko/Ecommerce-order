CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    total_price numeric(10, 2) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    orders_items int[],
    created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
);