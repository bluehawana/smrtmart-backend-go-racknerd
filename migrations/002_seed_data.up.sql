-- MySQL version of seed data migration

-- Insert sample categories
INSERT INTO categories (id, name, slug, description, is_active, sort_order) VALUES
('550e8400-e29b-41d4-a716-446655440010', 'Electronics', 'electronics', 'Electronic devices and gadgets', true, 1),
('550e8400-e29b-41d4-a716-446655440011', 'Computers', 'computers', 'Laptops, desktops, and computer accessories', true, 2),
('550e8400-e29b-41d4-a716-446655440012', 'Audio', 'audio', 'Headphones, speakers, and audio equipment', true, 3),
('550e8400-e29b-41d4-a716-446655440013', 'Wearables', 'wearables', 'Smart watches and wearable technology', true, 4),
('550e8400-e29b-41d4-a716-446655440014', 'Monitors', 'monitors', 'Computer monitors and displays', true, 5),
('550e8400-e29b-41d4-a716-446655440015', 'Networking', 'networking', 'Routers, switches, and networking equipment', true, 6),
('550e8400-e29b-41d4-a716-446655440016', 'Smartphones', 'smartphones', 'Mobile phones and accessories', true, 7);

-- Insert sample admin user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, status) VALUES
('550e8400-e29b-41d4-a716-446655440003', 'admin@smrtmart.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Admin', 'User', 'admin', 'active');

-- Insert sample vendor user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, status) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'vendor@smrtmart.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'John', 'Vendor', 'vendor', 'active');

-- Insert sample vendor
INSERT INTO vendors (id, user_id, business_name, business_type, description, address, status, verified_at) VALUES
('550e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440001', 'TechMart Store', 'Electronics Retailer', 'Premium electronics and technology products', 
'{"street": "123 Tech Street", "city": "San Francisco", "state": "CA", "postal_code": "94105", "country": "USA", "phone": "+1-555-0123"}', 
'approved', NOW());

-- Insert sample products
INSERT INTO products (id, vendor_id, name, description, price, compare_price, sku, category, tags, images, stock, status, featured, weight, dimensions, seo) VALUES
('550e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440002', 'MacBook Pro 16-inch', 
'Apple MacBook Pro 16-inch with M3 Pro chip, 18GB RAM, 512GB SSD. Perfect for professionals and creatives.', 
2499.00, 2799.00, 'MBP-16-M3-512', 'computers', 
JSON_ARRAY('apple', 'macbook', 'laptop', 'professional'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/macbook.jpg'), 
15, 'active', true, 2.1,
JSON_OBJECT('length', 35.57, 'width', 24.81, 'height', 1.68),
JSON_OBJECT('title', 'MacBook Pro 16-inch - Professional Laptop', 'description', 'Powerful MacBook Pro with M3 Pro chip for professional work', 'keywords', JSON_ARRAY('macbook', 'apple', 'laptop', 'professional', 'm3'))),

('550e8400-e29b-41d4-a716-446655440021', '550e8400-e29b-41d4-a716-446655440002', 'AirPods Pro 2nd Generation', 
'Apple AirPods Pro with Active Noise Cancellation, Transparency mode, and spatial audio.', 
249.00, 279.00, 'APP-2ND-GEN', 'audio', 
JSON_ARRAY('apple', 'airpods', 'wireless', 'noise-cancellation'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/airpods2.jpg'), 
50, 'active', true, 0.056,
JSON_OBJECT('length', 6.1, 'width', 4.5, 'height', 2.1),
JSON_OBJECT('title', 'AirPods Pro 2nd Generation - Wireless Earbuds', 'description', 'Premium wireless earbuds with active noise cancellation', 'keywords', JSON_ARRAY('airpods', 'apple', 'wireless', 'earbuds', 'noise-cancellation'))),

('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440002', 'Sony WH-1000XM5 Headphones', 
'Industry-leading noise canceling headphones with exceptional sound quality and 30-hour battery life.', 
399.00, 449.00, 'SONY-WH1000XM5', 'audio', 
JSON_ARRAY('sony', 'headphones', 'noise-cancellation', 'wireless'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/sony.jpg'), 
25, 'active', true, 0.25,
JSON_OBJECT('length', 26.4, 'width', 19.6, 'height', 8.0),
JSON_OBJECT('title', 'Sony WH-1000XM5 - Noise Canceling Headphones', 'description', 'Premium noise canceling wireless headphones', 'keywords', JSON_ARRAY('sony', 'headphones', 'wireless', 'noise-cancellation'))),

('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440002', 'Apple Watch Ultra 2', 
'The most rugged and capable Apple Watch. Featuring precision dual-frequency GPS and up to 36 hours of battery life.', 
799.00, 849.00, 'AW-ULTRA-2', 'wearables', 
JSON_ARRAY('apple', 'watch', 'smartwatch', 'fitness', 'gps'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/ultra.jpg'), 
20, 'active', true, 0.061,
JSON_OBJECT('length', 4.9, 'width', 4.4, 'height', 1.47),
JSON_OBJECT('title', 'Apple Watch Ultra 2 - Rugged Smartwatch', 'description', 'Most advanced Apple Watch for extreme sports and adventures', 'keywords', JSON_ARRAY('apple', 'watch', 'smartwatch', 'ultra', 'gps'))),

('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440002', 'iPhone 16 Pro Max', 
'iPhone 16 Pro Max with A18 Pro chip, advanced camera system, and titanium design.', 
1199.00, 1299.00, 'IP16-PM-256', 'smartphones', 
JSON_ARRAY('apple', 'iphone', 'smartphone', 'pro', 'camera'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/iphone.jpg'), 
30, 'active', true, 0.227,
JSON_OBJECT('length', 16.29, 'width', 7.79, 'height', 0.83),
JSON_OBJECT('title', 'iPhone 16 Pro Max - Premium Smartphone', 'description', 'Latest iPhone with advanced features and cameras', 'keywords', JSON_ARRAY('iphone', 'apple', 'smartphone', 'pro', 'camera'))),

('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440002', 'Dell XPS 15 (2023)', 
'Powerful laptop with Intel Core i7, 16GB RAM, and NVIDIA RTX 4050 graphics.', 
1899.00, 2099.00, 'DELL-XPS15-2023', 'computers', 
JSON_ARRAY('dell', 'xps', 'laptop', 'gaming', 'professional'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/dell-xps-15-2023.jpg'), 
12, 'active', false, 2.0,
JSON_OBJECT('length', 34.4, 'width', 23.0, 'height', 1.8),
JSON_OBJECT('title', 'Dell XPS 15 - High-Performance Laptop', 'description', 'Premium laptop for professionals and creators', 'keywords', JSON_ARRAY('dell', 'xps', 'laptop', 'performance'))),

('550e8400-e29b-41d4-a716-446655440026', '550e8400-e29b-41d4-a716-446655440002', 'ASUS ROG Strix Gaming Monitor 27"', 
'27-inch 1440p gaming monitor with 165Hz refresh rate and G-Sync compatibility.', 
449.00, 499.00, 'ASUS-ROG-27', 'monitors', 
JSON_ARRAY('asus', 'monitor', 'gaming', '1440p', 'g-sync'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/asus.jpg'), 
18, 'active', false, 5.8,
JSON_OBJECT('length', 61.4, 'width', 27.5, 'height', 36.7),
JSON_OBJECT('title', 'ASUS ROG Gaming Monitor - 27 inch QHD', 'description', 'High refresh rate gaming monitor with G-Sync', 'keywords', JSON_ARRAY('asus', 'monitor', 'gaming', 'qhd'))),

('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440002', 'Huawei GT 2 Pro Smartwatch', 
'Premium smartwatch with 14-day battery life, GPS, and comprehensive health monitoring.', 
299.00, 329.00, 'HUAWEI-GT2-PRO', 'wearables', 
JSON_ARRAY('huawei', 'smartwatch', 'fitness', 'gps', 'health'), 
JSON_ARRAY('https://d10qehs4k3bdf9.cloudfront.net/huawei-gt2-pro.jpg'), 
22, 'active', false, 0.052,
JSON_OBJECT('length', 4.6, 'width', 4.6, 'height', 1.1),
JSON_OBJECT('title', 'Huawei GT 2 Pro - Premium Smartwatch', 'description', 'Feature-rich smartwatch with long battery life', 'keywords', JSON_ARRAY('huawei', 'smartwatch', 'fitness', 'health')));