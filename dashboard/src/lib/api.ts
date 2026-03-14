const BASE_URL = import.meta.env.VITE_API_URL

export const api = {
    get: (path: string): Promise<any> => fetch(`${BASE_URL}${path}`).then(
        res => {
            console.log(res)
            res.json()}),
    post: (path: string, body: unknown): Promise<any> => fetch(`${BASE_URL}${path}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(body)
    }).then(res => res.json())
}