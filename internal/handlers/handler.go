package handlers

import (
	"focustracker/internal/repository"
	"focustracker/internal/services"
)

const DefaultUserID = 1

type Handler struct {
	TaskRepo       *repository.TaskRepository
	CompletionRepo *repository.CompletionRepository
	StreakService  *service.StreakService
	CompletionSvc  *service.CompletionService
}

func NewHandler(
	taskRepo *repository.TaskRepository,
	completionRepo *repository.CompletionRepository,
	streakService *service.StreakService,
	completionSvc *service.CompletionService,
) *Handler {
	return &Handler{
		TaskRepo:       taskRepo,
		CompletionRepo: completionRepo,
		StreakService:  streakService,
		CompletionSvc:  completionSvc,
	}
}
