import { useRouter } from "next/router"
import React from "react"

const messages = [
	{
		from: "8910719147",
		to: "Yash",
		message: "Hi",
	},
	{
		from: "8910719147",
		to: "Yash",
		message: "Hello",
	},
	{
		from: "Yash",
		to: "8910719147",
		message: "How are you?",
	},
	{
		from: "8910719147",
		to: "Yash",
		message: "I am fine.",
	},
]

const ChatScreen = () => {
	const {
		query: { number },
		back,
	} = useRouter()

	return (
		<div className="flex flex-col justify-between h-screen">
			<div className="flex flex-col justify-start h-5/6">
				<div className="flex justify-start items-center p-3 border-b border-gray-300">
					<span onClick={() => back()} className="text-blue-700 w-5">
						{"<"}
					</span>
					<span className="block ml-2 font-bold text-blue-700">{number}</span>
				</div>
				<div className="w-full p-6 min-h-full overflow-y-auto">
					<ul className="space-y-2">
						{messages.map((msg, i) => {
							const textColor =
								msg.from === "8910719147" ? "text-white" : "text-black"
							const bgColor =
								msg.from === "8910719147" ? "bg-blue-500" : "bg-gray-200"
							return (
								<React.Fragment key={i}>
									<li
										className={`flex justify-${
											msg.from === "8910719147" ? "end" : "start"
										}`}
										key={i}
									>
										<div
											className={`relative max-w-xl px-4 py-2 ${textColor} ${bgColor} rounded shadow text-`}
										>
											<span className="block">{msg.message}</span>
										</div>
									</li>
								</React.Fragment>
							)
						})}
					</ul>
				</div>
			</div>
			<div className="border-l-2 h-[5rem]">
				<div className="flex items-center justify-center w-full p-3 border-t border-gray-300">
					<input
						type="text"
						placeholder="Message"
						className="block w-full py-2 pl-4 mx-3 bg-gray-100 rounded-full outline-none focus:text-gray-700"
						name="message"
						required
					/>
					<button type="submit">
						<svg
							className="w-5 h-5 text-blue-500 origin-center transform rotate-90"
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							fill="currentColor"
						>
							<path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
						</svg>
					</button>
				</div>
			</div>
		</div>
	)
}

export default ChatScreen
