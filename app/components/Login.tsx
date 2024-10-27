'use client'

import { useState, useRef } from 'react'
import { useRouter } from 'next/navigation'
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Loader2, Upload } from 'lucide-react'

interface LoginProps {
  clientId: string
}

export default function Login({ clientId }: LoginProps) {
  const [publicKey, setPublicKey] = useState('')
  const [file, setFile] = useState<File | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)
  const router = useRouter()

  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (file && file.name.endsWith('.pub')) {
      setFile(file)
      const reader = new FileReader()
      reader.onload = (e) => {
        const content = e.target?.result as string
        setPublicKey(content.trim())
      }
      reader.readAsText(file)
    } else {
      alert('Please upload a valid .pub file')
      if (fileInputRef.current) {
        fileInputRef.current.value = ''
      }
    }
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()
    setIsLoading(true)

    const apiUrl = `${process.env.NEXT_PUBLIC_API_HOST || 'http://localhost:8000'}/login`
    
    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          kyberPublicKey: publicKey,
          clientId: clientId,
        }),
      })

      if (!response.ok) {
        throw new Error('Login failed')
      }

      const data = await response.json()
      router.push(`${data.redirectUrl}?token=${data.accessToken}`)
    } catch (error) {
      console.error('Login error:', error)
      alert('Login failed. Please try again.')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white shadow-lg rounded-2xl p-8 max-w-md w-full">
        <h2 className="text-3xl font-bold text-gray-800 text-center mb-8">
          Sign In
        </h2>

        <form onSubmit={handleSubmit}>
          <div className="space-y-6">
            <div>
              <Label htmlFor="publicKey" className="block text-gray-700 font-semibold mb-1">
                Enter Your Public Key
              </Label>
              <Input
                id="publicKey"
                type="text"
                placeholder="Enter Your Public Key"
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400 focus:outline-none transition duration-200"
                value={publicKey}
                onChange={(e) => setPublicKey(e.target.value)}
              />
            </div>

            <div className="relative my-6">
              <span className="absolute inset-x-0 top-1/2 transform -translate-y-1/2 text-center text-gray-400 bg-white px-2">
                OR
              </span>
              <div className="border-t border-gray-300"></div>
            </div>

            <div className="text-center">
              <Label htmlFor="file-upload" className="block text-lg font-semibold text-gray-700 mb-4">
                Upload Your Public Key File
              </Label>
              <div className="bg-gray-50 rounded-lg p-6 shadow-md flex justify-center">
                <input
                  id="file-upload"
                  type="file"
                  className="hidden"
                  onChange={handleFileUpload}
                  accept=".pub"
                  ref={fileInputRef}
                />
                <Button
                  type="button"
                  onClick={() => fileInputRef.current?.click()}
                  className="bg-blue-500 text-white px-6 py-3 rounded-full hover:bg-blue-600 transition-all duration-300 flex items-center justify-center"
                >
                  <Upload className="mr-2 h-5 w-5" />
                  Upload File
                </Button>
              </div>
              {file && (
                <p className="mt-2 text-sm text-gray-600">
                  File uploaded: {file.name}
                </p>
              )}
            </div>

            <div className="flex justify-center">
              <Button
                type="submit"
                className="bg-black text-white px-8 py-3 rounded-full text-lg font-semibold hover:bg-gray-900 transition-all duration-300 flex items-center justify-center"
                disabled={isLoading || !publicKey}
              >
                {isLoading ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Logging in...
                  </>
                ) : (
                  'Login'
                )}
              </Button>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}
