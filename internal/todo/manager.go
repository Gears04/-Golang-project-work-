package todo

import (
	"errors"
)

// Add добовляет новую задачу с заданным описанием.

func Add(tasks []Task, desc string) []Task {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	newTask := Task{
		ID:          maxID + 1,
		Description: desc,
		Done:        false,
	}
	return append(tasks, newTask)
}

// List возвращает список задач с учетом фильтра

func List(tasks []Task, filter string) []Task {
	if filter == "" || filter == "all" {
		return tasks
	}

	var result []Task
	for _, task := range tasks {
		switch filter {
		case "done":
			if task.Done {
				result = append(result, task)
			}

		case "pending":
			if !task.Done {
				result = append(result, task)
			}
		}
	}

	return result
}

// Complete отмечает задачу выполненной по id. Возвращает ошибку, если задача не найдена.

func Complete(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			return tasks, nil
		}
	}

	return tasks, errors.New("task not found")
}

// Delete e удаляет задачу по id.

func Delete(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}

	return tasks, errors.New("task not found")
}
