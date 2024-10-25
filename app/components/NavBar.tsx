import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function NavBar() {
  return (
    <div className="bg-yellow-400 w-full py-4">
      <div className="container mx-auto flex items-center justify-between">
        {/* Logo */}
        <div className="text-black text-xl font-bold">QSIGN</div>

        {/* Navigation Menu */}
        <nav className="flex items-center justify-center bg-white px-6 py-2 rounded-full shadow-md">
          <ul className="flex items-center space-x-4 text-black">
            <li>
              <Link href="/">Home</Link>
            </li>
            <span className="text-gray-400">|</span>
            <li>
              <Link href="/pricing">Pricing</Link>
            </li>
            <span className="text-gray-400">|</span>
            <li>
              <Link href="/register-client">Register Client Application</Link>
            </li>
            <span className="text-gray-400">|</span>
            <li>
              <a href="https://research.qsign.io">Our Research</a>
            </li>
          </ul>

          {/* Register and Login Buttons */}
          <div className="flex items-center space-x-2 ml-6">
						<Link href="/register">
							<Button className="bg-white text-black rounded-full px-4 py-1 hover:bg-black hover:text-white">
								Register
							</Button>
						</Link>
						<Link href="/login">
							<Button className="bg-black text-white rounded-full px-4 py-1">
								Login
							</Button>
						</Link>
          </div>
        </nav>
      </div>
    </div>
  );
}

