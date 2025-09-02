package models

import "time"

type QuizQuestion struct {
	ID        uint64    `db:"id" json:"id"`
	QuizID    uint64    `db:"quiz_id" json:"quizId"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type CreateQuestionDTO struct {
	Text    string            `json:"text" binding:"required"`
	Options []CreateOptionDTO `json:"options" binding:"required,dive,required"`
}
