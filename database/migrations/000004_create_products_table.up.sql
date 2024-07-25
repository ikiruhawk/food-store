CREATE TABLE IF NOT EXISTS products(
    product_id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    manufacturer_id INT NOT NULL,
    name VARCHAR(20) NOT NULL,
    price INT NOT NULL,
    amount INT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);