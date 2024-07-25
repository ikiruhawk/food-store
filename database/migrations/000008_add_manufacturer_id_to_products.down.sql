BEGIN;

ALTER TABLE products DROP CONSTRAINT fk_manufacturers;

COMMIT;