'use client';

import { useState, useEffect } from 'react';
import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Button } from "@/components/ui/button"; // ShadCN Button component
import { Label } from "@/components/ui/label"; // ShadCN Label component
import { ClipboardIcon } from "@heroicons/react/24/outline"; // For clipboard icon

export default function ClientRegistrationForm() {
  // State variables for form inputs and response data
  const [clientName, setClientName] = useState('');
  const [redirectUrl, setRedirectUrl] = useState('');
  const [clientId, setClientId] = useState('');
  const [publicKey, setPublicKey] = useState('');
  const [privateKey, setPrivateKey] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [redirectUrlError, setRedirectUrlError] = useState('');

  // Access the API host from environment variables or use default
  const API_HOST = process.env.NEXT_PUBLIC_API_HOST || 'localhost:8000';

  // Function to validate URL
  const validateUrl = (url: string): boolean => {
    try {
      new URL(url);
      return true;
    } catch {
      return false;
    }
  };

  // Handle form submission
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Reset error messages
    setErrorMessage('');
    setRedirectUrlError('');

    // Validate Redirect URL
    if (!validateUrl(redirectUrl)) {
      setRedirectUrlError('Please enter a valid URL.');
      return;
    }

    // Prepare the request body
    const requestBody = {
      clientName,
      redirectUrl,
    };

    try {
      const response = await fetch(`http://${API_HOST}/client-registration`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody),
      });

      if (response.ok) {
        const data = await response.json();
        // Update state with the returned data
        setClientId(data.clientId);
        setPublicKey(data.publicKey);
        setPrivateKey(data.privateKey);
        // Clear any previous error messages
        setErrorMessage('');
      } else {
        // Handle error response
        const errorData = await response.json();
        if (errorData && errorData.error) {
          setErrorMessage(errorData.error);
        } else {
          setErrorMessage('Internal Server Error');
        }
      }
    } catch (error) {
      console.error('Error:', error);
      setErrorMessage('Internal Server Error');
    }
  };

  // Function to handle copying text to the clipboard
  const handleCopy = (text: string) => {
    if (text) {
      navigator.clipboard.writeText(text).then(() => {
        // Optionally, provide feedback that the copy was successful
        // e.g., using a temporary message or tooltip
        alert('Copied to clipboard!');
      }).catch((err) => {
        console.error('Failed to copy: ', err);
      });
    }
  };

  // Real-time validation for Redirect URL
  useEffect(() => {
    if (redirectUrl === '') {
      setRedirectUrlError('');
      return;
    }
    if (validateUrl(redirectUrl)) {
      setRedirectUrlError('');
    } else {
      setRedirectUrlError('Please enter a valid URL.');
    }
  }, [redirectUrl]);

  return (
    <div className="bg-gradient-to-r from-gray-100 to-gray-200 py-12 px-6">
      {/* Removed min-h-screen and replaced with py-12 for top/bottom padding */}
      <div className="bg-white p-10 rounded-3xl shadow-2xl max-w-4xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-12">
          {/* Left Column: Form */}
          <div>
            <h1 className="text-4xl font-bold text-black mb-6">
              Register Your Client Application
            </h1>
            <p className="text-gray-600 mb-8 text-lg">
              Enter your application&apos;s name and its redirect URL.
            </p>

            {/* Display Error Message */}
            {errorMessage && (
              <div className="mb-4 text-red-500 text-center font-semibold">
                {errorMessage}
              </div>
            )}

            <form className="space-y-6" onSubmit={handleSubmit} noValidate>
              {/* Client Name Field */}
              <div>
                <Label
                  htmlFor="clientName"
                  className="block text-gray-700 font-semibold mb-2"
                >
                  Client Name
                </Label>
                <Input
                  id="clientName"
                  type="text"
                  placeholder="Client Name"
                  value={clientName}
                  onChange={(e) => setClientName(e.target.value)}
                  required
                  className={`w-full p-3 border rounded-lg focus:ring-2 ${
                    redirectUrlError
                      ? 'border-red-500'
                      : 'border-gray-300'
                  } focus:ring-yellow-400`}
                />
              </div>

              {/* Redirect URL Field */}
              <div>
                <Label
                  htmlFor="redirectUrl"
                  className="block text-gray-700 font-semibold mb-2"
                >
                  Redirect URL
                </Label>
                <Input
                  id="redirectUrl"
                  type="url"
                  placeholder="https://yourapp.com/callback"
                  value={redirectUrl}
                  onChange={(e) => setRedirectUrl(e.target.value)}
                  required
                  className={`w-full p-3 border rounded-lg focus:ring-2 ${
                    redirectUrlError
                      ? 'border-red-500'
                      : 'border-gray-300'
                  } focus:ring-yellow-400`}
                />
                {redirectUrlError && (
                  <p className="mt-1 text-red-500 text-sm">
                    {redirectUrlError}
                  </p>
                )}
              </div>

              {/* Register Button */}
              <div className="mt-6">
                <Button
                  type="submit"
                  className={`bg-black text-white w-full py-3 rounded-full text-lg font-semibold hover:bg-gray-800 transition-all duration-300 ${
                    (clientName === '' || redirectUrl === '' || redirectUrlError) &&
                    'opacity-50 cursor-not-allowed'
                  }`}
                  disabled={
                    clientName === '' ||
                    redirectUrl === '' ||
                    Boolean(redirectUrlError)
                  }
                >
                  REGISTER
                </Button>
              </div>
            </form>
          </div>

          {/* Right Column: Display Keys */}
          <div className="space-y-8">
            {/* Client ID Field */}
            <div>
              <Label
                htmlFor="clientId"
                className="block text-2xl font-bold text-black mb-2"
              >
                CLIENT ID
              </Label>
              <div className="flex items-center">
                <Input
                  id="clientId"
                  type="text"
                  placeholder="Your Client ID will appear here"
                  value={clientId}
                  readOnly
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                />
                <Button
                  className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300"
                  onClick={() => handleCopy(clientId)}
                  disabled={!clientId}
                >
                  <ClipboardIcon className="h-6 w-6 text-gray-600" />
                </Button>
              </div>
            </div>

            {/* Public Key Field */}
            <div>
              <Label
                htmlFor="publicKey"
                className="block text-2xl font-bold text-black mb-2"
              >
                PUBLIC KEY
              </Label>
              <div className="flex items-center">
                <Input
                  id="publicKey"
                  type="text"
                  placeholder="Your Public Key will appear here"
                  value={publicKey}
                  readOnly
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                />
                <Button
                  className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300"
                  onClick={() => handleCopy(publicKey)}
                  disabled={!publicKey}
                >
                  <ClipboardIcon className="h-6 w-6 text-gray-600" />
                </Button>
              </div>
            </div>

            {/* Private Key Field */}
            <div>
              <Label
                htmlFor="privateKey"
                className="block text-2xl font-bold text-black mb-2"
              >
                PRIVATE KEY
              </Label>
              <div className="flex items-center">
                <Input
                  id="privateKey"
                  type="text"
                  placeholder="Your Private Key will appear here"
                  value={privateKey}
                  readOnly
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                />
                <Button
                  className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300"
                  onClick={() => handleCopy(privateKey)}
                  disabled={!privateKey}
                >
                  <ClipboardIcon className="h-6 w-6 text-gray-600" />
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

