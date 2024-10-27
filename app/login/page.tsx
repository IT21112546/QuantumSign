'use client'

import { Suspense } from 'react'
import Login from './Login'
import { useSearchParams } from 'next/navigation'

const page = () => {

  const searchParams = useSearchParams()
  const token = searchParams.get('token')

	return (
		<Suspense>
			<Login token={token} />
		</Suspense>
	)
}

export default page
