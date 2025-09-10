-- Update product images to use Cloudflare R2 URLs
-- Using R2 public domain: https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev

-- MacBook Pro 16-inch
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/macbook.jpg'] 
WHERE name = 'MacBook Pro 16-inch';

-- AirPods Pro 2nd Generation  
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/airpods2.jpg']
WHERE name = 'AirPods Pro 2nd Generation';

-- Sony WH-1000XM5 Headphones
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/sony.jpg']
WHERE name = 'Sony WH-1000XM5 Headphones';

-- iPhone 16 Pro Max
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/iphone.jpg']
WHERE name LIKE '%iPhone 16 Pro Max%';

-- Dell XPS 15 (2023)
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/dell-xps-15-2023.jpg']
WHERE name LIKE '%Dell XPS 15%';

-- ASUS ROG Strix G15
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/asus.jpg']
WHERE name LIKE '%ASUS ROG%';

-- Apple Watch Ultra 2
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/ultra.jpg']
WHERE name LIKE '%Apple Watch Ultra%';

-- Huawei GT 2 Pro / Watch GT series
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/huawei-gt2-pro.jpg']
WHERE name LIKE '%Huawei%' AND name LIKE '%Watch%';

-- iPhone 16 Pro Max Case
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/iphone16-promax-case.jpg']
WHERE name LIKE '%iPhone%' AND name LIKE '%Case%' AND name LIKE '%16%';

-- MagSafe Case for iPhone
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/iphone-magsafe-case.jpg']
WHERE name LIKE '%MagSafe%' AND name LIKE '%Case%';

-- MacBook Air Case (Green)
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/macbook-air-case-green.jpg']
WHERE name LIKE '%MacBook Air%' AND name LIKE '%Case%' AND name LIKE '%Green%';

-- MacBook Air M3 Weaving Case  
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/macbookair-m3-weaving-case.jpg']
WHERE name LIKE '%MacBook Air%' AND name LIKE '%Weaving%';

-- MacBook M4 Charging Cable
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/macbook-m4-charging-cable.png']
WHERE name LIKE '%MacBook%' AND name LIKE '%M4%' AND name LIKE '%Charging%';

-- MacBook Air Adapter and Cable
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/macbookair-adaptor-cable.png']
WHERE name LIKE '%MacBook Air%' AND (name LIKE '%Adapter%' OR name LIKE '%Adaptor%') AND name LIKE '%Cable%';

-- 8K Data Cable Dell
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/8k-data-cable-dell.jpg']
WHERE name LIKE '%8K%' AND name LIKE '%Dell%';

-- Dell Thunderbolt Cable
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/dell-thunderbolt-cable.jpg']
WHERE name LIKE '%Dell%' AND name LIKE '%Thunderbolt%';

-- USB-C iPhone Cable
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/usb-c-iphone-cable.jpg']
WHERE name LIKE '%USB-C%' AND name LIKE '%iPhone%';

-- Magnetic Charging Cable
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/magnetic-charging-cable.jpg']
WHERE name LIKE '%Magnetic%' AND name LIKE '%Charging%';

-- MagSafe 3 Cable (Blue)
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/magsafe3-cable-blue.jpg']
WHERE name LIKE '%MagSafe 3%' AND name LIKE '%Blue%';

-- Apple 29W Adapter
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/apple-29w-adapter.jpg']
WHERE name LIKE '%Apple%' AND name LIKE '%29W%';

-- M Tracking Tag / AirTag Alternative
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/m-tracking-tag.jpg']
WHERE name LIKE '%Tracking%' AND name LIKE '%Tag%';

-- Smart Tracking Card
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/smart-tracking-card.jpg']
WHERE name LIKE '%Smart%' AND name LIKE '%Tracking%' AND name LIKE '%Card%';

-- AI Translate Pro
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/ai-translate-pro.jpg']
WHERE name LIKE '%AI%' AND name LIKE '%Translate%';

-- Smart Translator
UPDATE products SET images = ARRAY['https://pub-f181c83ced9f499bbd048ab1e553216c.r2.dev/products/smart-translator.jpg']
WHERE name LIKE '%Smart%' AND name LIKE '%Translator%';

-- Add indexes for better performance
CREATE INDEX IF NOT EXISTS idx_products_images ON products USING gin(images);
CREATE INDEX IF NOT EXISTS idx_products_name_search ON products USING gin(to_tsvector('english', name));