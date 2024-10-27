'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"

export default function Home() {
  const router = useRouter()
  const [clientId, setClientId] = useState('')
  const apiHost = process.env.NEXT_PUBLIC_SSO_URL || 'http://localhost:3000'

  const handleLogin = () => {
    if (!clientId) {
      alert('Please enter a valid Client ID')
      return
    }
    router.push(`${apiHost}/login?client-id=${clientId}`)
  }

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-4">
      <div className="max-w-md w-full bg-white rounded-lg shadow-md p-8 space-y-6">
        <h1 className="text-2xl font-bold text-center text-gray-800">OAuth 2.0 Client Authentication</h1>
        <p className="text-center text-gray-600">
          This is a testing page for OAuth 2.0 client authentication for the Qsign SSO.
        </p>
        <div className="space-y-4">
          <label htmlFor="clientId" className="block text-sm font-medium text-gray-700">
            Client ID:
          </label>
          <Input
            id="clientId"
            type="text"
            placeholder="Enter your Client ID"
            value={clientId}
            onChange={(e) => setClientId(e.target.value)}
            className="w-full"
          />
        </div>
        <Button onClick={handleLogin} className="w-full">
          Simulate Login with Qsign
        </Button>
      </div>
    </div>
  )
}
