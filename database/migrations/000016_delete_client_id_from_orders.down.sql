BEGIN;

ALTER TABLE orders
ADD client_id INT;
ALTER TABLE orders
ADD CONSTRAINT fk_clients
FOREIGN KEY (client_id)
REFERENCES clients (client_id);

COMMIT;