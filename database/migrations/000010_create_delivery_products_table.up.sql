CREATE TABLE IF NOT EXISTS delivery_products(
    delivery_id INT NOT NULL,
    product_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_deliveries
    FOREIGN KEY (delivery_id)
    REFERENCES deliveries (delivery_id),
    CONSTRAINT fk_products
    FOREIGN KEY (product_id)
    REFERENCES products (product_id)
);