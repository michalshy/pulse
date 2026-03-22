export type ProjectFormData = {
    key: string
    name: string
    description: string
    retention_days: number
}

export type Project = {
    id: number
    key: string
    name: string
    description: string
    created_at: string
    retention_days: number
    last_heartbeat: string | null
}