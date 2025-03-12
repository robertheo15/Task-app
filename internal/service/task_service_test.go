package service_test

import (
	"task-app/internal/model"
	"task-app/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) SaveTasks(tasks []model.Task) error {
	args := m.Called(tasks)
	return args.Error(0)
}

func (m *MockTaskRepository) LoadTasks() ([]model.Task, error) {
	args := m.Called()
	return args.Get(0).([]model.Task), args.Error(1)
}

func TestAddTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	existingTasks := []model.Task{}
	mockRepo.On("LoadTasks").Return(existingTasks, nil)
	mockRepo.On("SaveTasks", mock.Anything).Return(nil)

	taskService := service.NewTaskService(mockRepo)
	err := taskService.AddTask("Test task")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	existingTasks := []model.Task{{ID: 1, Description: "Old task", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	mockRepo.On("LoadTasks").Return(existingTasks, nil)
	mockRepo.On("SaveTasks", mock.Anything).Return(nil)

	taskService := service.NewTaskService(mockRepo)
	err := taskService.UpdateTask(1, "Updated task")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	existingTasks := []model.Task{{ID: 1, Description: "Task to delete", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	mockRepo.On("LoadTasks").Return(existingTasks, nil)
	mockRepo.On("SaveTasks", mock.Anything).Return(nil)

	taskService := service.NewTaskService(mockRepo)
	err := taskService.DeleteTask(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	existingTasks := []model.Task{{ID: 1, Description: "Task to mark", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	mockRepo.On("LoadTasks").Return(existingTasks, nil)
	mockRepo.On("SaveTasks", mock.Anything).Return(nil)

	taskService := service.NewTaskService(mockRepo)
	err := taskService.MarkTask(1, "done")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	existingTasks := []model.Task{{ID: 1, Description: "Task 1", Status: "todo"}, {ID: 2, Description: "Task 2", Status: "done"}}
	mockRepo.On("LoadTasks").Return(existingTasks, nil)

	taskService := service.NewTaskService(mockRepo)
	tasks, err := taskService.ListTasks("todo")

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "Task 1", tasks[0].Description)
	mockRepo.AssertExpectations(t)
}
