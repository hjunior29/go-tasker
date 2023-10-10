package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunior29/go-tasker/models"
	"github.com/hjunior29/go-tasker/store"
)

// InfoQueueHandler retrieves and returns various information about the task queue.
func InfoQueueHandler(c *gin.Context) {
	// Retrieve various pieces of information about the task queue from the database.
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

	// Return the retrieved information as JSON.
	c.JSON(200, gin.H{
		"totalTasks":          totalTasks,
		"totalProcessedTasks": totalProcessedTasks,
		"totalPostRequests":   totalPostRequests,
		"totalPutRequests":    totalPutRequests,
		"totalFailedTasks":    totalFailedTasks,
		"pendingTasks":        pendingTasks,
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
