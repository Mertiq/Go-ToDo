package models

import "time"

type Task struct {
	Id        uint      `json:"id"`
	TaskItem  string    `json:"task_item"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
