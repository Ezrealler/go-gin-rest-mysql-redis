package model

import "time"

type MCtegory struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}
