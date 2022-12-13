import { useRouter } from "next/router"
import React from "react"
import OtpInput from "react-otp-input"
import { Button, H1, TextInput } from "../components/atoms"
import { Container } from "../components/template"

const OTP = () => {
	const { push } = useRouter()

	return (
		<Container>
			<H1>Enter the OTP</H1>
			<OtpInput
				className="my-5"
				containerStyle="p-1"
				inputStyle="p-1"
				value={""}
				onChange={(otp: string) => console.log(otp)}
				numInputs={4}
				separator={<span>-</span>}
				isInputNum={true}
			/>
			<Button onClick={() => push("/chats")}>Login</Button>
		</Container>
	)
}

export default OTP
