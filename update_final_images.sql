-- Update all products with real uploaded images from Supabase Storage
UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/macbook.jpg']
WHERE name = 'MacBook Pro 16-inch';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/airpods2.jpg']
WHERE name = 'AirPods Pro 2nd Generation';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/sony.jpg']
WHERE name = 'Sony WH-1000XM5 Headphones';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/xps.jpg']
WHERE name = 'Dell XPS 13 Laptop';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/dell.jpg']
WHERE name = 'Dell Alienware 34 Curved Monitor';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/ultra.jpg']
WHERE name = 'Apple Watch Ultra';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/ai-translate-pro.jpg']
WHERE name = 'AI Translate Earphones Pro';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/smart-translator.jpg']
WHERE name = 'Smart Language Translator Buds';