CREATE TABLE IF NOT EXISTS deliveries(
    delivery_id SERIAL PRIMARY KEY,
    store_id INT NOT NULL,
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);