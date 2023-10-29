package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hjunior29/go-tasker/models"
	"gorm.io/gorm"
)

// EnqueueTask adds a new task to the PostgreSQL queue using GORM.
func EnqueueTask(task models.TasksRequest) error {
	var status, logMessage string

	taskID := uuid.New().String()
	logType := "Receive"

	payloadBytes, err := json.Marshal(task.Payload)
	if err != nil {
		return err
	}

	scheduledTime := DetermineScheduledTime(task.ScheduledAt)

	TasksRequest := models.Tasks{
		TaskID:      taskID,
		Payload:     string(payloadBytes),
		Method:      task.Method,
		URL:         task.URL,
		ScheduledAt: scheduledTime,
	}

	if err := DB.Create(&TasksRequest).Error; err != nil {
		status = "Received failed"
		logMessage = fmt.Sprintf("Failed to enqueue TasksRequest: %v", err)
		LogTask(DB, taskID, TasksRequest.Payload, status, logMessage, logType, task.Method, task.URL, time.Now())
		return err
	}

	status = "Received success"
	logMessage = "TasksRequest enqueued successfully"
	LogTask(DB, taskID, TasksRequest.Payload, status, logMessage, logType, task.Method, task.URL, time.Now())

	fmt.Println(logMessage)
	return nil
}

// DequeueTask retrieves and removes a task from the PostgreSQL queue using GORM.
func DequeueTask() (*models.Tasks, error) {
	var TasksRequest models.Tasks

	query := DB.Where("(scheduled_at IS NULL) OR (scheduled_at <= ?)", time.Now()).Order("created_at asc")
	if err := query.First(&TasksRequest).Error; err != nil {
		return nil, err
	}

	if err := DB.Delete(&TasksRequest).Error; err != nil {
		return nil, err
	}

	return &TasksRequest, nil
}

// DetermineScheduledTime calculates the future time based on a given delay in minutes.
func DetermineScheduledTime(minutesDelay int) *time.Time {
	if minutesDelay == 0 {
		return nil
	}

	delayDuration := time.Duration(minutesDelay) * time.Minute
	scheduledTime := time.Now().Add(delayDuration)

	return &scheduledTime
}

// LogTask records the details of a task processing attempt in the database.
func LogTask(db *gorm.DB, taskID, payload, status, logMessage, logType, method, url string, processedAt time.Time) error {
	logTask := models.TasksLogs{
		TaskID:      taskID,
		Payload:     payload,
		Status:      status,
		LogMessage:  logMessage,
		ProcessedAt: processedAt,
		LogType:     logType,
		Method:      method,
		URL:         url,
	}

	if err := db.Create(&logTask).Error; err != nil {
		return fmt.Errorf("failed to log TasksRequest: %w", err)
	}

	return nil
}

// UpdateConfig updates the task configuration in the database.
func UpdateConfig(newConfig models.TasksConfig) error {
	if err := DB.Save(&newConfig).Error; err != nil {
		return err
	}

	return nil
}