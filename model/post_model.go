package model

import "time"

type MPost struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	CategoryID  int64     `json:"category_id"`
	Category    string    `json:"category,omitempty"` // 可选：联表返回分类名
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MPosts array of MPost type
type MPosts []MPost
