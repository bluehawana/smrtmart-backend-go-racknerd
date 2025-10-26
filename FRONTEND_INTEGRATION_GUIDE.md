# Frontend Integration Guide - SmrtMart API

## 🎯 Overview

This guide shows how to integrate your Next.js frontend (Vercel) with the SmrtMart Go backend API.

## 📡 API Base URLs

### Production
```javascript
const API_BASE_URL = 'https://api.smrtmart.com/api/v1';
```

### Development/Local
```javascript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

## 🔐 Environment Variables

Add to your frontend `.env.local`:

```bash
# API Configuration
NEXT_PUBLIC_API_URL=https://api.smrtmart.com/api/v1

# Stripe (Frontend)
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_live_YOUR_PUBLISHABLE_KEY
```

## 📋 Available API Endpoints

### Public Endpoints (No Auth Required)

#### 1. Health Check
```javascript
GET /health

// Example
const response = await fetch(`${API_BASE_URL}/health`);
const data = await response.json();
// Response: { "status": "healthy", "service": "SmrtMart API v1", "version": "1.0.0" }
```

#### 2. Products
```javascript
// Get all products
GET /products?page=1&limit=20&category_id=1

// Get single product
GET /products/:id

// Search products
GET /products/search?q=keyword&category_id=1

// Get featured products
GET /products/featured

// Example:
const getProducts = async (page = 1, limit = 20) => {
  const response = await fetch(
    `${API_BASE_URL}/products?page=${page}&limit=${limit}`
  );
  return await response.json();
};
```

#### 3. Categories
```javascript
// Get all categories
GET /categories

// Get single category
GET /categories/:id

// Example:
const getCategories = async () => {
  const response = await fetch(`${API_BASE_URL}/categories`);
  return await response.json();
};
```

#### 4. Authentication
```javascript
// Register
POST /auth/register
Body: {
  "email": "user@example.com",
  "password": "securepassword",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890"
}

// Login
POST /auth/login
Body: {
  "email": "user@example.com",
  "password": "securepassword"
}
// Response: { "token": "jwt_token", "user": {...} }

// Refresh token
POST /auth/refresh
Headers: { "Authorization": "Bearer <refresh_token>" }

// Example:
const login = async (email, password) => {
  const response = await fetch(`${API_BASE_URL}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  return await response.json();
};
```

#### 5. Cart (Guest & Authenticated)
```javascript
// Get cart
GET /cart
Headers: { "Authorization": "Bearer <token>" } // Optional

// Add item to cart
POST /cart/items
Body: {
  "product_id": "123",
  "quantity": 2
}

// Update cart item
PUT /cart/items/:id
Body: {
  "quantity": 3
}

// Remove item
DELETE /cart/items/:id

// Clear cart
DELETE /cart
// or
POST /cart/clear

// Example:
const addToCart = async (productId, quantity, token) => {
  const headers = {
    'Content-Type': 'application/json'
  };
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const response = await fetch(`${API_BASE_URL}/cart/items`, {
    method: 'POST',
    headers,
    body: JSON.stringify({ product_id: productId, quantity })
  });
  return await response.json();
};
```

#### 6. Checkout (Stripe)
```javascript
POST /orders/checkout
Body: {
  "items": [
    {
      "product_id": "123",
      "name": "Product Name",
      "description": "Product description",
      "price": 29.99,
      "quantity": 2,
      "images": ["image1.jpg"]
    }
  ],
  "customer_email": "customer@example.com",
  "success_url": "https://smrtmart.com/checkout/success?session_id={CHECKOUT_SESSION_ID}",
  "cancel_url": "https://smrtmart.com/checkout/cancel"
}

// Response:
{
  "success": true,
  "message": "Checkout session created successfully",
  "data": {
    "session_id": "cs_live_...",
    "session_url": "https://checkout.stripe.com/..."
  }
}

// Example:
const createCheckout = async (cartItems, customerEmail) => {
  const response = await fetch(`${API_BASE_URL}/orders/checkout`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      items: cartItems,
      customer_email: customerEmail,
      success_url: `${window.location.origin}/checkout/success?session_id={CHECKOUT_SESSION_ID}`,
      cancel_url: `${window.location.origin}/checkout/cancel`
    })
  });

  const result = await response.json();

  if (result.success) {
    // Redirect to Stripe Checkout
    window.location.href = result.data.session_url;
  }

  return result;
};
```

### Protected Endpoints (Auth Required)

All protected endpoints require JWT token in Authorization header:
```javascript
Headers: {
  "Authorization": "Bearer <jwt_token>"
}
```

#### 7. User Profile
```javascript
// Get profile
GET /users/profile

// Update profile
PUT /users/profile
Body: {
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890"
}

// Change password
POST /users/change-password
Body: {
  "old_password": "oldpass",
  "new_password": "newpass"
}
```

#### 8. Orders
```javascript
// Get user orders
GET /orders?page=1&limit=10

// Get specific order
GET /orders/:id

// Cancel order
POST /orders/:id/cancel
```

#### 9. Reviews
```javascript
// Create review
POST /reviews
Body: {
  "product_id": "123",
  "rating": 5,
  "comment": "Great product!"
}

// Update review
PUT /reviews/:id

// Delete review
DELETE /reviews/:id
```

## 🛒 Complete Checkout Flow Example

```javascript
// components/CheckoutButton.jsx
import { useState } from 'react';

export default function CheckoutButton({ cart, userEmail }) {
  const [loading, setLoading] = useState(false);

  const handleCheckout = async () => {
    setLoading(true);

    try {
      // Format cart items for API
      const items = cart.map(item => ({
        product_id: item.id.toString(),
        name: item.name,
        description: item.description || '',
        price: parseFloat(item.price),
        quantity: item.quantity,
        images: item.images || []
      }));

      // Create checkout session
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/orders/checkout`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            items,
            customer_email: userEmail,
            success_url: `${window.location.origin}/checkout/success?session_id={CHECKOUT_SESSION_ID}`,
            cancel_url: `${window.location.origin}/checkout/cancel`
          })
        }
      );

      const result = await response.json();

      if (result.success) {
        // Redirect to Stripe Checkout
        window.location.href = result.data.session_url;
      } else {
        alert(`Error: ${result.message}`);
      }
    } catch (error) {
      console.error('Checkout error:', error);
      alert('Failed to create checkout session');
    } finally {
      setLoading(false);
    }
  };

  return (
    <button
      onClick={handleCheckout}
      disabled={loading || cart.length === 0}
      className="bg-blue-600 text-white px-6 py-3 rounded-lg disabled:opacity-50"
    >
      {loading ? 'Creating checkout...' : 'Proceed to Checkout'}
    </button>
  );
}
```

## 🔄 API Client Helper

Create a reusable API client:

```javascript
// lib/api.js
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL;

class APIClient {
  constructor() {
    this.baseURL = API_BASE_URL;
  }

  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`;
    const config = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    };

    // Add auth token if available
    const token = localStorage.getItem('auth_token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    const response = await fetch(url, config);
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'API request failed');
    }

    return data;
  }

  // Products
  async getProducts(params = {}) {
    const query = new URLSearchParams(params).toString();
    return this.request(`/products${query ? `?${query}` : ''}`);
  }

  async getProduct(id) {
    return this.request(`/products/${id}`);
  }

  async searchProducts(query, categoryId) {
    const params = new URLSearchParams({ q: query });
    if (categoryId) params.append('category_id', categoryId);
    return this.request(`/products/search?${params}`);
  }

  // Categories
  async getCategories() {
    return this.request('/categories');
  }

  // Auth
  async login(email, password) {
    const data = await this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password })
    });

    if (data.data?.token) {
      localStorage.setItem('auth_token', data.data.token);
    }

    return data;
  }

  async register(userData) {
    return this.request('/auth/register', {
      method: 'POST',
      body: JSON.stringify(userData)
    });
  }

  logout() {
    localStorage.removeItem('auth_token');
  }

  // Cart
  async getCart() {
    return this.request('/cart');
  }

  async addToCart(productId, quantity) {
    return this.request('/cart/items', {
      method: 'POST',
      body: JSON.stringify({ product_id: productId, quantity })
    });
  }

  async updateCartItem(itemId, quantity) {
    return this.request(`/cart/items/${itemId}`, {
      method: 'PUT',
      body: JSON.stringify({ quantity })
    });
  }

  async removeFromCart(itemId) {
    return this.request(`/cart/items/${itemId}`, {
      method: 'DELETE'
    });
  }

  async clearCart() {
    return this.request('/cart', {
      method: 'DELETE'
    });
  }

  // Checkout
  async createCheckout(items, customerEmail, successUrl, cancelUrl) {
    return this.request('/orders/checkout', {
      method: 'POST',
      body: JSON.stringify({
        items,
        customer_email: customerEmail,
        success_url: successUrl,
        cancel_url: cancelUrl
      })
    });
  }

  // Orders
  async getOrders(page = 1, limit = 10) {
    return this.request(`/orders?page=${page}&limit=${limit}`);
  }

  async getOrder(id) {
    return this.request(`/orders/${id}`);
  }

  // User
  async getProfile() {
    return this.request('/users/profile');
  }

  async updateProfile(userData) {
    return this.request('/users/profile', {
      method: 'PUT',
      body: JSON.stringify(userData)
    });
  }
}

export const api = new APIClient();
```

## 🎨 Usage in Components

```javascript
// pages/products/index.jsx
import { useEffect, useState } from 'react';
import { api } from '@/lib/api';

export default function ProductsPage() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadProducts();
  }, []);

  const loadProducts = async () => {
    try {
      const data = await api.getProducts({ page: 1, limit: 20 });
      setProducts(data.data || []);
    } catch (error) {
      console.error('Failed to load products:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div>Loading...</div>;

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
      {products.map(product => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  );
}
```

## 🔒 CORS Configuration

The backend already has CORS configured. Make sure your frontend domain is allowed:

**Backend .env:**
```bash
CORS_ORIGINS=https://smrtmart.com,https://www.smrtmart.com
```

## 🧪 Testing API Calls

```javascript
// Test API connectivity
const testAPI = async () => {
  try {
    const response = await fetch('https://api.smrtmart.com/api/v1/health');
    const data = await response.json();
    console.log('API Health:', data);
    // Expected: { "status": "healthy", ... }
  } catch (error) {
    console.error('API Error:', error);
  }
};
```

## 📊 API Response Format

All API responses follow this format:

```javascript
// Success
{
  "success": true,
  "message": "Operation successful",
  "data": { /* actual data */ }
}

// Error
{
  "success": false,
  "message": "Error description",
  "error": {
    "code": "ERROR_CODE",
    "message": "Detailed error message"
  }
}
```

## 🚨 Error Handling

```javascript
const handleAPICall = async (apiFunction) => {
  try {
    const result = await apiFunction();

    if (!result.success) {
      throw new Error(result.message || 'Operation failed');
    }

    return result.data;
  } catch (error) {
    console.error('API Error:', error);

    // Handle specific errors
    if (error.message.includes('401') || error.message.includes('Unauthorized')) {
      // Redirect to login
      window.location.href = '/login';
    }

    throw error;
  }
};
```

## 🔗 Useful Links

- **API Base URL**: https://api.smrtmart.com/api/v1
- **API Health Check**: https://api.smrtmart.com/api/v1/health
- **Swagger Docs**: https://api.smrtmart.com/swagger/index.html (if enabled)

## 📝 Next Steps

1. ✅ Update frontend `.env.local` with API URL
2. ✅ Create API client helper (`lib/api.js`)
3. ✅ Implement authentication flow
4. ✅ Integrate product listing
5. ✅ Implement cart functionality
6. ✅ Set up Stripe checkout flow
7. ✅ Add error handling and loading states
8. ✅ Test end-to-end checkout process

---

**Questions or issues?** Check the backend logs:
```bash
ssh harvad@107.175.235.220 'sudo journalctl -u smrtmart -f'
```
