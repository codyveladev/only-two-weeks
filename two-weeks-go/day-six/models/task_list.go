package models

import "fmt"

type TaskList struct {
	Tasks []Task
}

func AddTask(taskList TaskList, title string) TaskList {
	newTask := NewTask(len(taskList.Tasks)+1, title)
	taskList.Tasks = append(taskList.Tasks, newTask)
	return taskList
}

func findTaskIndexById(taskList TaskList, id int) (int, error) {
	for i, task := range taskList.Tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("unable to find task with id %d", id)
}

func CompleteTask(taskList TaskList, id int) (TaskList, error) {
	index, err := findTaskIndexById(taskList, id)
	if err != nil {
		return taskList, err
	}
	taskList.Tasks[index].Complete()
	return taskList, nil
}

func DeleteTask(taskList TaskList, id int) (TaskList, error) {
	index, err := findTaskIndexById(taskList, id)
	if err != nil {
		return taskList, err
	}
	taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
	return taskList, nil
}

func ListTasks(taskList TaskList) {
	fmt.Printf("TASK LIST\n")
	for _, task := range taskList.Tasks {
		fmt.Printf("%d - %s - Completed: %t - %s\n", task.ID, task.Title, task.Completed, task.CreatedAt)
	}
}
