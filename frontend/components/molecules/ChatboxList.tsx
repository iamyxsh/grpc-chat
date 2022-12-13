import React from "react"
import { ChatBoxMessage } from "../../types"
import ChatBox from "../atoms/ChatBox"

type Props = {
	contacts: ChatBoxMessage[]
}

const ChatboxList = ({ contacts }: Props) => {
	return (
		<div>
			{contacts.map((contact, i) => (
				<ChatBox contact={contact} key={i} />
			))}
		</div>
	)
}

export default ChatboxList
