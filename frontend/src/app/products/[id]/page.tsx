'use client';

import { useState, useEffect } from 'react';
import Image from 'next/image';
import { useParams } from 'next/navigation';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'https://smrtmart-go-backend-1753976056-b4c4ef7e5ab7.herokuapp.com/api/v1';

interface Product {
  id: string;
  vendor_id: string;
  name: string;
  description: string;
  price: number;
  compare_price: number;
  sku: string;
  category: string;
  tags: string[];
  images: string[];
  stock: number;
  status: string;
  featured: boolean;
  weight: number;
  dimensions: {
    length: number;
    width: number;
    height: number;
  };
  seo: {
    title: string;
    description: string;
    keywords: string[];
  };
  created_at: string;
  updated_at: string;
}

export default function ProductDetailPage() {
  const params = useParams();
  const productId = params.id as string;
  
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedImage, setSelectedImage] = useState(0);
  const [quantity, setQuantity] = useState(1);

  useEffect(() => {
    if (productId) {
      fetchProduct();
    }
  }, [productId]);

  const fetchProduct = async () => {
    try {
      setLoading(true);
      const response = await fetch(`${BASE_URL}/products/${productId}`);
      
      if (response.ok) {
        const data = await response.json();
        if (data.success) {
          setProduct(data.data);
        } else {
          setError('Product not found');
        }
      } else {
        setError('Product not found');
      }
    } catch (error) {
      setError('Failed to load product');
      console.error('Error fetching product:', error);
    } finally {
      setLoading(false);
    }
  };

  const addToCart = () => {
    // Placeholder for add to cart functionality
    alert(`Added ${quantity} ${product?.name} to cart!`);
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-8">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div className="animate-pulse">
            <div className="aspect-square bg-gray-200 rounded-lg mb-4"></div>
            <div className="grid grid-cols-4 gap-2">
              {[...Array(4)].map((_, i) => (
                <div key={i} className="aspect-square bg-gray-200 rounded"></div>
              ))}
            </div>
          </div>
          <div className="animate-pulse">
            <div className="h-8 bg-gray-200 rounded w-3/4 mb-4"></div>
            <div className="h-6 bg-gray-200 rounded w-1/4 mb-4"></div>
            <div className="h-24 bg-gray-200 rounded mb-6"></div>
            <div className="h-12 bg-gray-200 rounded w-1/2"></div>
          </div>
        </div>
      </div>
    );
  }

  if (error || !product) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-16 text-center">
        <h1 className="text-2xl font-bold text-gray-900 mb-4">Product Not Found</h1>
        <p className="text-gray-600 mb-8">{error || 'The product you are looking for does not exist.'}</p>
        <a
          href="/products"
          className="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          Browse All Products
        </a>
      </div>
    );
  }

  const savings = product.compare_price > product.price ? product.compare_price - product.price : 0;
  const savingsPercent = savings > 0 ? Math.round((savings / product.compare_price) * 100) : 0;

  return (
    <div className="max-w-7xl mx-auto px-4 py-8">
      {/* Breadcrumb */}
      <nav className="mb-8">
        <ol className="flex items-center space-x-2 text-sm text-gray-500">
          <li><a href="/" className="hover:text-blue-600">Home</a></li>
          <li>/</li>
          <li><a href="/products" className="hover:text-blue-600">Products</a></li>
          <li>/</li>
          <li className="capitalize"><a href={`/products?category=${product.category}`} className="hover:text-blue-600">{product.category}</a></li>
          <li>/</li>
          <li className="text-gray-900 truncate">{product.name}</li>
        </ol>
      </nav>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
        {/* Product Images */}
        <div>
          <div className="aspect-square mb-6 bg-gray-100 rounded-lg overflow-hidden">
            <Image
              src={product.images[selectedImage] || '/placeholder-product.jpg'}
              alt={product.name}
              width={600}
              height={600}
              className="w-full h-full object-cover"
              priority
            />
          </div>
          
          {product.images.length > 1 && (
            <div className="grid grid-cols-4 gap-2">
              {product.images.map((image, index) => (
                <button
                  key={index}
                  onClick={() => setSelectedImage(index)}
                  className={`aspect-square rounded-lg overflow-hidden border-2 ${
                    selectedImage === index ? 'border-blue-500' : 'border-gray-200'
                  }`}
                >
                  <Image
                    src={image}
                    alt={`${product.name} view ${index + 1}`}
                    width={150}
                    height={150}
                    className="w-full h-full object-cover"
                  />
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Product Info */}
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-4">{product.name}</h1>
          
          {/* Price */}
          <div className="mb-6">
            <div className="flex items-center space-x-4 mb-2">
              <span className="text-3xl font-bold text-gray-900">
                ${product.price.toFixed(2)}
              </span>
              {product.compare_price > product.price && (
                <span className="text-xl text-gray-500 line-through">
                  ${product.compare_price.toFixed(2)}
                </span>
              )}
            </div>
            {savings > 0 && (
              <div className="text-green-600 font-medium">
                Save ${savings.toFixed(2)} ({savingsPercent}% off)
              </div>
            )}
          </div>

          {/* Stock Status */}
          <div className="mb-6">
            {product.stock > 0 ? (
              <span className="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800">
                ✓ In Stock ({product.stock} available)
              </span>
            ) : (
              <span className="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-red-100 text-red-800">
                Out of Stock
              </span>
            )}
          </div>

          {/* Description */}
          <div className="mb-8">
            <h3 className="text-lg font-semibold text-gray-900 mb-3">Description</h3>
            <p className="text-gray-600 leading-relaxed">{product.description}</p>
          </div>

          {/* Product Details */}
          <div className="mb-8 space-y-4">
            <div className="grid grid-cols-2 gap-4 text-sm">
              <div>
                <span className="font-medium text-gray-900">SKU:</span>
                <span className="ml-2 text-gray-600">{product.sku}</span>
              </div>
              <div>
                <span className="font-medium text-gray-900">Category:</span>
                <span className="ml-2 text-gray-600 capitalize">{product.category}</span>
              </div>
              {product.weight && (
                <div>
                  <span className="font-medium text-gray-900">Weight:</span>
                  <span className="ml-2 text-gray-600">{product.weight} kg</span>
                </div>
              )}
              {product.dimensions && (
                <div>
                  <span className="font-medium text-gray-900">Dimensions:</span>
                  <span className="ml-2 text-gray-600">
                    {product.dimensions.length} × {product.dimensions.width} × {product.dimensions.height} cm
                  </span>
                </div>
              )}
            </div>
          </div>

          {/* Tags */}
          {product.tags && product.tags.length > 0 && (
            <div className="mb-8">
              <h3 className="text-lg font-semibold text-gray-900 mb-3">Tags</h3>
              <div className="flex flex-wrap gap-2">
                {product.tags.map((tag, index) => (
                  <span
                    key={index}
                    className="px-3 py-1 bg-gray-100 text-gray-700 text-sm rounded-full"
                  >
                    {tag}
                  </span>
                ))}
              </div>
            </div>
          )}

          {/* Add to Cart */}
          {product.stock > 0 && (
            <div className="border-t pt-8">
              <div className="flex items-center space-x-4 mb-6">
                <label htmlFor="quantity" className="font-medium text-gray-900">
                  Quantity:
                </label>
                <select
                  id="quantity"
                  value={quantity}
                  onChange={(e) => setQuantity(parseInt(e.target.value))}
                  className="border border-gray-300 rounded-md px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  {[...Array(Math.min(product.stock, 10))].map((_, i) => (
                    <option key={i + 1} value={i + 1}>
                      {i + 1}
                    </option>
                  ))}
                </select>
              </div>
              
              <div className="flex space-x-4">
                <button
                  onClick={addToCart}
                  className="flex-1 bg-blue-600 text-white py-3 px-6 rounded-lg hover:bg-blue-700 font-medium transition-colors"
                >
                  Add to Cart
                </button>
                <button className="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
                  ♡ Wishlist
                </button>
              </div>
            </div>
          )}
        </div>
      </div>

      {/* Additional Info Tabs */}
      <div className="mt-16">
        <div className="border-b border-gray-200">
          <nav className="-mb-px flex space-x-8">
            <button className="border-b-2 border-blue-500 text-blue-600 py-2 px-1 font-medium text-sm">
              Specifications
            </button>
          </nav>
        </div>
        <div className="py-8">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div>
              <h4 className="font-semibold text-gray-900 mb-4">Product Information</h4>
              <dl className="space-y-2">
                <div className="flex">
                  <dt className="font-medium text-gray-900 w-24">SKU:</dt>
                  <dd className="text-gray-600">{product.sku}</dd>
                </div>
                <div className="flex">
                  <dt className="font-medium text-gray-900 w-24">Category:</dt>
                  <dd className="text-gray-600 capitalize">{product.category}</dd>
                </div>
                {product.weight && (
                  <div className="flex">
                    <dt className="font-medium text-gray-900 w-24">Weight:</dt>
                    <dd className="text-gray-600">{product.weight} kg</dd>
                  </div>
                )}
              </dl>
            </div>
            {product.dimensions && (
              <div>
                <h4 className="font-semibold text-gray-900 mb-4">Dimensions</h4>
                <dl className="space-y-2">
                  <div className="flex">
                    <dt className="font-medium text-gray-900 w-24">Length:</dt>
                    <dd className="text-gray-600">{product.dimensions.length} cm</dd>
                  </div>
                  <div className="flex">
                    <dt className="font-medium text-gray-900 w-24">Width:</dt>
                    <dd className="text-gray-600">{product.dimensions.width} cm</dd>
                  </div>
                  <div className="flex">
                    <dt className="font-medium text-gray-900 w-24">Height:</dt>
                    <dd className="text-gray-600">{product.dimensions.height} cm</dd>
                  </div>
                </dl>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}