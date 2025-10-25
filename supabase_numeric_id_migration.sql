-- ===============================================
-- Supabase Migration: Add Numeric Product IDs
-- ===============================================
-- Run this in Supabase SQL Editor to add numeric_id column
-- This allows products to use numeric IDs (1, 2, 3) instead of UUIDs for frontend

-- Step 1: Add numeric_id column (nullable initially)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'products' AND column_name = 'numeric_id'
    ) THEN
        ALTER TABLE products ADD COLUMN numeric_id INTEGER;
    END IF;
END $$;

-- Step 2: Create sequence if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_sequences WHERE schemaname = 'public' AND sequencename = 'product_numeric_id_seq') THEN
        CREATE SEQUENCE product_numeric_id_seq START 1;
    END IF;
END $$;

-- Step 3: Update existing products with sequential numeric IDs
WITH numbered_products AS (
  SELECT id, ROW_NUMBER() OVER (ORDER BY created_at) as row_num
  FROM products
  WHERE numeric_id IS NULL
)
UPDATE products
SET numeric_id = numbered_products.row_num
FROM numbered_products
WHERE products.id = numbered_products.id;

-- Step 4: Set the sequence to continue from the highest numeric_id
SELECT setval('product_numeric_id_seq', COALESCE((SELECT MAX(numeric_id) FROM products), 0) + 1, false);

-- Step 5: Set numeric_id as NOT NULL with default value from sequence
ALTER TABLE products ALTER COLUMN numeric_id SET NOT NULL;
ALTER TABLE products ALTER COLUMN numeric_id SET DEFAULT nextval('product_numeric_id_seq');

-- Step 6: Add unique constraint on numeric_id
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'products_numeric_id_key'
    ) THEN
        ALTER TABLE products ADD CONSTRAINT products_numeric_id_key UNIQUE (numeric_id);
    END IF;
END $$;

-- Step 7: Create index for fast lookups
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes WHERE indexname = 'idx_products_numeric_id'
    ) THEN
        CREATE INDEX idx_products_numeric_id ON products(numeric_id);
    END IF;
END $$;

-- Verification query - run this to check the migration worked
SELECT
    id,
    numeric_id,
    name,
    created_at
FROM products
ORDER BY numeric_id
LIMIT 10;

-- ===============================================
-- Expected Result:
-- All products should now have numeric_id values
-- New products will auto-increment numeric_id
-- ===============================================
