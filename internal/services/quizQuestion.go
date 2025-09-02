package services

import (
	"context"
	"quiz-base-api/internal/models"
	"quiz-base-api/internal/repos"
)

type QuizQuestionService interface {
	AddQuestion(ctx context.Context, quizID uint64, dto models.CreateQuestionDTO) (uint64, error)
	ListQuestions(ctx context.Context, quizID uint64) ([]models.QuizQuestion, error)
	UpdateQuestion(ctx context.Context, questionID uint64, dto models.CreateQuestionDTO) error
	DeleteQuestion(ctx context.Context, questionID uint64) error
}

type quizQuestionService struct{ repo repos.QuizQuestionRepository }

func NewQuizQuestionService(r repos.QuizQuestionRepository) QuizQuestionService {
	return &quizQuestionService{repo: r}
}

func (s *quizQuestionService) AddQuestion(ctx context.Context, quizID uint64, dto models.CreateQuestionDTO) (uint64, error) {
	return s.repo.CreateQuestion(ctx, quizID, dto)
}
func (s *quizQuestionService) ListQuestions(ctx context.Context, quizID uint64) ([]models.QuizQuestion, error) {
	return s.repo.GetQuestions(ctx, quizID)
}
func (s *quizQuestionService) UpdateQuestion(ctx context.Context, questionID uint64, dto models.CreateQuestionDTO) error {
	return s.repo.UpdateQuestion(ctx, questionID, dto)
}
func (s *quizQuestionService) DeleteQuestion(ctx context.Context, questionID uint64) error {
	return s.repo.DeleteQuestion(ctx, questionID)
}
