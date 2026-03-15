const BASE_URL = import.meta.env.VITE_API_URL

export const api = {
    get: (path: string): Promise<any> => 
        fetch(`${BASE_URL}${path}`).then(res => {
            if(!res.ok) {
                console.error("API error:", res.status)
                return null
            }
            console.log(res)
            return res.json()
        }),
    post: (path: string, body: unknown) => 
        fetch(`${BASE_URL}${path}`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        }).then(res => {
            if(!res.ok) {
                console.error("API error:", res.status)
                return null
            }
            return res.json()
        })
}