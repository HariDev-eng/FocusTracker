package service

import (
	"time"

	"focustracker/internal/models"
	"focustracker/internal/repository"
)

type CompletionService struct {
	Repo *repository.CompletionRepository
}

func NewCompletionService(repo *repository.CompletionRepository) *CompletionService {
	return &CompletionService{Repo: repo}
}

// Toggle marks a task complete for a date, or un-marks it if it already
// was. Returns the resulting state.
func (s *CompletionService) Toggle(taskID uint, date time.Time) (bool, error) {
	existing, err := s.Repo.FindByTaskAndDate(taskID, date)
	if err != nil {
		return false, err
	}
	if existing != nil {
		if delErr := s.Repo.Delete(existing); delErr != nil {
			return false, delErr
		}
		return false, nil
	}

	completion := &models.Completion{TaskID: taskID, Date: date}
	if createErr := s.Repo.Create(completion); createErr != nil {
		return false, createErr
	}
	return true, nil
}
