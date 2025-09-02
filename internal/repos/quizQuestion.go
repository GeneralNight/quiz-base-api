package repos

import (
	"context"
	"quiz-base-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type QuizQuestionRepository interface {
	CreateQuestion(ctx context.Context, quizID uint64, dto models.CreateQuestionDTO) (uint64, error)
	GetQuestions(ctx context.Context, quizID uint64) ([]models.QuizQuestion, error)
	UpdateQuestion(ctx context.Context, questionID uint64, dto models.CreateQuestionDTO) error
	DeleteQuestion(ctx context.Context, questionID uint64) error
}

type quizQuestionRepository struct{ db *sqlx.DB }

func NewQuizQuestionRepository(db *sqlx.DB) QuizQuestionRepository {
	return &quizQuestionRepository{db: db}
}

func (r *quizQuestionRepository) CreateQuestion(ctx context.Context, quizID uint64, dto models.CreateQuestionDTO) (uint64, error) {
	res, err := r.db.ExecContext(ctx,
		`INSERT INTO quiz_question (quiz_id, text) VALUES (?, ?)`,
		quizID, dto.Text,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r *quizQuestionRepository) GetQuestions(ctx context.Context, quizID uint64) ([]models.QuizQuestion, error) {
	var items []models.QuizQuestion
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM quiz_question WHERE quiz_id=? ORDER BY id ASC`, quizID)
	return items, err
}

func (r *quizQuestionRepository) UpdateQuestion(ctx context.Context, questionID uint64, dto models.CreateQuestionDTO) error {
	_, err := r.db.ExecContext(ctx, `UPDATE quiz_question SET text=? WHERE id=?`, dto.Text, questionID)
	return err
}

func (r *quizQuestionRepository) DeleteQuestion(ctx context.Context, questionID uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM quiz_question WHERE id=?`, questionID)
	return err
}
