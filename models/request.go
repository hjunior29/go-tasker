package models

// TasksRequest represents the structure of a TasksRequest for processing.
type TasksRequest struct {
	URL         string                 `json:"url"`
	Method      string                 `json:"method"`
	ScheduledAt int                    `json:"scheduled_at"`
	Payload     map[string]interface{} `json:"payload"`
}
