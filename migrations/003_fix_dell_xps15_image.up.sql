-- MySQL version - Fix Dell XPS 15 Developer Edition image to use the correct dell-xps-15-2023.jpg file
UPDATE products 
SET images = JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/dell-xps-15-2023.jpg')
WHERE name = 'Dell XPS 15 Developer Edition';

-- Also update the description to be more specific
UPDATE products 
SET description = 'Dell XPS 15 Developer Edition with Ubuntu, Intel Core i7, 32GB RAM, 1TB SSD, NVIDIA GeForce RTX 4050. Perfect for developers and content creators.'
WHERE name = 'Dell XPS 15 Developer Edition';