CREATE TABLE IF NOT EXISTS store_products(
    store_id INT NOT NULL,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_stores
    FOREIGN KEY (store_id)
    REFERENCES stores (store_id),
    CONSTRAINT fk_products
    FOREIGN KEY (product_id)
    REFERENCES products (product_id)
);