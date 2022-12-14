import React from "react"

interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {}

const Button = (props: Props) => {
	return (
		<button
			{...props}
			className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
		>
			{props.children}
		</button>
	)
}

export default Button
