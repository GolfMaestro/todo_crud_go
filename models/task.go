package models

import "time"

type Task struct {
	ID         int       `json:"id"`
	PersonID   int       `json:"personId"`
	Title      string    `json:"title"`
	IsComplete bool      `json:"isComplete"`
	CreatedAt  time.Time `json:"createdAt"`
}
