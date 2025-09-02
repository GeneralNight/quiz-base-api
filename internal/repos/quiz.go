package repos

import (
	"context"
	"quiz-base-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type QuizRepository interface {
	CreateQuiz(ctx context.Context, dto models.CreateQuizDTO) (uint64, error)
	GetQuizzes(ctx context.Context) ([]models.Quiz, error)
	GetQuiz(ctx context.Context, id uint64) (*models.Quiz, error)
	UpdateQuiz(ctx context.Context, id uint64, dto models.CreateQuizDTO) error
	DeleteQuiz(ctx context.Context, id uint64) error
	UpdateQuizStatus(ctx context.Context, id uint64, status string) error
}

type quizRepository struct{ db *sqlx.DB }

func NewQuizRepository(db *sqlx.DB) QuizRepository { return &quizRepository{db: db} }

func (r *quizRepository) CreateQuiz(ctx context.Context, dto models.CreateQuizDTO) (uint64, error) {
	res, err := r.db.ExecContext(ctx,
		`INSERT INTO quiz (name, title, description, multiple_choice, answers_per_user, status) VALUES (?, ?, ?, ?, ?, ?)`,
		dto.Name, dto.Title, dto.Description, dto.MultipleChoice, dto.AnswersPerUser, dto.Status,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r *quizRepository) GetQuizzes(ctx context.Context) ([]models.Quiz, error) {
	var items []models.Quiz
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM quiz ORDER BY id DESC`)
	return items, err
}

func (r *quizRepository) GetQuiz(ctx context.Context, id uint64) (*models.Quiz, error) {
	var q models.Quiz
	err := r.db.GetContext(ctx, &q, `SELECT * FROM quiz WHERE id=?`, id)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *quizRepository) UpdateQuiz(ctx context.Context, id uint64, dto models.CreateQuizDTO) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE quiz SET name=?, title=?, description=?, multiple_choice=?, answers_per_user=?, status=? WHERE id=?`,
		dto.Name, dto.Title, dto.Description, dto.MultipleChoice, dto.AnswersPerUser, dto.Status, id,
	)
	return err
}

func (r *quizRepository) DeleteQuiz(ctx context.Context, id uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM quiz WHERE id=?`, id)
	return err
}

func (r *quizRepository) UpdateQuizStatus(ctx context.Context, id uint64, status string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE quiz SET status=? WHERE id=?`, status, id)
	return err
}
