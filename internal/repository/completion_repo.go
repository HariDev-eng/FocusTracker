package repository

import (
	"errors"
	"time"

	"focustracker/internal/models"

	"gorm.io/gorm"
)

type CompletionRepository struct {
	DB *gorm.DB
}

func NewCompletionRepository(db *gorm.DB) *CompletionRepository {
	return &CompletionRepository{DB: db}
}

func (r *CompletionRepository) FindByTaskID(taskID uint) ([]models.Completion, error) {
	var completions []models.Completion
	err := r.DB.Where("task_id = ?", taskID).Find(&completions).Error
	return completions, err
}

// FindByTaskAndDate returns (nil, nil) if nothing matches — "not found"
// is a normal outcome here, not an error, so callers in the service
// layer never need to know gorm.ErrRecordNotFound exists.
func (r *CompletionRepository) FindByTaskAndDate(taskID uint, date time.Time) (*models.Completion, error) {
	var completion models.Completion
	err := r.DB.Where("task_id = ? AND date = ?", taskID, date).First(&completion).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &completion, nil
}

func (r *CompletionRepository) Create(c *models.Completion) error {
	return r.DB.Create(c).Error
}

func (r *CompletionRepository) Delete(c *models.Completion) error {
	return r.DB.Delete(c).Error
}
