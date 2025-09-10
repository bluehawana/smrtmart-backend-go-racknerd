-- Add numeric ID column to products table for cleaner URLs
-- Keep UUID as primary key for backend relationships, but add numeric ID for frontend

-- Add numeric_id column
ALTER TABLE products ADD COLUMN numeric_id INTEGER UNIQUE;

-- Create a sequence for auto-incrementing numeric IDs (PostgreSQL)
CREATE SEQUENCE product_numeric_id_seq START 1;

-- Update existing products with numeric IDs based on order
-- Using ROW_NUMBER() to assign sequential IDs starting from 1
WITH numbered_products AS (
  SELECT id, ROW_NUMBER() OVER (ORDER BY created_at) as row_num
  FROM products
)
UPDATE products 
SET numeric_id = numbered_products.row_num
FROM numbered_products 
WHERE products.id = numbered_products.id;

-- Make numeric_id NOT NULL after populating data
ALTER TABLE products ALTER COLUMN numeric_id SET NOT NULL;

-- Create index for fast numeric_id lookups
CREATE INDEX idx_products_numeric_id ON products(numeric_id);