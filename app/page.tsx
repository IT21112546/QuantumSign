'use client'

import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"

export default function Home() {
  const router = useRouter()
  const clientId = process.env.NEXT_CLIENT_ID || '26d9b87b-8b44-4bc3-8f10-2b966258b136'
  const apiHost = process.env.NEXT_PUBLIC_SSO_URL || 'localhost:3000'

  const handleLogin = () => {
    router.push(`http://${apiHost}/login?client-id=${clientId}`)
  }

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-4">
      <div className="max-w-md w-full bg-white rounded-lg shadow-md p-8 space-y-6">
        <h1 className="text-2xl font-bold text-center text-gray-800">OAuth 2.0 Client Authentication</h1>
        <p className="text-center text-gray-600">
          This is a testing page for OAuth 2.0 client authentication for the Qsign SSO.
        </p>
        <div className="text-center">
          <p className="font-semibold text-gray-700">Client ID: {clientId}</p>
        </div>
        <Button onClick={handleLogin} className="w-full">
          Login with Qsign
        </Button>
      </div>
    </div>
  )
}
