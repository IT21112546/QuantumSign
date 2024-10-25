import { Button } from "@/components/ui/button"; // ShadCN Button component
import { CheckIcon } from "@heroicons/react/24/solid"; // Heroicons Check Icon
import Link from "next/link";

export default function HomePage() {
	return (
		<div>
			{/* Hero Section */}
			<section className="bg-gradient-to-r from-gray-900 to-gray-700 text-white py-20">
				<div className="container mx-auto text-center px-6">
					<h1 className="text-5xl font-extrabold mb-6">
						Quantum-resistant authentication
					</h1>
					<p className="text-xl text-gray-300 mb-12">
						Qsign provides robust, scalable, and easy-to-implement future-proof authentication solutions.
					</p>
					<div className="flex justify-center space-x-6">
						<Link href='/login'>
							<Button className="bg-yellow-500 text-black px-6 py-4 rounded-full hover:bg-yellow-600 transition-all text-lg font-semibold shadow-lg">
								Get Started
							</Button>
						</Link>
					</div>
				</div>
			</section>

			{/* Features Section */}
			<section className="py-20 bg-gray-100">
				<div className="container mx-auto px-6 text-center">
					<h2 className="text-4xl font-extrabold text-gray-900 mb-12">
						Why Choose Us?
					</h2>
					<div className="grid grid-cols-1 md:grid-cols-3 gap-10">
						{/* Feature 1 */}
						<div className="bg-white p-8 rounded-xl shadow-xl hover:shadow-2xl transition-shadow">
							<CheckIcon className="h-12 w-12 text-green-500 mx-auto mb-4" />
							<h3 className="text-2xl font-bold mb-2">Post-Quantum Authentication</h3>
							<p className="text-gray-600">
								Our solution is quantum-resistant, ensuring your data is secure against future threats.
							</p>
						</div>

						{/* Feature 2 */}
						<div className="bg-white p-8 rounded-xl shadow-xl hover:shadow-2xl transition-shadow">
							<CheckIcon className="h-12 w-12 text-green-500 mx-auto mb-4" />
							<h3 className="text-2xl font-bold mb-2">Easy Integration</h3>
							<p className="text-gray-600">
								We provide seamless integration with your existing systems, allowing for faster adoption.
							</p>
						</div>

						{/* Feature 3 */}
						<div className="bg-white p-8 rounded-xl shadow-xl hover:shadow-2xl transition-shadow">
							<CheckIcon className="h-12 w-12 text-green-500 mx-auto mb-4" />
							<h3 className="text-2xl font-bold mb-2">Open Source</h3>
							<p className="text-gray-600">
								Our code is open source, allowing you to audit and customize it to your needs.
							</p>
						</div>
					</div>
				</div>
			</section>

			{/* Call to Action Section */}
			<section className="bg-yellow-500 py-20">
				<div className="container mx-auto text-center px-6">
					<h2 className="text-4xl font-extrabold text-black mb-6">
						Secure your application&apos;s user authentication today
					</h2>
					<p className="text-lg text-gray-800 mb-10">
						Join thousands of developers who trust Qsign for their authentication needs.
					</p>
					<Link href='/login'>
						<Button className="bg-black text-white px-6 py-4 rounded-full hover:bg-gray-900 transition-all text-lg font-semibold shadow-lg">
							Get Started Now
						</Button>
					</Link>
				</div>
			</section>
		</div>
	);
}

