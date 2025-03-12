package main

import (
	"fmt"
	"os"
	"strconv"
	"task-app/internal/repository"
	"task-app/internal/service"
)

const taskFile = "tasks.json"

func main() {
	repo := repository.NewTaskRepository(taskFile)
	svc := service.NewTaskService(repo)

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}
	command := os.Args[1]
	showMenu(command, svc)
}

func showMenu(command string, svc service.TaskService) {
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := os.Args[2]
		err := svc.AddTask(description)
		if err != nil {
			return
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <description>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		err := svc.UpdateTask(id, description)
		if err != nil {
			return
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := svc.DeleteTask(id)
		if err != nil {
			return
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := svc.MarkTask(id, "in-progress")
		if err != nil {
			return
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := svc.MarkTask(id, "done")
		if err != nil {
			return
		}

	case "list":
		if len(os.Args) == 3 {
			tasks, err := svc.ListTasks(os.Args[2])
			if err != nil {
				return
			}

			for _, task := range tasks {
				fmt.Printf("[%d] %s - %s\n", task.ID, task.Description, task.Status)
			}

		} else {
			tasks, err := svc.ListTasks("")
			if err != nil {
				return
			}

			for _, task := range tasks {
				fmt.Printf("[%d] %s - %s\n", task.ID, task.Description, task.Status)
			}
		}
	default:
		fmt.Println("Unknown command")
	}
}
