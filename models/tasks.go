package models

import (
	"gorm.io/gorm"
)

// Tasks represents the structure of a TasksRequest in the database.
type Tasks struct {
	gorm.Model
	TaskID  string
	Payload string
	Method  string
	URL     string
}
