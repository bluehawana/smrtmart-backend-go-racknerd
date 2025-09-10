-- MySQL version of seed data down migration

-- Delete in reverse order of dependencies
DELETE FROM products WHERE vendor_id = '550e8400-e29b-41d4-a716-446655440002';
DELETE FROM vendors WHERE id = '550e8400-e29b-41d4-a716-446655440002';
DELETE FROM users WHERE id = '550e8400-e29b-41d4-a716-446655440001';
DELETE FROM users WHERE email = 'admin@smrtmart.com';
DELETE FROM categories WHERE slug IN ('electronics', 'computers', 'audio', 'wearables', 'monitors', 'networking', 'smartphones');