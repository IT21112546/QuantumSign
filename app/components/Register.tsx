import { Button } from "@/components/ui/button"; // ShadCN Button component
import { Input } from "@/components/ui/input"; // ShadCN Input component
import { Label } from "@/components/ui/label"; // ShadCN Label component

export default function SignUpForm() {
  return (
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-r from-gray-100 to-gray-200">
      <div className="bg-white p-10 rounded-3xl shadow-2xl w-full max-w-lg">
        <h2 className="text-3xl font-bold text-center mb-6 text-gray-800">
          SIGN UP
        </h2>

        {/* Form Inputs */}
        <form>
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
              <Button
                variant="outline"
                className="bg-yellow-400 text-black font-semibold py-2 px-6 rounded-full hover:bg-yellow-500"
              >
                Upload
              </Button>
            </div>
          </div>

          {/* Register Button */}
          <div className="mt-6 text-center">
            <Button
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

