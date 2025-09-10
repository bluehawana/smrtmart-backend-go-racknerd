-- Rollback: Revert product images back to local file names

-- MacBook Pro 16-inch
UPDATE products SET images = ARRAY['macbook.jpg'] 
WHERE name = 'MacBook Pro 16-inch';

-- AirPods Pro 2nd Generation  
UPDATE products SET images = ARRAY['airpods2.jpg']
WHERE name = 'AirPods Pro 2nd Generation';

-- Sony WH-1000XM5 Headphones
UPDATE products SET images = ARRAY['sony.jpg']
WHERE name = 'Sony WH-1000XM5 Headphones';

-- iPhone 16 Pro Max
UPDATE products SET images = ARRAY['iphone.jpg']
WHERE name LIKE '%iPhone 16 Pro Max%';

-- Dell XPS 15 (2023)
UPDATE products SET images = ARRAY['dell-xps-15-2023.jpg']
WHERE name LIKE '%Dell XPS 15%';

-- ASUS ROG Strix G15
UPDATE products SET images = ARRAY['asus.jpg']
WHERE name LIKE '%ASUS ROG%';

-- Apple Watch Ultra 2
UPDATE products SET images = ARRAY['ultra.jpg']
WHERE name LIKE '%Apple Watch Ultra%';

-- Huawei GT 2 Pro / Watch GT series
UPDATE products SET images = ARRAY['huawei-gt2-pro.jpg']
WHERE name LIKE '%Huawei%' AND name LIKE '%Watch%';

-- iPhone 16 Pro Max Case
UPDATE products SET images = ARRAY['iphone16-promax-case.jpg']
WHERE name LIKE '%iPhone%' AND name LIKE '%Case%' AND name LIKE '%16%';

-- Remove indexes
DROP INDEX IF EXISTS idx_products_images;
DROP INDEX IF EXISTS idx_products_name_search;