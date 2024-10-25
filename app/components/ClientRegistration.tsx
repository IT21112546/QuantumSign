import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Button } from "@/components/ui/button"; // ShadCN Button component
import { Label } from "@/components/ui/label"; // ShadCN Label component
import { ClipboardIcon } from "@heroicons/react/24/outline"; // For clipboard icon

export default function ClientRegistrationForm() {
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

            <form className="space-y-6">
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
                  className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
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
                  type="text"
                  placeholder="Redirect URL"
                  className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
                />
              </div>

              {/* Register Button */}
              <div className="mt-6">
                <Button className="bg-black text-white w-full py-3 rounded-full text-lg font-semibold hover:bg-gray-800 transition-all duration-300">
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
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                  readOnly
                />
                <Button className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300">
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
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                  readOnly
                />
                <Button className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300">
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
                  className="w-full p-3 border border-gray-300 rounded-lg bg-gray-50 focus:outline-none"
                  readOnly
                />
                <Button className="ml-3 bg-gray-200 p-2 rounded-lg hover:bg-gray-300">
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

