-- MySQL version - Update product images to use actual uploaded filenames with proper URL encoding
-- Using R2 CloudFront distribution: https://d10qehs4k3bdf9.cloudfront.net/

-- MacBook Pro 16-inch
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbook.jpg') 
WHERE name = 'MacBook Pro 16-inch';

-- AirPods Pro 2nd Generation  
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/airpods2.jpg')
WHERE name = 'AirPods Pro 2nd Generation';

-- Sony WH-1000XM5 Headphones
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/sony.jpg')
WHERE name = 'Sony WH-1000XM5 Headphones';

-- iPhone 16 Pro Max
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/iphone.jpg')
WHERE name LIKE '%iPhone 16 Pro Max%';

-- Dell XPS 15 (2023)
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/dell-xps-15-2023.jpg')
WHERE name LIKE '%Dell XPS 15%';

-- ASUS ROG Strix G15
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/asus.jpg')
WHERE name LIKE '%ASUS ROG%';

-- Apple Watch Ultra 2
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/ultra.jpg')
WHERE name LIKE '%Apple Watch Ultra%';

-- Huawei GT 2 Pro / Watch GT series
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/huaweismartwatch.jpg')
WHERE name LIKE '%Huawei%' AND name LIKE '%Watch%';

-- iPhone 16 Pro Max Case
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/iphone16-promaxcase.jpg')
WHERE name LIKE '%iPhone%' AND name LIKE '%Case%' AND name LIKE '%16%';

-- MagSafe Case for iPhone
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/iphone-magsafe-case.jpg')
WHERE name LIKE '%MagSafe%' AND name LIKE '%Case%';

-- MacBook Air Case (Green)
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbook-air-case-green.jpg')
WHERE name LIKE '%MacBook Air%' AND name LIKE '%Case%' AND name LIKE '%Green%';

-- MacBook Air M3 Weaving Case  
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbookair-m3-weaving-case.jpg')
WHERE name LIKE '%MacBook Air%' AND name LIKE '%Weaving%';

-- MacBook M4 Charging Cable
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbook-m4-charging-cable.png')
WHERE name LIKE '%MacBook%' AND name LIKE '%M4%' AND name LIKE '%Charging%';

-- MacBook Air Adapter and Cable
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbookair-adaptor-and-cable.png')
WHERE name LIKE '%MacBook Air%' AND (name LIKE '%Adapter%' OR name LIKE '%Adaptor%') AND name LIKE '%Cable%';

-- 8K Data Cable Dell
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/8k-data-cable-dell.jpg')
WHERE name LIKE '%8K%' AND name LIKE '%Dell%';

-- Dell Thunderbolt Cable
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/dell-thunderbolt-cable.jpg')
WHERE name LIKE '%Dell%' AND name LIKE '%Thunderbolt%';

-- USB-C iPhone Cable
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/usb-c-iphone-cable.jpg')
WHERE name LIKE '%USB-C%' AND name LIKE '%iPhone%';

-- Magnetic Charging Cable
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/magnetic-charging-cable.jpg')
WHERE name LIKE '%Magnetic%' AND name LIKE '%Charging%';

-- MagSafe 3 Cable (Blue)
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/magsafe3-cable-blue.jpg')
WHERE name LIKE '%MagSafe 3%' AND name LIKE '%Blue%';

-- Apple 29W Adapter
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/apple-29w-adapter.jpg')
WHERE name LIKE '%Apple%' AND name LIKE '%29W%';

-- M Tracking Tag / AirTag Alternative
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/mtrackingtag.jpg')
WHERE name LIKE '%Tracking%' AND name LIKE '%Tag%';

-- MTag Tracker
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/mtag-tracker.jpg')
WHERE name LIKE '%MTag%' OR name LIKE '%M-Tag%';

-- Smart Tracking Card
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/smart-tracking-card.jpg')
WHERE name LIKE '%Smart%' AND name LIKE '%Tracking%' AND name LIKE '%Card%';

-- AI Translate Pro
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/ai-translate-pro.jpg')
WHERE name LIKE '%AI%' AND name LIKE '%Translate%';

-- Smart Translator
UPDATE products SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/smart-translator.jpg')
WHERE name LIKE '%Smart%' AND name LIKE '%Translator%';

-- MySQL doesn't support GIN indexes like PostgreSQL
-- Basic indexes are already created in the initial schema