package models

import "time"

type Task struct {
	ID          int
	PersonID    int
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
}
