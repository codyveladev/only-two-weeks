package models

import "time"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"`
}

func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func (t *Task) Complete() {
	t.Completed = true
}
