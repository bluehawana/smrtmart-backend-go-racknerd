-- Rollback: Revert product images back to local file names

-- Remove indexes
DROP INDEX IF EXISTS idx_products_images;
DROP INDEX IF EXISTS idx_products_name_search;

-- Revert images to original local paths
UPDATE products SET images = ARRAY['macbook.jpg'] 
WHERE name = 'MacBook Pro 16-inch';

UPDATE products SET images = ARRAY['airpods2.jpg']
WHERE name = 'AirPods Pro 2nd Generation';

UPDATE products SET images = ARRAY['sony.jpg']
WHERE name = 'Sony WH-1000XM5 Headphones';

UPDATE products SET images = ARRAY['iphone.jpg']
WHERE name LIKE '%iPhone 16 Pro Max%';

UPDATE products SET images = ARRAY['dell-xps-15-2023.jpg']
WHERE name LIKE '%Dell XPS 15%';

UPDATE products SET images = ARRAY['asus.jpg']
WHERE name LIKE '%ASUS ROG%';

UPDATE products SET images = ARRAY['ultra.jpg']
WHERE name LIKE '%Apple Watch Ultra%';