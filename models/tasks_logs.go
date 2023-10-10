package models

import (
	"time"

	"gorm.io/gorm"
)

// TasksLogs represents the structure of a TasksRequest log in the database.
type TasksLogs struct {
	gorm.Model
	TaskID      string
	Payload     string
	Status      string
	LogMessage  string
	ProcessedAt time.Time
	LogType     string
	Method      string
	URL         string
}
