package quiz

import "context"

type Service interface {
	CreateQuiz(ctx context.Context, dto CreateQuizDTO) (uint64, error)
	ListQuizzes(ctx context.Context) ([]Quiz, error)
	GetQuiz(ctx context.Context, id uint64) (*Quiz, error)
	UpdateQuiz(ctx context.Context, id uint64, dto CreateQuizDTO) error
	DeleteQuiz(ctx context.Context, id uint64) error
	UpdateQuizStatus(ctx context.Context, id uint64, status string) error

	AddQuestion(ctx context.Context, quizID uint64, dto CreateQuestionDTO) (uint64, error)
	ListQuestions(ctx context.Context, quizID uint64) ([]QuizQuestion, error)
	UpdateQuestion(ctx context.Context, questionID uint64, dto CreateQuestionDTO) error
	DeleteQuestion(ctx context.Context, questionID uint64) error

	ListOptions(ctx context.Context, questionID uint64) ([]QuestionOption, error)
	AddOption(ctx context.Context, questionID uint64, dto CreateOptionDTO) (uint64, error)
	UpdateOption(ctx context.Context, optionID uint64, dto CreateOptionDTO) error
	DeleteOption(ctx context.Context, optionID uint64) error
}

type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) CreateQuiz(ctx context.Context, dto CreateQuizDTO) (uint64, error) {
	return s.repo.CreateQuiz(ctx, dto)
}
func (s *service) ListQuizzes(ctx context.Context) ([]Quiz, error) {
	return s.repo.GetQuizzes(ctx)
}
func (s *service) GetQuiz(ctx context.Context, id uint64) (*Quiz, error) {
	return s.repo.GetQuiz(ctx, id)
}
func (s *service) UpdateQuiz(ctx context.Context, id uint64, dto CreateQuizDTO) error {
	return s.repo.UpdateQuiz(ctx, id, dto)
}
func (s *service) DeleteQuiz(ctx context.Context, id uint64) error {
	return s.repo.DeleteQuiz(ctx, id)
}
func (s *service) UpdateQuizStatus(ctx context.Context, id uint64, status string) error {
	return s.repo.UpdateQuizStatus(ctx, id, status)
}

func (s *service) AddQuestion(ctx context.Context, quizID uint64, dto CreateQuestionDTO) (uint64, error) {
	return s.repo.CreateQuestion(ctx, quizID, dto)
}
func (s *service) ListQuestions(ctx context.Context, quizID uint64) ([]QuizQuestion, error) {
	return s.repo.GetQuestions(ctx, quizID)
}
func (s *service) UpdateQuestion(ctx context.Context, questionID uint64, dto CreateQuestionDTO) error {
	return s.repo.UpdateQuestion(ctx, questionID, dto)
}
func (s *service) DeleteQuestion(ctx context.Context, questionID uint64) error {
	return s.repo.DeleteQuestion(ctx, questionID)
}

func (s *service) ListOptions(ctx context.Context, questionID uint64) ([]QuestionOption, error) {
	return s.repo.GetOptions(ctx, questionID)
}
func (s *service) AddOption(ctx context.Context, questionID uint64, dto CreateOptionDTO) (uint64, error) {
	return s.repo.CreateOption(ctx, questionID, dto)
}
func (s *service) UpdateOption(ctx context.Context, optionID uint64, dto CreateOptionDTO) error {
	return s.repo.UpdateOption(ctx, optionID, dto)
}
func (s *service) DeleteOption(ctx context.Context, optionID uint64) error {
	return s.repo.DeleteOption(ctx, optionID)
}
