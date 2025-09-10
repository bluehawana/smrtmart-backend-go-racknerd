-- Remove sample data in reverse order of dependencies

-- Remove cart items (if any)
DELETE FROM cart_items WHERE cart_id = '550e8400-e29b-41d4-a716-446655440004';

-- Remove sample cart
DELETE FROM carts WHERE id = '550e8400-e29b-41d4-a716-446655440004';

-- Remove sample products
DELETE FROM products WHERE vendor_id = '550e8400-e29b-41d4-a716-446655440002';

-- Remove sample vendor
DELETE FROM vendors WHERE id = '550e8400-e29b-41d4-a716-446655440002';

-- Remove sample users
DELETE FROM users WHERE email IN ('admin@smrtmart.com', 'vendor@smrtmart.com', 'customer@example.com');

-- Remove sample categories
DELETE FROM categories WHERE slug IN ('electronics', 'computers', 'audio', 'wearables', 'monitors');