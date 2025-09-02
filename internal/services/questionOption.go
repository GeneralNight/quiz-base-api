package services

import (
	"context"
	"quiz-base-api/internal/models"
	"quiz-base-api/internal/repos"
)

type QuestionOptionService interface {
	ListOptions(ctx context.Context, questionID uint64) ([]models.QuestionOption, error)
	AddOption(ctx context.Context, questionID uint64, dto models.CreateOptionDTO) (uint64, error)
	UpdateOption(ctx context.Context, optionID uint64, dto models.CreateOptionDTO) error
	DeleteOption(ctx context.Context, optionID uint64) error
}

type questionOptionService struct {
	repo repos.QuestionOptionRepository
}

func NewQuestionOptionService(r repos.QuestionOptionRepository) QuestionOptionService {
	return &questionOptionService{repo: r}
}

func (s *questionOptionService) ListOptions(ctx context.Context, questionID uint64) ([]models.QuestionOption, error) {
	return s.repo.GetOptions(ctx, questionID)
}
func (s *questionOptionService) AddOption(ctx context.Context, questionID uint64, dto models.CreateOptionDTO) (uint64, error) {
	return s.repo.CreateOption(ctx, questionID, dto)
}
func (s *questionOptionService) UpdateOption(ctx context.Context, optionID uint64, dto models.CreateOptionDTO) error {
	return s.repo.UpdateOption(ctx, optionID, dto)
}
func (s *questionOptionService) DeleteOption(ctx context.Context, optionID uint64) error {
	return s.repo.DeleteOption(ctx, optionID)
}
