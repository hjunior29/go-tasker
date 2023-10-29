package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hjunior29/go-tasker/models"
	"github.com/hjunior29/go-tasker/store"
)

// HomeQueueHandler get infos about API.
func HomeQueueHandler(c *gin.Context) {
	infoMessage := map[string]interface{}{
		"*Welcome": map[string]string{
			"message":   "Welcome to the Go-Tasker API!",
			"version":   "1.0.0",
			"timestamp": time.Now().Format(time.UnixDate),
		},
		"Docs": map[string]string{
			"DockerHub": "https://hub.docker.com/r/hjunior29/go-tasker",
			"GitHub":    "https://github.com/hjunior29/go-tasker",
		},
	}
	c.JSON(200, infoMessage)
}

// InfoQueueHandler retrieves and returns various information about the task queue.
func InfoQueueHandler(c *gin.Context) {
	averageProcessingTime, err := store.GetAverageProcessingTime()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get average processing time tasks"})
		return
	}

	totalTasks, err := store.GetTotalTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve total tasks"})
		return
	}

	totalProcessedTasks, err := store.GetTotalProcessedTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve total processed tasks"})
		return
	}

	totalPostRequests, err := store.GetTotalPostRequests()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve total POST requests"})
		return
	}

	totalPutRequests, err := store.GetTotalPutRequests()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve total PUT requests"})
		return
	}

	totalFailedTasks, err := store.GetTotalFailedTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve total failed tasks"})
		return
	}

	pendingTasks, err := store.GetPendingTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve pending tasks"})
		return
	}

	c.JSON(200, gin.H{
		"tasksMetrics": gin.H{
			"AverageProcessingTime": fmt.Sprintf("%.10f ms", averageProcessingTime),
			"total":                 totalTasks,
			"processed":             totalProcessedTasks,
			"failed":                totalFailedTasks,
			"pending":               pendingTasks,
		},
		"tasksTypes": gin.H{
			"post": totalPostRequests,
			"put":  totalPutRequests,
		},
	})

}

// EnqueueTaskHandler handles the enqueuing of a new task.
func EnqueueTaskHandler(c *gin.Context) {
	var task models.TasksRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := store.EnqueueTask(task); err != nil {
		c.JSON(500, gin.H{"error": "Failed to enqueue task"})
		return
	}

	c.JSON(200, gin.H{"status": "Task enqueued successfully!"})
}

func GetQueueConfigHandler(c *gin.Context) {
	config, err := store.GetConfig()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get configuration"})
		return
	}
	c.JSON(200, gin.H{
		"TasksConfigs": gin.H{
			"workers": config.Workers,
		},
	})
}

func UpdateQueueConfigHandler(c *gin.Context) {
	var config models.TasksConfig

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if config.Workers <= 0 {
		c.JSON(400, gin.H{"error": "Invalid number of workers"})
		return
	}

	if err := store.UpdateConfig(config); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update configuration"})
		return
	}

	c.JSON(200, gin.H{"status": "Configuration updated successfully!"})

}
