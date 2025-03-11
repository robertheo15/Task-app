package service

import (
	"errors"
	"fmt"
	"task-app/internal/model"
	"task-app/internal/repository"
	"time"
)

type TaskService interface {
	AddTask(description string) error
	UpdateTask(id int, description string) error
	DeleteTask(id int) error
	MarkTask(id int, status string) error
	ListTasks(filter string) ([]model.Task, error)
}

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskRepository(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) AddTask(description string) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	newTask := model.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	err = s.repo.SaveTasks(tasks)
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)

	return nil
}

func (s *TaskServiceImpl) UpdateTask(id int, description string) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			err = s.repo.SaveTasks(tasks)
			if err != nil {
				return err
			}
			fmt.Println("Task updated successfully")
			return nil
		}
	}

	return errors.New("task not found")
}

func (s *TaskServiceImpl) DeleteTask(id int) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err = s.repo.SaveTasks(tasks)
			if err != nil {
				return err
			}
			fmt.Println("Task deleted successfully")
			return nil
		}
	}

	return errors.New("task not found")
}

func (s *TaskServiceImpl) MarkTask(id int, status string) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			err = s.repo.SaveTasks(tasks)
			if err != nil {
				return err
			}
			fmt.Printf("Task marked as %s\n", status)
			return nil
		}
	}

	return errors.New("task not found")
}

func (s *TaskServiceImpl) ListTasks(filter string) ([]model.Task, error) {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []model.Task
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}
