package repository

import (
	"errors"
	"focustracker/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) FindActiveByUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("user_id = ? AND is_archived = ?", userID, false).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Create(t *models.Task) error {
	return r.DB.Create(t).Error
}

func (r *TaskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.DB.First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(t *models.Task) error {
	return r.DB.Save(t).Error
}

func (r *TaskRepository) Archive(id uint) error {
	return r.DB.Model(&models.Task{}).Where("id = ?", id).Update("is_archived", true).Error
}
