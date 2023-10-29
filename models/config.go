package models

import (
	"gorm.io/gorm"
)

// TasksConfig represents the task configuration structure in the database
type TasksConfig struct {
	gorm.Model
	Workers int `json:"workers" gorm:"default:1"`
}
