'use client';

import { useState, useEffect } from 'react';
import Image from 'next/image';
import Link from 'next/link';

// Product Image Component with error handling
function ProductImage({ src, alt, className }: { src: string; alt: string; className?: string }) {
  const [imgSrc, setImgSrc] = useState(src);
  const [hasError, setHasError] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  const handleError = () => {
    if (!hasError) {
      setHasError(true);
      setIsLoading(false);
      console.log('Image failed to load:', src);
    }
  };

  const handleLoad = () => {
    setIsLoading(false);
  };

  // If no src provided or has error, show placeholder
  if (!src || hasError) {
    return (
      <div className={`bg-gray-200 flex items-center justify-center ${className}`}>
        <div className="text-center text-gray-500">
          <div className="text-4xl mb-2">📦</div>
          <div className="text-xs">{alt.split(' ')[0]}</div>
        </div>
      </div>
    );
  }

  return (
    <div className="relative w-full h-full">
      {isLoading && (
        <div className="absolute inset-0 bg-gray-200 animate-pulse flex items-center justify-center">
          <div className="text-gray-400 text-4xl">📦</div>
        </div>
      )}
      <img
        src={imgSrc}
        alt={alt}
        className={className}
        onLoad={handleLoad}
        onError={handleError}
        style={{ display: hasError ? 'none' : 'block' }}
      />
    </div>
  );
}

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'https://smrtmart-go-backend-1753976056-b4c4ef7e5ab7.herokuapp.com/api/v1';

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  compare_price: number;
  images: string[];
  stock: number;
  status: string;
  featured: boolean;
  category: string;
  tags: string[];
}

export default function HomePage() {
  const [products, setProducts] = useState<Product[]>([]);
  const [featuredProducts, setFeaturedProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const [productsRes, featuredRes] = await Promise.all([
        fetch(`${BASE_URL}/products?status=active&limit=6`),
        fetch(`${BASE_URL}/products/featured?limit=6`)
      ]);

      if (productsRes.ok) {
        const productsData = await productsRes.json();
        if (productsData.success) {
          setProducts(productsData.data.data);
        }
      }

      if (featuredRes.ok) {
        const featuredData = await featuredRes.json();
        if (featuredData.success) {
          setFeaturedProducts(featuredData.data);
        }
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[...Array(6)].map((_, i) => (
            <div key={i} className="bg-white rounded-lg shadow-md overflow-hidden animate-pulse">
              <div className="h-64 bg-gray-200" />
              <div className="p-4">
                <div className="h-4 bg-gray-200 rounded w-3/4 mb-2" />
                <div className="h-4 bg-gray-200 rounded w-1/4" />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 py-8">
      {/* Hero Section */}
      <div className="text-center mb-12">
        <h1 className="text-4xl font-bold text-gray-900 mb-4">
          Welcome to SmartMart
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          Discover the latest in technology and premium electronics
        </p>
        <Link
          href="/products"
          className="inline-flex px-8 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 font-medium"
        >
          Shop Now
        </Link>
      </div>

      {/* Featured Products */}
      <section className="mb-12">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-900">Featured Products</h2>
          <Link href="/products" className="text-blue-600 hover:underline">
            View all
          </Link>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {featuredProducts.slice(0, 6).map((product) => (
            <Link
              href={`/products/${product.id}`}
              key={product.id}
              className="group bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
            >
              <div className="relative w-full pt-[75%]">
                <ProductImage
                  src={product.images[0] || ''}
                  alt={product.name}
                  className="absolute inset-0 w-full h-full object-cover group-hover:opacity-75 transition-opacity duration-300"
                />
                {product.featured && (
                  <div className="absolute top-2 left-2 bg-red-500 text-white px-2 py-1 text-xs rounded">
                    Featured
                  </div>
                )}
              </div>
              <div className="p-4">
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  {product.name}
                </h3>
                <p className="text-sm text-gray-600 mb-2 line-clamp-2">
                  {product.description}
                </p>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <span className="text-xl font-bold text-gray-900">
                      ${product.price.toFixed(2)}
                    </span>
                    {product.compare_price > product.price && (
                      <span className="text-sm text-gray-500 line-through">
                        ${product.compare_price.toFixed(2)}
                      </span>
                    )}
                  </div>
                  <span className="text-sm text-gray-500">
                    {product.stock > 0 ? 'In Stock' : 'Out of Stock'}
                  </span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      </section>

      {/* Latest Products */}
      <section>
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-900">Latest Products</h2>
          <Link href="/products" className="text-blue-600 hover:underline">
            View all
          </Link>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {products.slice(0, 6).map((product) => (
            <Link
              href={`/products/${product.id}`}
              key={product.id}
              className="group bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
            >
              <div className="relative w-full pt-[75%]">
                <ProductImage
                  src={product.images[0] || ''}
                  alt={product.name}
                  className="absolute inset-0 w-full h-full object-cover group-hover:opacity-75 transition-opacity duration-300"
                />
              </div>
              <div className="p-4">
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  {product.name}
                </h3>
                <p className="text-sm text-gray-600 mb-2 line-clamp-2">
                  {product.description}
                </p>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <span className="text-xl font-bold text-gray-900">
                      ${product.price.toFixed(2)}
                    </span>
                    {product.compare_price > product.price && (
                      <span className="text-sm text-gray-500 line-through">
                        ${product.compare_price.toFixed(2)}
                      </span>
                    )}
                  </div>
                  <span className="text-sm text-gray-500">
                    {product.stock > 0 ? 'In Stock' : 'Out of Stock'}
                  </span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      </section>
    </div>
  );
}