'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Menu, X } from 'lucide-react'

const navItems = [
  { name: 'Home', href: '/' },
  { name: 'Pricing', href: '/pricing' },
  { name: 'Register Application', href: '/register-client' },
  { name: 'Generate Keypair', href: '/gen-keypair' },
  { name: 'Our Research', href: 'https://research.qsign.io', external: true },
]

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <nav className="bg-gray-900 text-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex-shrink-0">
            <Link href="/" className="text-xl font-bold">
              QSIGN
            </Link>
          </div>
          <div className="hidden md:block">
            <div className="ml-10 flex items-center space-x-4">
              {navItems.map((item) => (
                <Link
                  key={item.name}
                  href={item.href}
                  className="text-gray-300 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                  {...(item.external ? { target: '_blank', rel: 'noopener noreferrer' } : {})}
                >
                  {item.name}
                </Link>
              ))}
            </div>
          </div>
          <div className="hidden md:flex items-center space-x-2">
            <Link href="/register">
              <button className="bg-transparent text-white border border-white px-4 py-2 rounded-md text-sm font-medium hover:bg-white hover:text-gray-900 transition-colors">
                Register
              </button>
            </Link>
            <Link href="/login">
              <button className="bg-teal-500 text-white px-4 py-2 rounded-md text-sm font-medium hover:bg-teal-600 transition-colors">
                Simulate a Login
              </button>
            </Link>
          </div>
          <div className="md:hidden">
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
            >
              <span className="sr-only">Open main menu</span>
              {isMenuOpen ? (
                <X className="block h-6 w-6" aria-hidden="true" />
              ) : (
                <Menu className="block h-6 w-6" aria-hidden="true" />
              )}
            </button>
          </div>
        </div>
      </div>

      {isMenuOpen && (
        <div className="md:hidden">
          <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
            {navItems.map((item) => (
              <Link
                key={item.name}
                href={item.href}
                className="text-gray-300 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
                {...(item.external ? { target: '_blank', rel: 'noopener noreferrer' } : {})}
              >
                {item.name}
              </Link>
            ))}
            <Link href="/register">
              <button className="w-full text-left bg-transparent text-white border border-white px-3 py-2 rounded-md text-base font-medium hover:bg-white hover:text-gray-900 transition-colors mt-2">
                Register
              </button>
            </Link>
            <Link href="/login">
              <button className="w-full text-left bg-teal-500 text-white px-3 py-2 rounded-md text-base font-medium hover:bg-teal-600 transition-colors mt-2">
                Login
              </button>
            </Link>
          </div>
        </div>
      )}
    </nav>
  )
}
