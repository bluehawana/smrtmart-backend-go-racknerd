-- Update product images to use Supabase Storage URLs
-- Base URL: https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/macbook-pro-16.jpg']
WHERE name = 'MacBook Pro 16-inch';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/airpods-pro-2.jpg']
WHERE name = 'AirPods Pro 2nd Generation';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/sony-wh1000xm5.jpg']
WHERE name = 'Sony WH-1000XM5 Headphones';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/dell-xps-13.jpg']
WHERE name = 'Dell XPS 13 Laptop';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/dell-alienware-34.jpg']
WHERE name = 'Dell Alienware 34 Curved Monitor';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/apple-watch-ultra.jpg']
WHERE name = 'Apple Watch Ultra';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/ai-translate-earphones.jpg']
WHERE name = 'AI Translate Earphones Pro';

UPDATE products 
SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/smart-translator-buds.jpg']
WHERE name = 'Smart Language Translator Buds';