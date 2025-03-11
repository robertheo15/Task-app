package service

import (
	"errors"
	"task-app/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockTaskRepository struct {
	tasks []model.Task
	err   error
}

func (m *mockTaskRepository) SaveTasks(tasks []model.Task) error {
	if m.err != nil {
		return m.err
	}
	m.tasks = tasks
	return nil
}

func (m *mockTaskRepository) LoadTasks() ([]model.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.tasks, nil
}

func TestAddTask(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskRepository(mockRepo)

	err := service.AddTask("Test Task")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(mockRepo.tasks))
	assert.Equal(t, "Test Task", mockRepo.tasks[0].Description)
}

func TestAddTask_RepositoryError(t *testing.T) {
	mockRepo := &mockTaskRepository{err: errors.New("database error")}
	service := NewTaskRepository(mockRepo)

	err := service.AddTask("Test Task")
	assert.Error(t, err)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := &mockTaskRepository{
		tasks: []model.Task{
			{ID: 1, Description: "Old Task", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	service := NewTaskRepository(mockRepo)

	err := service.UpdateTask(1, "Updated Task")
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", mockRepo.tasks[0].Description)
}

func TestUpdateTask_NotFound(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskRepository(mockRepo)

	err := service.UpdateTask(1, "Updated Task")
	assert.Error(t, err)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := &mockTaskRepository{
		tasks: []model.Task{
			{ID: 1, Description: "Task to Delete", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	service := NewTaskRepository(mockRepo)

	err := service.DeleteTask(1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(mockRepo.tasks))
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskRepository(mockRepo)

	err := service.DeleteTask(1)
	assert.Error(t, err)
}

func TestMarkTask(t *testing.T) {
	mockRepo := &mockTaskRepository{
		tasks: []model.Task{
			{ID: 1, Description: "Task to Mark", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	service := NewTaskRepository(mockRepo)

	err := service.MarkTask(1, "done")
	assert.NoError(t, err)
	assert.Equal(t, "done", mockRepo.tasks[0].Status)
}

func TestMarkTask_NotFound(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskRepository(mockRepo)

	err := service.MarkTask(1, "done")
	assert.Error(t, err)
}

func TestListTasks(t *testing.T) {
	mockRepo := &mockTaskRepository{
		tasks: []model.Task{
			{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Description: "Task 2", Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	service := NewTaskRepository(mockRepo)

	tasks, err := service.ListTasks("done")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, "done", tasks[0].Status)
}

func TestListTasks_All(t *testing.T) {
	mockRepo := &mockTaskRepository{
		tasks: []model.Task{
			{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Description: "Task 2", Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	service := NewTaskRepository(mockRepo)

	tasks, err := service.ListTasks("")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(tasks))
}
