'use client'

import { useEffect } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import Login from './Login'
import NavBar from './NavBar'

const LoginPage = () => {
	const router = useRouter()
	const searchParams = useSearchParams()
	const clientId = searchParams.get('client-id')

	useEffect(() => {
		const redirectUrl = process.env.NEXT_PUBLIC_TEST_CLIENT_URL || "http://localhost:3001"
		if (!clientId) {
			router.push(redirectUrl)
		}
	}, [clientId, router])

	if (!clientId) {
		return null // or a loading spinner
	}

	return (
		<div>
			<NavBar />
			<Login clientId={clientId} />
		</div>
	)
}

export default LoginPage
