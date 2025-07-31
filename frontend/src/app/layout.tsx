import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'SmartMart - Premium Electronics & Technology',
  description: 'Discover the latest in technology and premium electronics at SmartMart. Quality products, competitive prices, and exceptional service.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <nav className="bg-white shadow-md">
          <div className="max-w-7xl mx-auto px-4">
            <div className="flex justify-between items-center h-16">
              <div className="flex items-center">
                <a href="/" className="text-xl font-bold text-gray-900">
                  SmartMart
                </a>
              </div>
              <div className="hidden md:block">
                <div className="ml-10 flex items-baseline space-x-4">
                  <a href="/" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Home
                  </a>
                  <a href="/products" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Products
                  </a>
                  <a href="/search" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Search
                  </a>
                  <a href="/about" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    About
                  </a>
                  <a href="/contact" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Contact
                  </a>
                  <a href="/support" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Support
                  </a>
                  <a href="/profile" className="text-gray-900 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                    Account
                  </a>
                </div>
              </div>
            </div>
          </div>
        </nav>
        <main>{children}</main>
        <footer className="bg-gray-800 text-white py-8 mt-12">
          <div className="max-w-7xl mx-auto px-4">
            <div className="text-center">
              <p>&copy; 2024 SmartMart. All rights reserved.</p>
            </div>
          </div>
        </footer>
      </body>
    </html>
  )
}