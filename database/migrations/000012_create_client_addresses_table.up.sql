CREATE TABLE IF NOT EXISTS client_addresses(
    client_address_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    address VARCHAR(30),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_clients
    FOREIGN KEY (client_id)
    REFERENCES clients (client_id)
);