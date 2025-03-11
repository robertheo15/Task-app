package repository

import (
	"os"
	"task-app/internal/model"
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

	repo := NewTaskRepository(tempFile)

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
	repo := NewTaskRepository("nonexistent.json")
	tasks, err := repo.LoadTasks()

	assert.NoError(t, err)
	assert.Empty(t, tasks)
}

func TestLoadTasks_InvalidJSON(t *testing.T) {
	tempFile := setupTempFile(t)
	defer os.Remove(tempFile)

	err := os.WriteFile(tempFile, []byte("invalid json"), 0644)
	assert.NoError(t, err)

	repo := NewTaskRepository(tempFile)
	_, err = repo.LoadTasks()

	assert.Error(t, err)
}
