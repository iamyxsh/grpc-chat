import React from "react"

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {}

const TextInput = (props: Props) => {
	return <input className="m-10 p-2 rounded-md" {...props} />
}

export default TextInput
