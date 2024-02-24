package models

type Task struct {
	ExecutionName string `json:"execution_name"`
	Name          string `json:"name"`
	Type          string `json:"type"`
}
