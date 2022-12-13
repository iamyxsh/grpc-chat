import { useEffect, useState } from "react"
import LoginPage from "./login"

export default function Home() {
	const [contacts, setContacts] = useState<any[]>([])
	const [isAndroid, setIsAndroid] = useState(true)

	useEffect(() => {
		const getContacts = async () => {
			try {
				// @ts-ignore
				const contacts = await navigator.contacts.select(["name", "tel"], {
					multiple: true,
				})

				setContacts(contacts)
			} catch (err) {
				setIsAndroid(false)
				console.log("err", err)
			}
		}
		getContacts()
	}, [])

	return (
		<div>
			<LoginPage />
		</div>
	)
}
