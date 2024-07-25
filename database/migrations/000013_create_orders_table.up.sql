CREATE TABLE IF NOT EXISTS orders(
    order_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    client_address_id INT NOT NULL,
    date TIMESTAMP NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_clients
    FOREIGN KEY (client_id)
    REFERENCES clients (client_id),
    CONSTRAINT fk_client_addresses
    FOREIGN KEY (client_address_id)
    REFERENCES client_addresses (client_address_id)
);