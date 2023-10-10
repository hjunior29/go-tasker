package store

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hjunior29/go-tasker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// InitDatabase establishes a connection to the PostgreSQL database and initializes it.
func InitDatabase() error {

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SSLMODE"),
	)

	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.Tasks{}, &models.TasksLogs{})

	return nil
}

// EnqueueTask adds a new task to the PostgreSQL queue using GORM.
func EnqueueTask(task models.TasksRequest) error {
	var status, logMessage string

	taskID := uuid.New().String()
	logType := "Receive"

	payloadBytes, err := json.Marshal(task.Payload)
	if err != nil {
		return err
	}

	TasksRequest := models.Tasks{
		TaskID:  taskID,
		Payload: string(payloadBytes),
		Method:  task.Method,
		URL:     task.URL,
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
	if err := DB.Order("created_at asc").First(&TasksRequest).Error; err != nil {
		return nil, err
	}

	if err := DB.Delete(&TasksRequest).Error; err != nil {
		return nil, err
	}

	return &TasksRequest, nil
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

// GetTotalTasks retrieves the total number of tasks from the database.
func GetTotalTasks() (int64, error) {
	var count int64
	if err := DB.Model(&models.Tasks{}).Count(&count).Error; err != nil {
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
	if err := DB.Model(&models.TasksLogs{}).Where("method = ?", "POST").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetTotalPutRequests retrieves the total number of PUT requests from the database.
func GetTotalPutRequests() (int64, error) {
	var count int64
	if err := DB.Model(&models.TasksLogs{}).Where("method = ?", "PUT").Count(&count).Error; err != nil {
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

