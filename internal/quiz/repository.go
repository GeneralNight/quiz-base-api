package quiz

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateQuiz(ctx context.Context, dto CreateQuizDTO) (uint64, error)
	GetQuizzes(ctx context.Context) ([]Quiz, error)
	GetQuiz(ctx context.Context, id uint64) (*Quiz, error)
	UpdateQuiz(ctx context.Context, id uint64, dto CreateQuizDTO) error
	DeleteQuiz(ctx context.Context, id uint64) error
	UpdateQuizStatus(ctx context.Context, id uint64, status string) error

	CreateQuestion(ctx context.Context, quizID uint64, dto CreateQuestionDTO) (uint64, error)
	GetQuestions(ctx context.Context, quizID uint64) ([]QuizQuestion, error)
	UpdateQuestion(ctx context.Context, questionID uint64, dto CreateQuestionDTO) error
	DeleteQuestion(ctx context.Context, questionID uint64) error

	GetOptions(ctx context.Context, questionID uint64) ([]QuestionOption, error)
	CreateOption(ctx context.Context, questionID uint64, dto CreateOptionDTO) (uint64, error)
	UpdateOption(ctx context.Context, optionID uint64, dto CreateOptionDTO) error
	DeleteOption(ctx context.Context, optionID uint64) error
}

type repository struct{ db *sqlx.DB }

func NewRepository(db *sqlx.DB) Repository { return &repository{db: db} }

func (r *repository) CreateQuiz(ctx context.Context, dto CreateQuizDTO) (uint64, error) {
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

func (r *repository) GetQuizzes(ctx context.Context) ([]Quiz, error) {
	var items []Quiz
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM quiz ORDER BY id DESC`)
	return items, err
}

func (r *repository) GetQuiz(ctx context.Context, id uint64) (*Quiz, error) {
	var q Quiz
	err := r.db.GetContext(ctx, &q, `SELECT * FROM quiz WHERE id=?`, id)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *repository) UpdateQuiz(ctx context.Context, id uint64, dto CreateQuizDTO) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE quiz SET name=?, title=?, description=?, multiple_choice=?, answers_per_user=?, status=? WHERE id=?`,
		dto.Name, dto.Title, dto.Description, dto.MultipleChoice, dto.AnswersPerUser, dto.Status, id,
	)
	return err
}

func (r *repository) DeleteQuiz(ctx context.Context, id uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM quiz WHERE id=?`, id)
	return err
}

func (r *repository) UpdateQuizStatus(ctx context.Context, id uint64, status string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE quiz SET status=? WHERE id=?`, status, id)
	return err
}

func (r *repository) CreateQuestion(ctx context.Context, quizID uint64, dto CreateQuestionDTO) (uint64, error) {
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

func (r *repository) GetQuestions(ctx context.Context, quizID uint64) ([]QuizQuestion, error) {
	var items []QuizQuestion
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM quiz_question WHERE quiz_id=? ORDER BY id ASC`, quizID)
	return items, err
}

func (r *repository) UpdateQuestion(ctx context.Context, questionID uint64, dto CreateQuestionDTO) error {
	_, err := r.db.ExecContext(ctx, `UPDATE quiz_question SET text=? WHERE id=?`, dto.Text, questionID)
	return err
}

func (r *repository) DeleteQuestion(ctx context.Context, questionID uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM quiz_question WHERE id=?`, questionID)
	return err
}

func (r *repository) GetOptions(ctx context.Context, questionID uint64) ([]QuestionOption, error) {
	var items []QuestionOption
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM question_option WHERE question_id=? ORDER BY id ASC`, questionID)
	return items, err
}

func (r *repository) CreateOption(ctx context.Context, questionID uint64, dto CreateOptionDTO) (uint64, error) {
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

func (r *repository) UpdateOption(ctx context.Context, optionID uint64, dto CreateOptionDTO) error {
	_, err := r.db.ExecContext(ctx, `UPDATE question_option SET text=?, is_correct=? WHERE id=?`, dto.Text, dto.IsCorrect, optionID)
	return err
}

func (r *repository) DeleteOption(ctx context.Context, optionID uint64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM question_option WHERE id=?`, optionID)
	return err
}
