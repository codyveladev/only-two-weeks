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

func TestCompleteTask(t *testing.T) {
	tests := []struct {
		name     string
		task     Task
		wantDone bool
	}{
		{
			name:     "marks incomplete task as done",
			task:     NewTask(1, "Buy Milk"),
			wantDone: true,
		},
		{
			name:     "already done task stays done",
			task:     Task{ID: 2, Completed: true},
			wantDone: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.task.Complete()
			if tt.task.Completed != tt.wantDone {
				t.Errorf("got %v, want %v", tt.task.Completed, tt.wantDone)
			}
		})
	}
}
