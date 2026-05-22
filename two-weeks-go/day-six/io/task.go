package io

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codyveladev/day-six/models"
)

func LoadTasks(filename string) (models.TaskList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return models.TaskList{}, fmt.Errorf("issue loading file %w", err)
	}
	var t []models.Task
	err = json.Unmarshal(data, &t)
	if err != nil {
		return models.TaskList{}, fmt.Errorf("issue parsing file %w", err)
	}
	return models.TaskList{Tasks: t}, nil
}

func SaveTasks(filename string, taskList models.TaskList) error {
	data, err := json.MarshalIndent(taskList.Tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
