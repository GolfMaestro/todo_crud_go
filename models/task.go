package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	PersonID    int       `json:"personId"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}
