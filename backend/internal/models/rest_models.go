package models

type TriggerRequest struct {
	Type string `json:"type"`
}

type CreateProjectRequest struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
