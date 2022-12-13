import React from "react"

interface Props extends React.HtmlHTMLAttributes<HTMLHeadElement> {
	children: any
}

const H1 = (props: Props) => {
	return (
		<h1 className={"text-center text-blue-700 font-bold text-2xl"}>
			{props.children}
		</h1>
	)
}

export default H1
