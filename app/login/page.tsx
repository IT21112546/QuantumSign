import LoginPage from '../components/LoginPage'
import { Suspense } from 'react'
import Footer from '../components/Footer'

const page = () => {
	return (
		<Suspense>
			<LoginPage />
			<Footer />
		</Suspense>
	)
}

export default page
