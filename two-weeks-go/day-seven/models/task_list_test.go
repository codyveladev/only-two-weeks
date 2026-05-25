package models

import "testing"

func TestGetNextId(t *testing.T) {
	tests := []struct {
		name       string
		taskList   TaskList
		expectedId int
	}{
		{
			name:       "empty task returns 1",
			taskList:   TaskList{Tasks: []Task{}},
			expectedId: 1,
		},
		{
			name:       "tasklist with 1 task returns 2",
			taskList:   TaskList{Tasks: []Task{{ID: 1, Title: "Buy Milk"}}},
			expectedId: 2,
		},
		{
			name: "non-contiguous returns next max ID",
			taskList: TaskList{Tasks: []Task{
				{ID: 1, Title: "Buy Milk"},
				{ID: 2, Title: "Get Oil Change"},
				{ID: 3, Title: "Dentist Appt"},
				{ID: 5, Title: "Laundry"},
			},
			},
			expectedId: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNextId(tt.taskList)
			if got != tt.expectedId {
				t.Errorf("got %d, expected %d", got, tt.expectedId)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	tests := []struct {
		name          string
		taskList      TaskList
		expectedSize  int
		expectedNewId int
	}{
		{
			name:          "add task to empty task list",
			taskList:      TaskList{Tasks: []Task{}},
			expectedSize:  1,
			expectedNewId: 1,
		},
		{
			name: "sequential id increment",
			taskList: TaskList{Tasks: []Task{
				{ID: 1, Title: "Buy Milk"},
				{ID: 2, Title: "Get Oil Change"},
				{ID: 3, Title: "Dentist Appt"},
				{ID: 4, Title: "Laundry"},
			},
			},
			expectedSize:  5,
			expectedNewId: 5,
		},
		{
			name: "non-sequential id increment",
			taskList: TaskList{Tasks: []Task{
				{ID: 1, Title: "Buy Milk"},
				{ID: 2, Title: "Get Oil Change"},
				{ID: 5, Title: "Dentist Appt"},
				{ID: 6, Title: "Laundry"},
			},
			},
			expectedSize:  5,
			expectedNewId: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := AddTask(tt.taskList, "new task")
			if len(got.Tasks) != tt.expectedSize {
				t.Errorf("got %d, expected %d", len(got.Tasks), tt.expectedSize)
			}

			newTask := got.Tasks[len(got.Tasks)-1]
			if newTask.ID != tt.expectedNewId {
				t.Errorf("new task ID got %d, expected %d", newTask.ID, tt.expectedSize)
			}

		})
	}

}

func TestMarkTaskComplete(t *testing.T) {
	tests := []struct {
		name                 string
		taskList             TaskList
		id                   int
		indexOfCompletedTask int
		shouldBeComplete     bool
		returnsError         bool
	}{
		{
			name:                 "invalid ID",
			taskList:             TaskList{Tasks: []Task{}},
			id:                   -1,
			indexOfCompletedTask: -1,
			shouldBeComplete:     false,
			returnsError:         true,
		},
		{
			name:                 "marks task complete",
			taskList:             TaskList{Tasks: []Task{{ID: 1, Title: "Buy Milk", Completed: false}}},
			id:                   1,
			indexOfCompletedTask: 0,
			shouldBeComplete:     true,
			returnsError:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompleteTask(tt.taskList, tt.id)
			if (err != nil) != tt.returnsError {
				t.Errorf("unexpected error for mark task complete %s", err)
				return
			}
			if tt.indexOfCompletedTask != -1 {
				completedTask := got.Tasks[tt.indexOfCompletedTask]
				if completedTask.Completed != tt.shouldBeComplete {
					t.Errorf("task not marked complete")
				}
			}

		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name         string
		taskList     TaskList
		idToDelete   int
		expectedSize int
		remainingIDs []int
		returnsError bool
	}{
		{
			name:         "delete from empty task",
			taskList:     TaskList{Tasks: []Task{}},
			idToDelete:   1,
			expectedSize: 0,
			remainingIDs: []int{},
			returnsError: true,
		},
		{
			name: "delete task from existing list",
			taskList: TaskList{Tasks: []Task{
				{ID: 1, Title: "Buy Milk"},
				{ID: 2, Title: "Get Oil Change"},
				{ID: 3, Title: "Dentist Appt"},
				{ID: 4, Title: "Laundry"},
			}},
			idToDelete:   3,
			expectedSize: 3,
			remainingIDs: []int{1, 2, 4},
			returnsError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteTask(tt.taskList, tt.idToDelete)
			if (err != nil) != tt.returnsError {
				t.Errorf("unexpected error for delete task %s", err)
				return
			}

			if len(got.Tasks) != tt.expectedSize {
				t.Errorf("got size %d, expected %d", len(got.Tasks), tt.expectedSize)
			}

			for i, task := range got.Tasks {
				if task.ID != tt.remainingIDs[i] {
					t.Errorf("expected ID %d at index %d, got %d", tt.remainingIDs[i], i, task.ID)
				}
			}
		})
	}
}

func BenchmarkGetNextId(b *testing.B) {
	taskList := TaskList{}

	for i := 0; i < 10000; i++ {
		taskList.Tasks = append(taskList.Tasks, NewTask(i, "task"))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		getNextId(taskList)
	}

}
