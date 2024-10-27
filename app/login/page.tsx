import LoginPage from '../components/LoginPage'
import NavBar from '../components/NavBar'
import { Suspense } from 'react'

const page = () => {
  return (
		<Suspense>
			<LoginPage />
		</Suspense>
  )
}

export default page
