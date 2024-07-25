BEGIN;

ALTER TABLE deliveries
ADD CONSTRAINT fk_stores
FOREIGN KEY (store_id)
REFERENCES stores(store_id);

COMMIT;