package services

import (
	"context"
	"quiz-base-api/internal/models"
	"quiz-base-api/internal/repos"
)

type QuizService interface {
	CreateQuiz(ctx context.Context, dto models.CreateQuizDTO) (uint64, error)
	ListQuizzes(ctx context.Context) ([]models.Quiz, error)
	GetQuiz(ctx context.Context, id uint64) (*models.Quiz, error)
	UpdateQuiz(ctx context.Context, id uint64, dto models.CreateQuizDTO) error
	DeleteQuiz(ctx context.Context, id uint64) error
	UpdateQuizStatus(ctx context.Context, id uint64, status string) error
}

type quizService struct{ repo repos.QuizRepository }

func NewQuizService(r repos.QuizRepository) QuizService { return &quizService{repo: r} }

func (s *quizService) CreateQuiz(ctx context.Context, dto models.CreateQuizDTO) (uint64, error) {
	return s.repo.CreateQuiz(ctx, dto)
}
func (s *quizService) ListQuizzes(ctx context.Context) ([]models.Quiz, error) {
	return s.repo.GetQuizzes(ctx)
}
func (s *quizService) GetQuiz(ctx context.Context, id uint64) (*models.Quiz, error) {
	return s.repo.GetQuiz(ctx, id)
}
func (s *quizService) UpdateQuiz(ctx context.Context, id uint64, dto models.CreateQuizDTO) error {
	return s.repo.UpdateQuiz(ctx, id, dto)
}
func (s *quizService) DeleteQuiz(ctx context.Context, id uint64) error {
	return s.repo.DeleteQuiz(ctx, id)
}
func (s *quizService) UpdateQuizStatus(ctx context.Context, id uint64, status string) error {
	return s.repo.UpdateQuizStatus(ctx, id, status)
}
