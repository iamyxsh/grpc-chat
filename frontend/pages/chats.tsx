import React, { useEffect } from "react"
import { H1 } from "../components/atoms"
import { ChatboxList } from "../components/molecules"
import { Container } from "../components/template"

const contacts = [
	{
		number: "Yash",
		lastMessage: "Hello",
	},
	{
		number: "Yash",
		lastMessage: "Hello",
	},
	{
		number: "Yash",
		lastMessage: "Hello",
	},
	{
		number: "Yash",
		lastMessage: "Hello",
	},
]

const Chats = () => {
	return (
		<Container align="start">
			<div className="py-5">
				<H1>Chats</H1>
			</div>
			<div className="w-full p-0.5 bg-blue-500"></div>
			<ChatboxList contacts={contacts} />
		</Container>
	)
}

export default Chats
