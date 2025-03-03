import { useEffect, useState } from "react"
import { UserAccount } from "../types"


export const useUser = () => {
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [user, setUser] = useState<UserAccount | null>(null)
    
    const getUser = async () => {
        setIsLoading(true)
        const resp = await fetch("/api/get-me", {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        })

        console.log(resp.status);
        

        const json = await resp.json()

        setIsLoading(true)
    }

    useEffect(() => {
        getUser()
    }, [])

    return {user, isLoading}
}