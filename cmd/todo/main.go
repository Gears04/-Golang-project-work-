package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {
	const filename = "tasks.json"

	// Загрузка задач из файла
	tasks, err := storage.LoadJSON(filename)
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		runAdd(tasks, filename, os.Args[2:])
	case "list":
		runList(tasks, os.Args[2:])
	case "complete":
		runComplete(tasks, filename, os.Args[2:])
	case "delete":
		runDelete(tasks, filename, os.Args[2:])
	case "export":
		runExport(tasks, os.Args[2:])
	case "load":
		runLoad(filename, os.Args[2:])
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  todo add --desc \"описание\"")
	fmt.Println("  todo list --filter all|done|pending")
	fmt.Println("  todo complete --id 1")
	fmt.Println("  todo delete --id 1")
	fmt.Println("  todo export --format csv|json --out file")
	fmt.Println("  todo load --format csv|json --file file")
}

func runAdd(tasks []todo.Task, filename string, args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)

	desc := fs.String("desc", "", "описание задачи")

	fs.Parse(args)

	if strings.TrimSpace(*desc) == "" {
		fmt.Println("Флаг --desc обязателен")
		return
	}

	tasks = todo.Add(tasks, *desc)

	if err := storage.SaveJSON(filename, tasks); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Println("Задача добавлена.")
}

func runList(tasks []todo.Task, args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)

	filter := fs.String("filter", "all", "all|done|pending")

	_ = fs.Parse(args)

	// Валидация значений фильтра

	switch *filter {
	case "all", "done", "pending":
		// ok
	default:
		fmt.Println("Некорректный --filter. Доступно: all|done|pending")
		return
	}

	// Печать с применением фильтра
	for _, task := range tasks {
		// Применяем фильтр:
		// - done: только выполнение
		// - pending: только не выполненные
		// - all: всё
		if *filter == "done" && !task.Done {
			continue
		}
		if *filter == "pending" && task.Done {
			continue
		}

		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Description)
	}
}

func runComplete(tasks []todo.Task, filename string, args []string) {
	fs := flag.NewFlagSet("complete", flag.ExitOnError)
	id := fs.Int("id", 0, "ID задачи")

	_ = fs.Parse(args)

	if *id <= 0 {
		fmt.Println("Флаг --id обязателен и должен быть > 0")
		return
	}

	updated, err := todo.Complete(tasks, *id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := storage.SaveJSON(filename, updated); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Println("Задача отмечена как выполненная.")
}

func runDelete(tasks []todo.Task, filename string, args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.Int("id", 0, "ID задачи")

	_ = fs.Parse(args)

	if *id <= 0 {
		fmt.Println("Флаг --id обязателен")
		return
	}

	updated, err := todo.Delete(tasks, *id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := storage.SaveJSON(filename, updated); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Println("Задача удалена.")
}

func runExport(tasks []todo.Task, args []string) {
	fs := flag.NewFlagSet("export", flag.ExitOnError)

	format := fs.String("format", "", "csv|json")
	out := fs.String("out", "", "output file path")

	_ = fs.Parse(args)

	if strings.TrimSpace(*format) == "" {
		fmt.Println("Флаг --format обязателен (csv|json)")
		return
	}

	if strings.TrimSpace(*out) == "" {
		fmt.Println("Флаг --out обязателен")
		return
	}

	f := strings.ToLower(strings.TrimSpace(*format))

	switch f {
	case "json":
		if err := storage.SaveJSON(*out, tasks); err != nil {
			fmt.Println("Ошибка экспорта JSON:", err)
			return
		}
		fmt.Println("Экспортировано в JSON:", *out)

	case "csv":
		if err := storage.SaveCSV(*out, tasks); err != nil {
			fmt.Println("Ошибка экспорта CSV:", err)
			return
		}
		fmt.Println("Экспортировано в CSV:", *out)

	default:
		fmt.Println("Некорректный --format. Доступно: csv|json")
		return
	}
}

func runLoad(mainFilename string, args []string) {
	fs := flag.NewFlagSet("load", flag.ExitOnError)

	format := fs.String("format", "", "csv|json")
	file := fs.String("file", "", "input file path")

	_ = fs.Parse(args)

	if strings.TrimSpace(*format) == "" {
		fmt.Println("Флаг --format обязателен (csv|json)")
		return
	}
	if strings.TrimSpace(*file) == "" {
		fmt.Println("Флаг --file обязателен")
		return
	}

	var (
		tasks []todo.Task
		err   error
	)

	switch *format {
	case "json":
		tasks, err = storage.LoadJSON(*file)
		if err != nil {
			fmt.Println("Ошибка загрузки JSON:", err)
			return
		}

	case "csv":
		tasks, err = storage.LoadCSV(*file)
		if err != nil {
			fmt.Println("Ошибка загрузки CSV:", err)
			return
		}

	default:
		fmt.Println("Некорректный --format. Допустимо: csv|json")
		return
	}

	// ВАЖНО: сохраняем загруженные задачи в основной файл хранилища (tasks.json)
	if err := storage.SaveJSON(mainFilename, tasks); err != nil {
		fmt.Println("Ошибка сохранения в основное хранилище:", err)
		return
	}

	fmt.Println("Задачи загружены из", *file, "и сохранены в", mainFilename)
}
