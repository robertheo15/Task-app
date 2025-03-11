package repository

import (
	"encoding/json"
	"errors"
	"os"
	"task-app/internal/model"
)

type TaskRepository interface {
	SaveTasks(taskList []model.Task) error
	LoadTasks() ([]model.Task, error)
}

type TaskRepositoryImpl struct {
	FilePath string
}

func NewTaskRepository(filepath string) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		FilePath: filepath,
	}
}

func (repo *TaskRepositoryImpl) SaveTasks(tasks []model.Task) error {
	taskList := map[string][]model.Task{"tasks": tasks}
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(repo.FilePath, data, 0644)
}

func (repo *TaskRepositoryImpl) LoadTasks() ([]model.Task, error) {
	var taskWrapper struct {
		Tasks []model.Task `json:"tasks"`
	}

	var tasks []model.Task

	if _, err := os.Stat(repo.FilePath); errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	}

	data, err := os.ReadFile(repo.FilePath)
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(data, &taskWrapper)
	if err != nil {
		return nil, err
	}

	return taskWrapper.Tasks, err
}
