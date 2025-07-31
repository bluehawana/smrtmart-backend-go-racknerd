-- Update products with placeholder images from unsplash/placeholder services
-- These will display properly while you source real product images

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=500&h=500&fit=crop']
WHERE name = 'MacBook Pro 16-inch';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1606220945770-b5b6c2c55bf1?w=500&h=500&fit=crop']
WHERE name = 'AirPods Pro 2nd Generation';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1618366712010-f4ae9c647dcb?w=500&h=500&fit=crop']
WHERE name = 'Sony WH-1000XM5 Headphones';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=500&h=500&fit=crop']
WHERE name = 'Dell XPS 13 Laptop';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1527443224154-c4a3942d3acf?w=500&h=500&fit=crop']
WHERE name = 'Dell Alienware 34 Curved Monitor';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1551816230-ef5deaed4a26?w=500&h=500&fit=crop']
WHERE name = 'Apple Watch Ultra';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1590658268037-6bf12165a8df?w=500&h=500&fit=crop']
WHERE name = 'AI Translate Earphones Pro';

UPDATE products 
SET images = ARRAY['https://images.unsplash.com/photo-1572569511254-d8f925fe2cbb?w=500&h=500&fit=crop']
WHERE name = 'Smart Language Translator Buds';