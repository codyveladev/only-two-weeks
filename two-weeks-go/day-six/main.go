package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codyveladev/day-six/config"
	"github.com/codyveladev/day-six/io"
	"github.com/codyveladev/day-six/models"
)

func convertInputToId(input string) int {
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error: ", err)
		return -1
	}
	return id
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command> [args]")
		return
	}

	taskList, err := io.LoadTasks(config.TASK_FILE_NAME)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	op := os.Args[1]

	switch op {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add \"task title\"")
			return
		}
		title := os.Args[2]
		taskList = models.AddTask(taskList, title)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . done [id]")
			return
		}
		id := convertInputToId(os.Args[2])
		taskList, err = models.CompleteTask(taskList, id)
		if err != nil {
			fmt.Println("unable to mark task complete ", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . delete [id]")
			return
		}
		id := convertInputToId(os.Args[2])
		taskList, err = models.DeleteTask(taskList, id)
		if err != nil {
			fmt.Println("unable to delete ", err)
		}
	case "list":
		models.ListTasks(taskList)
	default:
		fmt.Println("invalid operation")
	}
	err = io.SaveTasks(config.TASK_FILE_NAME, taskList)
	if err != nil {
		fmt.Println("error saving file ", err)
	}
}
