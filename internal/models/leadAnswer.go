package models

import "time"

type LeadAnswer struct {
	ID         uint64    `db:"id" json:"id"`
	LeadID     uint64    `db:"lead_id" json:"leadId"`
	QuestionID uint64    `db:"question_id" json:"questionId"`
	OptionID   uint64    `db:"option_id" json:"optionId"`
	AnsweredAt time.Time `db:"answered_at" json:"answeredAt"`
}
