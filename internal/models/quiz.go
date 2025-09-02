package models

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

// payloads
type CreateQuizDTO struct {
	Name           string  `json:"name" binding:"required"`
	Title          string  `json:"title" binding:"required"`
	Description    *string `json:"description"`
	MultipleChoice bool    `json:"multipleChoice"`
	AnswersPerUser *uint   `json:"answersPerUser"`
	Status         string  `json:"status"`
}
