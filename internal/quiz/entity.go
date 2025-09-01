package quiz

import "time"

type Quiz struct {
	ID             uint64    `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Title          string    `db:"title" json:"title"`
	Description    *string   `db:"description" json:"description,omitempty"`
	MultipleChoice bool      `db:"multiple_choice" json:"multipleChoice"`
	AnswersPerUser *uint     `db:"answers_per_user" json:"answersPerUser,omitempty"`
	Status         string    `db:"status" json:"status"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
}

type QuizQuestion struct {
	ID        uint64    `db:"id" json:"id"`
	QuizID    uint64    `db:"quiz_id" json:"quizId"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type QuestionOption struct {
	ID         uint64    `db:"id" json:"id"`
	QuestionID uint64    `db:"question_id" json:"questionId"`
	Text       string    `db:"text" json:"text"`
	IsCorrect  *bool     `db:"is_correct" json:"isCorrect,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
}

type LeadUser struct {
	ID        uint64    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type LeadAnswer struct {
	ID         uint64    `db:"id" json:"id"`
	LeadID     uint64    `db:"lead_id" json:"leadId"`
	QuestionID uint64    `db:"question_id" json:"questionId"`
	OptionID   uint64    `db:"option_id" json:"optionId"`
	AnsweredAt time.Time `db:"answered_at" json:"answeredAt"`
}

// payloads
type CreateQuizDTO struct {
	Name           string  `json:"name" binding:"required"`
	Title          string  `json:"title" binding:"required"`
	Description    *string `json:"description"`
	MultipleChoice bool    `json:"multipleChoice"`
	AnswersPerUser *uint   `json:"answersPerUser"`
	Status         string  `json:"status"`
}

type CreateQuestionDTO struct {
	Text string `json:"text" binding:"required"`
}

type CreateOptionDTO struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect *bool  `json:"isCorrect"`
}
