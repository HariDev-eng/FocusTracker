package router

import (
	"focustracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func New(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "everything working properly"})
	})

	r.GET("/api/tasks", h.ListTasks)
	r.POST("/api/tasks", h.CreateTask)
	r.POST("/api/tasks/:id/complete", h.ToggleCompletion)
	r.PUT("/api/tasks/:id", h.UpdateTask)
	r.DELETE("/api/tasks/:id", h.ArchiveTask)

	return r
}
