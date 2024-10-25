'use client'

import { useState } from "react";
import { Button } from "@/components/ui/button"; // Assuming you're using ShadCN's Button
import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Label } from "@/components/ui/label"; // ShadCN Label component
import { ClipboardIcon, ArrowDownTrayIcon } from "@heroicons/react/24/outline"; // Icons for copy and download

export default function GenerateKeypair() {
  const [publicKey, setPublicKey] = useState("");
  const [privateKey, setPrivateKey] = useState("");
  const [error, setError] = useState("");

  const API_HOST = process.env.NEXT_PUBLIC_API_HOST || "localhost:8000";

  const handleGenerateKeypair = async () => {
    setError(""); // Reset the error before the request
    try {
      const response = await fetch(`http://${API_HOST}/generate/kyber`);
      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }

      const data = await response.json();
      setPublicKey(data.publicKey);
      setPrivateKey(data.privateKey);
    } catch (err) {
      console.error(err);
      setError("Failed to generate keypair. Please try again.");
    }
  };

  const handleCopyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
  };

  const handleDownload = (text: string, filename: string) => {
    const element = document.createElement("a");
    const file = new Blob([text], { type: "text/plain" });
    element.href = URL.createObjectURL(file);
    element.download = filename;
    document.body.appendChild(element); // Required for this to work in FireFox
    element.click();
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-2xl shadow-lg w-full max-w-lg">
        <h2 className="text-3xl font-bold mb-6 text-gray-900 text-center">Generate Kyber Keypair</h2>

        {/* Error Display */}
        {error && (
          <div className="bg-red-100 text-red-600 p-4 rounded-md mb-6">
            {error}
          </div>
        )}

        <div className="text-center mb-8">
          <Button className="bg-black text-white px-6 py-3 rounded-full hover:bg-gray-900 transition-all" onClick={handleGenerateKeypair}>
            Generate Keypair
          </Button>
        </div>

        {/* Public Key Display */}
        {publicKey && (
          <div className="mb-6">
            <Label htmlFor="publicKey" className="block text-lg font-semibold text-gray-700 mb-2">
              Public Key
            </Label>
            <div className="flex items-center">
              <Input
                id="publicKey"
                type="text"
                value={publicKey}
                readOnly
                className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50"
              />
              <Button
                className="ml-2 p-2 bg-gray-200 hover:bg-gray-300 rounded-lg"
                onClick={() => handleCopyToClipboard(publicKey)}
              >
                <ClipboardIcon className="h-5 w-5 text-gray-600" />
              </Button>
              <Button
                className="ml-2 p-2 bg-gray-200 hover:bg-gray-300 rounded-lg"
                onClick={() => handleDownload(publicKey, "id_kyber.pub")}
              >
                <ArrowDownTrayIcon className="h-5 w-5 text-gray-600" />
              </Button>
            </div>
          </div>
        )}

        {/* Private Key Display */}
        {privateKey && (
          <div className="mb-6">
            <Label htmlFor="privateKey" className="block text-lg font-semibold text-gray-700 mb-2">
              Private Key
            </Label>
            <div className="flex items-center">
              <Input
                id="privateKey"
                type="text"
                value={privateKey}
                readOnly
                className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50"
              />
              <Button
                className="ml-2 p-2 bg-gray-200 hover:bg-gray-300 rounded-lg"
                onClick={() => handleCopyToClipboard(privateKey)}
              >
                <ClipboardIcon className="h-5 w-5 text-gray-600" />
              </Button>
              <Button
                className="ml-2 p-2 bg-gray-200 hover:bg-gray-300 rounded-lg"
                onClick={() => handleDownload(privateKey, "id_kyber")}
              >
                <ArrowDownTrayIcon className="h-5 w-5 text-gray-600" />
              </Button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
