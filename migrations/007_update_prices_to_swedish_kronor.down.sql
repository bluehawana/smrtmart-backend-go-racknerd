-- Rollback Swedish kronor prices to original USD prices

UPDATE products SET price = 149, compare_price = 249 
WHERE name LIKE '%Mtag Apple AirTag%';

UPDATE products SET price = 1999, compare_price = 2499 
WHERE name LIKE '%Huawei GT2 Pro%';

UPDATE products SET price = 599, compare_price = 799 
WHERE name LIKE '%MagSafe 3 Charging Cable%';

UPDATE products SET price = 199, compare_price = 299 
WHERE name LIKE '%MacBook Air M3%' AND name LIKE '%Case%';

UPDATE products SET price = 399, compare_price = 599 
WHERE name LIKE '%Apple 29W%';

UPDATE products SET price = 299, compare_price = 499 
WHERE name LIKE '%iPhone%' AND name LIKE '%MagSafe%' AND name LIKE '%Case%';

UPDATE products SET price = 199, compare_price = 299 
WHERE name LIKE '%Dell Thunderbolt%';

UPDATE products SET price = 249, compare_price = 349 
WHERE name LIKE '%Braided Magnetic%';

UPDATE products SET price = 1199, compare_price = 1299 
WHERE name LIKE '%iPhone 15 Pro Max%';

UPDATE products SET price = 899, compare_price = 999 
WHERE name LIKE '%ASUS ROG%';

UPDATE products SET price = 1899, compare_price = 2199 
WHERE name LIKE '%Dell XPS 15%';

UPDATE products SET price = 149, compare_price = 179 
WHERE name LIKE '%Smart Language Translator%';

UPDATE products SET price = 199, compare_price = 249 
WHERE name LIKE '%AI Translate%';

UPDATE products SET price = 799, compare_price = 849 
WHERE name LIKE '%Apple Watch Ultra%';

UPDATE products SET price = 1299, compare_price = 1499 
WHERE name LIKE '%Dell XPS 13%';

UPDATE products SET price = 399, compare_price = 449 
WHERE name LIKE '%Sony WH-1000XM5%';

UPDATE products SET price = 2499, compare_price = 2799 
WHERE name LIKE '%MacBook Pro 16-inch%';

UPDATE products SET price = 249, compare_price = 279 
WHERE name LIKE '%AirPods Pro 2nd%';