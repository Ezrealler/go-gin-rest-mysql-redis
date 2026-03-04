package model

import "time"

// MPost struct
type MPost struct {
	ID          int64     `db:"ID" json:"id" example:"1"`
	UserID      int64     `db:"user_id" json:"user_id" example:"1"`
	Title       string    `db:"title" json:"title" example:"xyz"`
	Description string    `db:"description" json:"description" example:"abcdescription"`
	Status      int       `db:"status" json:"status" example:"1"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// MPosts array of MPost type
type MPosts []MPost
