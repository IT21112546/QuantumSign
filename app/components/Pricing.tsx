import { Button } from "@/components/ui/button"

export default function Pricing() {
  return (
    <section className="py-12 bg-gray-100">
      <div className="text-center mb-8">
        <h2 className="text-3xl font-bold text-black">CHOOSE YOUR PACKAGE</h2>
        <div className="mt-4 border-t-2 w-16 mx-auto border-black"></div>
      </div>

      <div className="container mx-auto grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Free Trial Plan */}
        <div className="bg-white p-6 rounded-xl shadow-lg text-center">
          <h3 className="text-xl font-bold">FREE TRIAL</h3>
          <ul className="text-left mt-4 mb-6">
            <li>• 1000 API calls per day</li>
            <li>• Rate limited to 100 authentications per second</li>
            <li>• 10,000 API calls per month</li>
          </ul>
          <div className="text-3xl font-bold">$0.00</div>
          <div className="text-gray-500 text-sm">/API Request</div>
          <Button className="mt-6 bg-yellow-400 text-black font-semibold py-2 px-6 rounded-full hover:bg-yellow-500">
            BUY NOW
          </Button>
        </div>

        {/* Basic Plan */}
        <div className="bg-white p-6 rounded-xl shadow-lg text-center">
          <h3 className="text-xl font-bold">BASIC</h3>
          <ul className="text-left mt-4 mb-6">
            <li>• 10,000 API calls per day</li>
            <li>• Rate limited to 1000 authentications per second</li>
            <li>• 1,000,000 API calls per month</li>
          </ul>
          <div className="text-3xl font-bold">$0.01</div>
          <div className="text-gray-500 text-sm">/API Request</div>
          <Button className="mt-6 bg-yellow-400 text-black font-semibold py-2 px-6 rounded-full hover:bg-yellow-500">
            BUY NOW
          </Button>
        </div>

        {/* Standard Plan */}
        <div className="bg-white p-6 rounded-xl shadow-lg text-center">
          <h3 className="text-xl font-bold">STANDARD</h3>
          <ul className="text-left mt-4 mb-6">
            <li>• Unlimited API calls per day</li>
            <li>• Rate limited to 10,000 authentications per second</li>
            <li>• Unlimited API calls per month</li>
          </ul>
          <div className="text-3xl font-bold">$0.05</div>
          <div className="text-gray-500 text-sm">/API Request</div>
          <Button className="mt-6 bg-yellow-400 text-black font-semibold py-2 px-6 rounded-full hover:bg-yellow-500">
            BUY NOW
          </Button>
        </div>
      </div>
    </section>
  );
}
