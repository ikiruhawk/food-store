CREATE TABLE IF NOT EXISTS order_products(
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_orders
    FOREIGN KEY (order_id)
    REFERENCES orders (order_id),
    CONSTRAINT fk_products
    FOREIGN KEY (product_id)
    REFERENCES products (product_id)
);