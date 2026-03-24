package storage

import (
	"encoding/json"
	"os"
	"todo-app/internal/todo"
)

// LoadJSON загружает задачи из JSON файла.
func LoadJSON(filename string) ([]todo.Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []todo.Task{}, nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tasks []todo.Task
	if len(data) == 0 {
		return []todo.Task{}, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// saveJSON сохраняет задачи в JSON файл.
func SaveJSON(filename string, tasks []todo.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
