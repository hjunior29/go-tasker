package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunior29/go-tasker/api/handlers"
)

// SetupRoutes defines the API routes.
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.HomeQueueHandler)
	r.GET("/info", api.InfoQueueHandler)
	r.POST("/enqueue", api.EnqueueTaskHandler)
	r.GET("/config", api.GetQueueConfigHandler)
	r.PUT("/config", api.UpdateQueueConfigHandler)

	return r
}
