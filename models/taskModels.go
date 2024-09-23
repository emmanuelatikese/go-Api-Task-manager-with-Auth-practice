package model

import "time"

type TaskModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt time.Time
	Completed bool `json:"completed"`
}
