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
    active: boolean
}