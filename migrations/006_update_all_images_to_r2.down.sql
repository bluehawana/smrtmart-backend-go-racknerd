-- Rollback migration - revert R2 URLs back to simple filenames

UPDATE products SET images = ARRAY['mtrackingtag.jpg'] 
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/mtrackingtag.jpg'];

UPDATE products SET images = ARRAY['huaweismartwatch.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/huaweismartwatch.jpg'];

UPDATE products SET images = ARRAY['macbook m4 charging cable.png']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbook-m4-charging-cable.png'];

UPDATE products SET images = ARRAY['macbookair m3 weaving case.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbookair-m3-weaving-case.jpg'];

UPDATE products SET images = ARRAY['macbookair adaptor and cable.png']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbookair-adaptor-and-cable.png'];

UPDATE products SET images = ARRAY['iphone16 promaxcase.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/iphone16-promaxcase.jpg'];

UPDATE products SET images = ARRAY['8k data cable dell.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/8k-data-cable-dell.jpg'];

UPDATE products SET images = ARRAY['usb c iphone cable.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/usb-c-iphone-cable.jpg'];

UPDATE products SET images = ARRAY['iphone.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/iphone.jpg'];

UPDATE products SET images = ARRAY['asus.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/asus.jpg'];

UPDATE products SET images = ARRAY['dell.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/dell.jpg'];

UPDATE products SET images = ARRAY['smart-translator.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/smart-translator.jpg'];

UPDATE products SET images = ARRAY['ai-translate-pro.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/ai-translate-pro.jpg'];

UPDATE products SET images = ARRAY['ultra.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/ultra.jpg'];

UPDATE products SET images = ARRAY['sony.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/sony.jpg'];

UPDATE products SET images = ARRAY['macbook.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/macbook.jpg'];

UPDATE products SET images = ARRAY['airpods2.jpg']
WHERE images @> ARRAY['https://2a35af424f8734e497a5d707344d79d5.r2.cloudflarestorage.com/smrtmart/airpods2.jpg'];