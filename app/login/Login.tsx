'use client'

import { useSearchParams } from 'next/navigation'
import { jwtDecode } from 'jwt-decode' 
import { useState, useEffect } from 'react'

interface DecodedToken {
  // Add the shape of your JWT here, e.g., based on your token structure
  sub: string;
  name: string;
  iat: number;
  exp: number;
}

export default function Login() {
  const searchParams = useSearchParams()
  const token = searchParams.get('token')

  const [decodedToken, setDecodedToken] = useState<DecodedToken | null>(null)
  const [isAuthenticated, setIsAuthenticated] = useState(false)

  useEffect(() => {
    if (token) {
      try {
        const decoded = jwtDecode<DecodedToken>(token) // Use the DecodedToken interface
        setDecodedToken(decoded)
        setIsAuthenticated(true)
      } catch (error) {
        console.error('Error decoding token:', error)
      }
    }
  }, [token])

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-4">
      <div className="max-w-md w-full bg-white rounded-lg shadow-md p-8 space-y-6">
        <h1 className="text-2xl font-bold text-center text-gray-800">
          {isAuthenticated ? 'User Authenticated' : 'User Not Authenticated'}
        </h1>
        <div className="bg-gray-50 p-4 rounded-md">
          {isAuthenticated ? (
            <pre className="whitespace-pre-wrap break-words text-sm">
              {JSON.stringify(decodedToken, null, 2)}
            </pre>
          ) : (
            <p className="text-center text-gray-600">
              Please include a ?token= parameter in the URL to authenticate.
            </p>
          )}
        </div>
      </div>
    </div>
  )
}
