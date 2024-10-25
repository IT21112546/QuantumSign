import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Button } from "@/components/ui/button"; // ShadCN Button component
import { Label } from "@/components/ui/label"; // ShadCN Label component

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white shadow-lg rounded-2xl p-8 max-w-md w-full">
        {/* Title */}
        <h2 className="text-3xl font-bold text-gray-800 text-center mb-8">
          Sign In
        </h2>

        {/* Public Key Input */}
        <form>
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
              />
            </div>

            {/* OR Divider */}
            <div className="relative my-6">
              <span className="absolute inset-x-0 top-1/2 transform -translate-y-1/2 text-center text-gray-400 bg-white px-2">
                OR
              </span>
              <div className="border-t border-gray-300"></div>
            </div>

            {/* File Upload Section */}
            <div className="text-center">
              <Label className="block text-lg font-semibold text-gray-700 mb-4">
                Upload Your File
              </Label>
              <div className="bg-gray-50 rounded-lg p-6 shadow-md">
                <Button className="bg-green-500 text-white px-6 py-2 rounded-full hover:bg-green-600 transition-all duration-300">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    className="w-6 h-6 inline-block mr-2"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M3 16.5V18a2.25 2.25 0 002.25 2.25h13.5A2.25 2.25 0 0021 18v-1.5M7.5 12l4.5-4.5m0 0L16.5 12m-4.5-4.5V18"
                    />
                  </svg>
                  Upload
                </Button>
              </div>
            </div>

            {/* Login Button */}
            <div>
              <Button className="bg-black text-white w-full py-3 rounded-full text-lg font-semibold hover:bg-gray-900 transition-all duration-300">
                Login
              </Button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}

