package repository_test

import (
	"os"
	"task-app/internal/model"
	"task-app/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupTempFile(t *testing.T) string {
	tmpfile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()
	return tmpfile.Name()
}

func TestSaveAndLoadTasks(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	repo := repository.NewTaskRepository(tempFile)

	tasks := []model.Task{
		{ID: 1, Description: "Test Task", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	err := repo.SaveTasks(tasks)
	assert.NoError(t, err)

	loadedTasks, err := repo.LoadTasks()
	assert.NoError(t, err)
	assert.Equal(t, len(tasks), len(loadedTasks))
	assert.Equal(t, tasks[0].Description, loadedTasks[0].Description)
}

func TestLoadTasks_FileNotExist(t *testing.T) {
	repo := repository.NewTaskRepository("nonexistent.json")
	tasks, err := repo.LoadTasks()

	assert.NoError(t, err)
	assert.Empty(t, tasks)
}

func TestLoadTasks_InvalidJSON(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	err := os.WriteFile(tempFile, []byte("invalid json"), 0644)
	assert.NoError(t, err)

	repo := repository.NewTaskRepository(tempFile)
	_, err = repo.LoadTasks()

	assert.Error(t, err)
}
