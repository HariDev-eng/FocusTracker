package handlers

import (
	"focustracker/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

const DefaultUserID = 1

func (h *Handler) ListTasks(c *gin.Context) {
	var tasks []models.Task
	if err := h.DB.Where("user_id = ? AND is_archived = ?", DefaultUserID, false).
		Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Frequency   string `json:"frequency"`
}

func (h *Handler) CreateTask(c *gin.Context) {
	var input CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Frequency == "" {
		input.Frequency = "daily"
	}

	task := models.Task{
		UserID:      DefaultUserID,
		Title:       input.Title,
		Description: input.Description,
		Frequency:   input.Frequency,
	}

	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) ToggleCompletion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	var body struct {
		Date string `json:"date"`
	}

	_ = c.ShouldBindJSON(&body)

	dateStr := body.Date
	if dateStr == "" {
		dateStr = time.Now().Format("2006-01-02")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date must be YYYY-MM-DD"})
		return
	}

	var existing models.Completion
	err = h.DB.Where("task_id = ? AND date = ?", id, date).First(&existing).Error
	if err == nil {
		if err = h.DB.Delete(&existing).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"completed": false,
		})
		return
	}

	completion := models.Completion{
		TaskID: uint(id),
		Date:   date,
	}

	if err := h.DB.Create(&completion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"completed": true,
	})
}
