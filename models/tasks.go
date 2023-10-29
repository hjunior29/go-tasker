package models

import (
	"time"

	"gorm.io/gorm"
)

// Tasks represents the structure of a TasksRequest in the database.
type Tasks struct {
	gorm.Model
	TaskID      string
	Payload     string
	Method      string
	URL         string
	ScheduledAt *time.Time `gorm:"type:timestamp"`
}
