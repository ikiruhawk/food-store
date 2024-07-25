BEGIN;

ALTER TABLE products
ADD CONSTRAINT fk_manufacturers
FOREIGN KEY (manufacturer_id)
REFERENCES manufacturers (manufacturer_id);

COMMIT;