-- Insert sample categories
INSERT INTO categories (id, name, slug, description, is_active, sort_order) VALUES
(uuid_generate_v4(), 'Electronics', 'electronics', 'Electronic devices and gadgets', true, 1),
(uuid_generate_v4(), 'Computers', 'computers', 'Laptops, desktops, and computer accessories', true, 2),
(uuid_generate_v4(), 'Audio', 'audio', 'Headphones, speakers, and audio equipment', true, 3),
(uuid_generate_v4(), 'Wearables', 'wearables', 'Smart watches and wearable technology', true, 4),
(uuid_generate_v4(), 'Monitors', 'monitors', 'Computer monitors and displays', true, 5),
(uuid_generate_v4(), 'Networking', 'networking', 'Routers, switches, and networking equipment', true, 6),
(uuid_generate_v4(), 'Smartphones', 'smartphones', 'Mobile phones and accessories', true, 7);

-- Insert sample admin user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, status) VALUES
(uuid_generate_v4(), 'admin@smrtmart.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Admin', 'User', 'admin', 'active');

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
(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'MacBook Pro 16-inch', 
'Apple MacBook Pro 16-inch with M3 Pro chip, 18GB RAM, 512GB SSD. Perfect for professionals and creatives.', 
2499.00, 2799.00, 'MBP-16-M3-512', 'computers', 
ARRAY['apple', 'macbook', 'laptop', 'professional'], 
ARRAY['macbook-pro-16.jpg', 'macbook-pro-16-2.jpg'], 
15, 'active', true, 2.1,
'{"length": 35.57, "width": 24.81, "height": 1.68}',
'{"title": "MacBook Pro 16-inch - Professional Laptop", "description": "Powerful MacBook Pro with M3 Pro chip for professional work", "keywords": ["macbook", "apple", "laptop", "professional", "m3"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'AirPods Pro 2nd Generation', 
'Apple AirPods Pro with Active Noise Cancellation, Transparency mode, and spatial audio.', 
249.00, 279.00, 'APP-2ND-GEN', 'audio', 
ARRAY['apple', 'airpods', 'wireless', 'noise-cancellation'], 
ARRAY['airpods-pro-2.jpg'], 
50, 'active', true, 0.056,
'{"length": 6.1, "width": 4.5, "height": 2.1}',
'{"title": "AirPods Pro 2nd Generation - Wireless Earbuds", "description": "Premium wireless earbuds with active noise cancellation", "keywords": ["airpods", "apple", "wireless", "earbuds", "noise-cancellation"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'Sony WH-1000XM5 Headphones', 
'Industry-leading noise canceling headphones with exceptional sound quality and 30-hour battery life.', 
399.00, 449.00, 'SONY-WH1000XM5', 'audio', 
ARRAY['sony', 'headphones', 'noise-cancellation', 'wireless'], 
ARRAY['sony-wh1000xm5.jpg'], 
25, 'active', true, 0.25,
'{"length": 26.4, "width": 19.5, "height": 8.0}',
'{"title": "Sony WH-1000XM5 - Premium Noise Canceling Headphones", "description": "Professional noise canceling headphones with superior sound", "keywords": ["sony", "headphones", "noise-canceling", "wireless", "premium"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'Dell XPS 13 Laptop', 
'Ultra-portable Dell XPS 13 with Intel Core i7, 16GB RAM, 512GB SSD, and stunning InfinityEdge display.', 
1299.00, 1499.00, 'DELL-XPS13-I7', 'computers', 
ARRAY['dell', 'xps', 'laptop', 'ultrabook', 'portable'], 
ARRAY['dell-xps-13.jpg'], 
20, 'active', false, 1.27,
'{"length": 29.6, "width": 19.9, "height": 1.48}',
'{"title": "Dell XPS 13 - Ultra-portable Laptop", "description": "Compact and powerful laptop for professionals on the go", "keywords": ["dell", "xps", "laptop", "ultrabook", "portable", "intel"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'Dell Alienware 34 Curved Monitor', 
'34-inch curved gaming monitor with 144Hz refresh rate, NVIDIA G-SYNC, and stunning WQHD resolution.', 
899.00, 1099.00, 'DELL-AW34-144HZ', 'monitors', 
ARRAY['dell', 'alienware', 'monitor', 'gaming', 'curved', '144hz'], 
ARRAY['dell-alienware-34.jpg'], 
10, 'active', true, 8.2,
'{"length": 81.3, "width": 36.3, "height": 24.1}',
'{"title": "Dell Alienware 34 Curved Gaming Monitor", "description": "Premium curved gaming monitor with high refresh rate", "keywords": ["dell", "alienware", "monitor", "gaming", "curved", "144hz", "gsync"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'Apple Watch Ultra', 
'The most rugged and capable Apple Watch, designed for endurance athletes and outdoor adventurers.', 
799.00, 849.00, 'AW-ULTRA-49MM', 'wearables', 
ARRAY['apple', 'watch', 'ultra', 'fitness', 'rugged'], 
ARRAY['apple-watch-ultra.jpg'], 
30, 'active', true, 0.061,
'{"length": 4.9, "width": 4.4, "height": 1.45}',
'{"title": "Apple Watch Ultra - Rugged Smartwatch", "description": "Most advanced Apple Watch for extreme sports and adventures", "keywords": ["apple", "watch", "ultra", "fitness", "rugged", "smartwatch"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'AI Translate Earphones Pro', 
'Revolutionary intelligent translate earphones with real-time translation in 40+ languages. Perfect for travelers, business professionals, and language learners. Features noise cancellation, 8-hour battery life, and instant voice translation.', 
199.00, 249.00, 'AI-TRANSLATE-PRO', 'audio', 
ARRAY['translate', 'earphones', 'ai', 'language', 'travel', 'wireless'], 
ARRAY['ai-translate-earphones.jpg'], 
25, 'active', true, 0.08,
'{"length": 6.5, "width": 4.2, "height": 2.8}',
'{"title": "AI Translate Earphones Pro - Real-time Translation", "description": "Smart earphones with instant translation in 40+ languages", "keywords": ["translate", "earphones", "ai", "language", "travel", "wireless"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'Smart Language Translator Buds', 
'Next-generation wireless earbuds with built-in AI translator. Supports conversation mode, offline translation for 12 languages, and crystal-clear audio quality. Ideal for international business and travel.', 
149.00, 179.00, 'SMART-LANG-BUDS', 'audio', 
ARRAY['translator', 'earbuds', 'smart', 'wireless', 'business', 'travel'], 
ARRAY['smart-translator-buds.jpg'], 
40, 'active', true, 0.06,
'{"length": 5.8, "width": 3.9, "height": 2.4}',
'{"title": "Smart Language Translator Buds - Wireless Translation", "description": "Advanced wireless earbuds with AI-powered real-time translation", "keywords": ["translator", "earbuds", "smart", "wireless", "business", "travel"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'ASUS ROG Rapture GT-BE98 Gaming Router', 
'Quad-band Gaming Router with WiFi 7, advanced QoS, and ultra-low latency for competitive gaming. Features 10Gb ports, RGB lighting, and advanced gaming acceleration.', 
899.99, 999.99, 'ASUS-ROG-GT-BE98', 'networking', 
ARRAY['asus', 'router', 'gaming', 'wifi7', 'networking'], 
ARRAY['asus.jpg'], 
8, 'active', false, 1.2,
'{"length": 26.0, "width": 26.0, "height": 16.8}',
'{"title": "ASUS ROG Gaming Router - WiFi 7", "description": "Professional gaming router with ultra-low latency", "keywords": ["asus", "router", "gaming", "wifi7", "networking"]}'),

(uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440002', 'iPhone 15 Pro Max', 
'The ultimate iPhone with titanium design, A17 Pro chip, and professional camera system. Features 6.7-inch Super Retina XDR display and advanced photography capabilities.', 
1199.99, 1299.99, 'IPHONE-15-PRO-MAX', 'smartphones', 
ARRAY['apple', 'iphone', 'smartphone', 'titanium', 'pro'], 
ARRAY['iphone.jpg'], 
18, 'active', true, 0.221,
'{"length": 15.99, "width": 7.69, "height": 0.83}',
'{"title": "iPhone 15 Pro Max - Premium Smartphone", "description": "Most advanced iPhone with titanium design and A17 Pro chip", "keywords": ["apple", "iphone", "smartphone", "titanium", "pro", "a17"]}');

-- Insert sample customer user
INSERT INTO users (id, email, password_hash, first_name, last_name, role, status) VALUES
('550e8400-e29b-41d4-a716-446655440003', 'customer@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Jane', 'Customer', 'customer', 'active');

-- Insert sample cart
INSERT INTO carts (id, customer_id) VALUES
('550e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440003');

-- Note: Cart items and orders would be inserted through API calls in real usage