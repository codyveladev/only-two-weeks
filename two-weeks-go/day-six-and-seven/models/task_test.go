package models

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask(1, "Buy Milk")

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Title != "Buy Milk" {
		t.Errorf("expected title 'Buy Milk' got %s", task.Title)
	}

	if task.Completed {
		t.Errorf("expected new task to not be completed")
	}

}
