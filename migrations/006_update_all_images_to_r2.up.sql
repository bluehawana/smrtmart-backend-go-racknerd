-- Final migration to ensure all product images use CloudFlare R2 URLs
-- Using the updated R2 URL from user: https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart

-- Update products that still have simple filenames to use full R2 URLs
UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/mtrackingtag.jpg'] 
WHERE images @> ARRAY['mtrackingtag.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/huaweismartwatch.jpg']
WHERE images @> ARRAY['huaweismartwatch.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbook-m4-charging-cable.png']
WHERE images @> ARRAY['macbook m4 charging cable.png'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbookair-m3-weaving-case.jpg']
WHERE images @> ARRAY['macbookair m3 weaving case.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbookair-adaptor-and-cable.png']
WHERE images @> ARRAY['macbookair adaptor and cable.png'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/iphone16-promaxcase.jpg']
WHERE images @> ARRAY['iphone16 promaxcase.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/8k-data-cable-dell.jpg']
WHERE images @> ARRAY['8k data cable dell.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/usb-c-iphone-cable.jpg']
WHERE images @> ARRAY['usb c iphone cable.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/iphone.jpg']
WHERE images @> ARRAY['iphone.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/asus.jpg']
WHERE images @> ARRAY['asus.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/dell.jpg']
WHERE images @> ARRAY['dell.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/smart-translator.jpg']
WHERE images @> ARRAY['smart-translator.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/ai-translate-pro.jpg']
WHERE images @> ARRAY['ai-translate-pro.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/ultra.jpg']
WHERE images @> ARRAY['ultra.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/sony.jpg']
WHERE images @> ARRAY['sony.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbook.jpg']
WHERE images @> ARRAY['macbook.jpg'];

UPDATE products SET images = ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/airpods2.jpg']
WHERE images @> ARRAY['airpods2.jpg'];