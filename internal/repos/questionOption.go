package repos

import (
	"context"
	"quiz-base-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type QuestionOptionRepository interface {
	GetOptions(ctx context.Context, questionID uint64) ([]models.QuestionOption, error)
	CreateOption(ctx context.Context, questionID uint64, dto models.CreateOptionDTO) (uint64, error)
	UpdateOption(ctx context.Context, optionID uint64, dto models.CreateOptionDTO) error
	DeleteOption(ctx context.Context, optionID uint64) error
}

type questionOptionRepository struct{ db *sqlx.DB }

func NewQuestionOptionRepository(db *sqlx.DB) QuestionOptionRepository {
	return &questionOptionRepository{db: db}
}

func (r *questionOptionRepository) GetOptions(ctx context.Context, questionID uint64) ([]models.QuestionOption, error) {
	var items []models.QuestionOption
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM question_option WHERE question_id=? ORDER BY id ASC`, questionID)
	return items, err
}

func (r questionOptionRepository) CreateOption(ctx context.Context, questionID uint64, dto models.CreateOptionDTO) (uint64, error) {
	res, err := r.db.ExecContext(ctx,
		`INSERT INTO question_option (question_id, text, is_correct) VALUES (?, ?, ?)`,
		questionID, dto.Text, dto.IsCorrect,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r questionOptionRepository) UpdateOption(ctx context.Context, optionID uint64, dto models.CreateOptionDTO) error {
	_, err := r.db.ExecContext(ctx, `UPDATE question_option SET text=?, is_correct=? WHERE id=?`, dto.Text, dto.IsCorrect, optionID)
	return err
}

func (r questionOptionRepository) DeleteOption(ctx context.Context, optionID uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM question_option WHERE id=?`, optionID)
	return err
}
