import React from "react"

type Props = {
	children: any
	class?: string
	align?: "center" | "start"
}

const Container = (props: Props) => {
	return (
		<div
			className={
				"sm:container bg-blue-100 h-screen flex flex-col " +
					props.class +
					props.align ===
				"start"
					? props.align
					: "justify-center items-center"
			}
		>
			{props.children}
		</div>
	)
}

export default Container
