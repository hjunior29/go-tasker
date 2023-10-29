package db

import (
	"github.com/hjunior29/go-tasker/models"
	"gorm.io/gorm"
)

// GetTotalTasks retrieves the total number of tasks from the database.
func GetTotalTasks() (int64, error) {
	var count int64
	processed, err := GetTotalProcessedTasks()
	failed, err := GetTotalFailedTasks()
	pending, err := GetPendingTasks()
	count = processed + failed + pending
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetTotalProcessedTasks retrieves the total number of processed tasks from the database.
func GetTotalProcessedTasks() (int64, error) {
	var count int64
	if err := DB.Model(&models.TasksLogs{}).Where("status = ?", "Sent Success").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetTotalPostRequests retrieves the total number of POST requests from the database.
func GetTotalPostRequests() (int64, error) {
	var count int64
	if err := DB.Unscoped().Model(&models.Tasks{}).Where("method = ?", "POST").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetTotalPutRequests retrieves the total number of PUT requests from the database.
func GetTotalPutRequests() (int64, error) {
	var count int64
	if err := DB.Unscoped().Model(&models.Tasks{}).Where("method = ?", "PUT").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetTotalFailedTasks retrieves the total number of failed tasks from the database.
func GetTotalFailedTasks() (int64, error) {
	var count int64
	if err := DB.Model(&models.TasksLogs{}).Where("status = ?", "Sent Failed").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetPendingTasks retrieves the total number of tasks that are in the queue (not processed yet).
func GetPendingTasks() (int64, error) {
	var count int64
	if err := DB.Model(&models.Tasks{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetAverageProcessingTime will return the average processing duration in seconds
func GetAverageProcessingTime() (float64, error) {
	var result struct {
		AverageTime float64
	}

	err := DB.Raw(`
		SELECT AVG(EXTRACT(EPOCH FROM (created_at - processed_at)) * 1000) AS average_time
		FROM tasks_logs
		WHERE processed_at IS NOT NULL
		AND status = ?`, "Sent Success").Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.AverageTime, nil
}

// GetConfig retrieves the latest task configuration or creates a default one.
func GetConfig() (*models.TasksConfig, error) {
	var config models.TasksConfig

	if err := DB.Last(&config).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			defaultConfig := models.TasksConfig{Workers: 1}
			if createErr := DB.Create(&defaultConfig).Error; createErr != nil {
				return nil, createErr
			}
			return &defaultConfig, nil
		}
		return nil, err
	}

	return &config, nil
}
