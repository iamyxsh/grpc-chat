import { useRouter } from "next/router"
import React from "react"
import { H1, TextInput } from "../components/atoms"
import Button from "../components/atoms/Button"
import { Container } from "../components/template"

const LoginPage = () => {
	const { push } = useRouter()
	return (
		<Container>
			<H1>Login via number</H1>
			<p className="text-center mt-5 text-blue-00">
				Enter your number to get an OTP
			</p>
			<TextInput placeholder="Number" />
			<Button onClick={() => push("/otp")}>Send</Button>
		</Container>
	)
}

export default LoginPage
