import LoginPage from '../components/LoginPage'
import { Suspense } from 'react'

const page = () => {
  return (
		<Suspense>
			<LoginPage />
		</Suspense>
  )
}

export default page
