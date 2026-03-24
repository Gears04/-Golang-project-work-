package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

func SaveCSV(filename string, tasks []todo.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"id", "description", "done"}); err != nil {
		return err
	}

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Done),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func LoadCSV(filename string) ([]todo.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []todo.Task

	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		done, err := strconv.ParseBool(record[2])
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, todo.Task{
			ID:          id,
			Description: record[1],
			Done:        done,
		})
	}

	return tasks, nil
}
