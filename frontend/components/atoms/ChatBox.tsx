import { useRouter } from "next/router"
import React from "react"
import { ChatBoxMessage } from "../../types"

type Props = {
	contact: ChatBoxMessage
}

const ChatBox = ({ contact }: Props) => {
	const { push } = useRouter()

	return (
		<div
			onClick={() => push(`/chat/${contact.number}`)}
			className="w-full border-b-4 flex flex-col justify-start items-start px-5 py-5"
		>
			<div className="py-2 text-blue-700 font-semibold text-2xl">
				{contact.number}
			</div>
			<div className="py-1 text-lg">{contact.lastMessage}</div>
		</div>
	)
}

export default ChatBox
