-- Update product prices to Swedish kronor with competitive pricing (20-30% below market)
-- Based on Swedish market prices as of 2024

-- Mtag Apple AirTag Compatible Tracker - Market: ~300 kr, Our price: 199 kr (-34%)
UPDATE products SET price = 199, compare_price = 299 
WHERE name LIKE '%Mtag Apple AirTag%';

-- Huawei GT2 Pro Smart Watch - Market: ~2800 kr, Our price: 1999 kr (-29%)
UPDATE products SET price = 1999, compare_price = 2799 
WHERE name LIKE '%Huawei GT2 Pro%';

-- MacBook Pro MagSafe 3 Charging Cable - Market: ~850 kr, Our price: 599 kr (-30%)
UPDATE products SET price = 599, compare_price = 849 
WHERE name LIKE '%MagSafe 3 Charging Cable%';

-- MacBook Air M3 Case - Market: ~299 kr, Our price: 199 kr (-33%)
UPDATE products SET price = 199, compare_price = 299 
WHERE name LIKE '%MacBook Air M3%' AND name LIKE '%Case%';

-- Apple 29W USB-C Power Adapter - Market: ~599 kr, Our price: 399 kr (-33%)
UPDATE products SET price = 399, compare_price = 599 
WHERE name LIKE '%Apple 29W%';

-- iPhone 13/13 Pro MagSafe Case - Market: ~449 kr, Our price: 299 kr (-33%)
UPDATE products SET price = 299, compare_price = 449 
WHERE name LIKE '%iPhone%' AND name LIKE '%MagSafe%' AND name LIKE '%Case%';

-- Dell Thunderbolt 5/USB4 Cable - Market: ~299 kr, Our price: 199 kr (-33%)
UPDATE products SET price = 199, compare_price = 299 
WHERE name LIKE '%Dell Thunderbolt%';

-- Braided Magnetic Charging Cable - Market: ~349 kr, Our price: 249 kr (-29%)
UPDATE products SET price = 249, compare_price = 349 
WHERE name LIKE '%Braided Magnetic%';

-- iPhone 15 Pro Max - Market: ~15990 kr, Our price: 11999 kr (-25%)
UPDATE products SET price = 11999, compare_price = 15990 
WHERE name LIKE '%iPhone 15 Pro Max%';

-- ASUS ROG Gaming Router - Market: ~12990 kr, Our price: 8990 kr (-31%)
UPDATE products SET price = 8990, compare_price = 12990 
WHERE name LIKE '%ASUS ROG%';

-- Dell XPS 15 Developer Edition - Market: ~25990 kr, Our price: 18990 kr (-27%)
UPDATE products SET price = 18990, compare_price = 25990 
WHERE name LIKE '%Dell XPS 15%';

-- Smart Language Translator Buds - Market: ~1999 kr, Our price: 1499 kr (-25%)
UPDATE products SET price = 1499, compare_price = 1999 
WHERE name LIKE '%Smart Language Translator%';

-- AI Translate Earphones Pro - Market: ~2499 kr, Our price: 1990 kr (-20%)
UPDATE products SET price = 1990, compare_price = 2499 
WHERE name LIKE '%AI Translate%';

-- Apple Watch Ultra - Market: ~10490 kr, Our price: 7990 kr (-24%)
UPDATE products SET price = 7990, compare_price = 10490 
WHERE name LIKE '%Apple Watch Ultra%';

-- Dell XPS 13 Laptop - Market: ~17990 kr, Our price: 12990 kr (-28%)
UPDATE products SET price = 12990, compare_price = 17990 
WHERE name LIKE '%Dell XPS 13%';

-- Sony WH-1000XM5 Headphones - Market: ~4990 kr, Our price: 3990 kr (-20%)
UPDATE products SET price = 3990, compare_price = 4990 
WHERE name LIKE '%Sony WH-1000XM5%';

-- MacBook Pro 16-inch - Market: ~34990 kr, Our price: 24990 kr (-29%)
UPDATE products SET price = 24990, compare_price = 34990 
WHERE name LIKE '%MacBook Pro 16-inch%';

-- AirPods Pro 2nd Generation - Market: ~2990 kr, Our price = 2490 kr (-17%)
UPDATE products SET price = 2490, compare_price = 2990 
WHERE name LIKE '%AirPods Pro 2nd%';