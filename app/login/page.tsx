import { Suspense } from 'react'
import Login from './Login'

const page = () => {
	return (
		<Suspense>
			<Login />
		</Suspense>
	)
}

export default page
