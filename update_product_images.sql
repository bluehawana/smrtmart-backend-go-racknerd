-- Update product images to match actual uploaded files in uploads folder
-- This script maps the existing products to their corresponding uploaded images

-- Update MacBook Pro image
UPDATE products 
SET images = ARRAY['macbook.jpg']
WHERE name LIKE '%MacBook%' AND category = 'computers';

-- Update AirPods Pro image  
UPDATE products 
SET images = ARRAY['airpods2.jpg']
WHERE name LIKE '%AirPods%' AND category = 'audio';

-- Update Sony headphones image
UPDATE products 
SET images = ARRAY['sony.jpg']
WHERE name LIKE '%Sony%' AND category = 'audio';

-- Update Dell XPS laptop image
UPDATE products 
SET images = ARRAY['dell.jpg']
WHERE name LIKE '%Dell XPS%' AND category = 'computers';

-- Update Dell Alienware monitor image (using ultra.jpg as it's likely a monitor)
UPDATE products 
SET images = ARRAY['ultra.jpg']
WHERE name LIKE '%Dell Alienware%' AND category = 'monitors';

-- Update Apple Watch Ultra image (using ultra.jpg)
UPDATE products 
SET images = ARRAY['ultra.jpg']
WHERE name LIKE '%Apple Watch Ultra%' AND category = 'wearables';

-- Update AI Translate Earphones image
UPDATE products 
SET images = ARRAY['ai-translate-pro.jpg']
WHERE name LIKE '%AI Translate%' AND category = 'audio';

-- Update Smart Language Translator image
UPDATE products 
SET images = ARRAY['smart-translator.jpg']
WHERE name LIKE '%Smart Language%' AND category = 'audio';

-- Update ASUS Router image
UPDATE products 
SET images = ARRAY['asus.jpg']
WHERE name LIKE '%ASUS%' AND category = 'networking';

-- Update iPhone 15 Pro Max image
UPDATE products 
SET images = ARRAY['iphone.jpg']
WHERE name LIKE '%iPhone%' AND category = 'smartphones';

-- Insert new products if they don't exist for remaining images

-- Check if we need to add XPS specific product (using xps.jpg)
INSERT INTO products (id, vendor_id, name, description, price, compare_price, sku, category, tags, images, stock, status, featured, weight, dimensions, seo)
SELECT 
    uuid_generate_v4(),
    '550e8400-e29b-41d4-a716-446655440002',
    'Dell XPS 15 Developer Edition',
    'Powerful Dell XPS 15 with Ubuntu pre-installed, Intel Core i9, 32GB RAM, 1TB SSD. Perfect for developers and content creators.',
    1899.00,
    2199.00,
    'DELL-XPS15-DEV',
    'computers',
    ARRAY['dell', 'xps', 'laptop', 'developer', 'ubuntu', 'linux'],
    ARRAY['xps.jpg'],
    12,
    'active',
    true,
    1.83,
    '{"length": 34.4, "width": 23.0, "height": 1.8}',
    '{"title": "Dell XPS 15 Developer Edition - Linux Laptop", "description": "High-performance laptop with Ubuntu for developers", "keywords": ["dell", "xps", "laptop", "developer", "ubuntu", "linux"]}'
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE images @> ARRAY['xps.jpg']
);

-- Show updated products with their images
SELECT name, category, images, status FROM products ORDER BY category, name;