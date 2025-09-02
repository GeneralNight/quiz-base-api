package models

import "time"

type QuestionOption struct {
	ID         uint64    `db:"id" json:"id"`
	QuestionID uint64    `db:"question_id" json:"questionId"`
	Text       string    `db:"text" json:"text"`
	IsCorrect  *bool     `db:"is_correct" json:"isCorrect,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
}

type CreateOptionDTO struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect *bool  `json:"isCorrect"`
}
