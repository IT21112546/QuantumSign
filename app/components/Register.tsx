'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from "@/components/ui/button"; // ShadCN Button component
import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Label } from "@/components/ui/label"; // ShadCN Label component

export default function SignUpForm() {
  // State management for form inputs and messages
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [publicKey, setPublicKey] = useState('');
  const [file, setFile] = useState<File | null>(null);
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const router = useRouter();

  // Access the API host from environment variables
  const API_HOST = process.env.NEXT_PUBLIC_API_HOST || 'localhost:8000';

  // Redirect to '/login' after showing success message
  useEffect(() => {
    if (successMessage) {
      const timer = setTimeout(() => {
        router.push('/login');
      }, 3000); // Redirect after 3 seconds

      return () => clearTimeout(timer);
    }
  }, [successMessage, router]);

  // Handle form submission
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    let kyberPublicKey = publicKey;

    // If a file is uploaded, read its content
    if (file) {
      try {
        const fileContent = await file.text();
        kyberPublicKey = fileContent;
      } catch (err) {
        setErrorMessage('Failed to read the public key file.');
				console.log(err)
        return;
      }
    }

    // Prepare the request body
    const requestBody = {
      kyberPublicKey,
      username,
      email
    };

    try {
      const response = await fetch(`http://${API_HOST}/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
      });

      if (response.ok) {
        const data = await response.json();
        // Display success message and redirect
        setSuccessMessage(data.message || 'Registration successful!');
        // Clear error message if any
        setErrorMessage('');
      } else {
        const errorData = await response.json();
        if (errorData && errorData.error) {
          setErrorMessage(errorData.error);
        } else if (errorData && errorData.message) {
					setErrorMessage(errorData.message);
				}
				else {
          setErrorMessage('Internal Server Error');
        }
        // Clear success message if any
        setSuccessMessage('');
      }
    } catch (error) {
      console.error('Error:', error);
      setErrorMessage('Internal Server Error');
      // Clear success message if any
      setSuccessMessage('');
    }
  };

  // Handle file input change
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setFile(e.target.files[0]);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-r from-gray-100 to-gray-200">
      <div className="bg-white p-10 rounded-3xl shadow-2xl w-full max-w-lg">
        <h2 className="text-3xl font-bold text-center mb-6 text-gray-800">
          SIGN UP
        </h2>

        {/* Display Success Message */}
        {successMessage && (
          <div className="mb-4 text-green-500 text-center font-semibold">
            {successMessage}
          </div>
        )}

        {/* Display Error Message */}
        {errorMessage && (
          <div className="mb-4 text-red-500 text-center font-semibold">
            {errorMessage}
          </div>
        )}

        {/* Form Inputs */}
        <form onSubmit={handleSubmit}>
          <div className="space-y-6">
            {/* Username Field */}
            <div>
              <Label
                htmlFor="username"
                className="block text-gray-700 font-semibold mb-1"
              >
                Enter Your Username
              </Label>
              <Input
                id="username"
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
              />
            </div>

            {/* Email Field */}
            <div>
              <Label
                htmlFor="email"
                className="block text-gray-700 font-semibold mb-1"
              >
                Enter Your Email
              </Label>
              <Input
                id="email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
              />
            </div>

            {/* Public Key Field */}
            <div>
              <Label
                htmlFor="publicKey"
                className="block text-gray-700 font-semibold mb-1"
              >
                Enter Your Public Key
              </Label>
              <Input
                id="publicKey"
                type="text"
                placeholder="Public Key"
                value={publicKey}
                onChange={(e) => setPublicKey(e.target.value)}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
              />
            </div>
          </div>

          {/* File Upload Section */}
          <div className="text-center my-8">
            <p className="text-gray-500 mb-4 text-sm">OR</p>
            <Label className="block font-bold text-gray-700 mb-2">
              UPLOAD YOUR PUBLIC KEY FILE
            </Label>
            <div className="bg-gray-50 shadow-md rounded-xl p-5">
              <input
                type="file"
                accept=".pub"
                onChange={handleFileChange}
              />
            </div>
          </div>

          {/* Register Button */}
          <div className="mt-6 text-center">
            <Button
              type="submit"
              className="bg-black text-white w-full py-3 rounded-full text-lg font-bold hover:bg-gray-900 focus:ring-2 focus:ring-yellow-400"
            >
              REGISTER
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
}

