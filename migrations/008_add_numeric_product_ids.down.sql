-- Rollback numeric product IDs

-- Drop the index
DROP INDEX IF EXISTS idx_products_numeric_id;

-- Drop the sequence
DROP SEQUENCE IF EXISTS product_numeric_id_seq;

-- Drop the numeric_id column
ALTER TABLE products DROP COLUMN IF EXISTS numeric_id;