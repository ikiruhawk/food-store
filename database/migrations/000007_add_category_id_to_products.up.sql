BEGIN;

ALTER TABLE products
ADD CONSTRAINT fk_categories
FOREIGN KEY (category_id)
REFERENCES categories(category_id);

COMMIT;