package models

import (
	"gorm.io/gorm"
)

type TasksConfig struct {
	gorm.Model
	Workers int `json:"workers" gorm:"default:1"`
}
