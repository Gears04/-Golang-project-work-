# Todo App (CLI)

Консольное приложение для управления списком задач (Todo List), написанное на Go.
Поддерживает добавление, просмотр, завершение, удаление задач, а также импорт и экспорт данных в форматах JSON и CSV.

Проект реализован как CLI-приложение с использованием стандартного пакета `flag` и подкоманд.

---

## Структура проекта

todo-app/
├── cmd/
│ └── todo/
│ └── main.go # точка входа CLI-приложения
├── internal/
│ ├── todo/ # бизнес-логика (задачи)
│ └── storage/ # работа с JSON и CSV
├── tasks.json # основное хранилище задач
├── go.mod
└── README.md

## КОманды Todo CLI

1) Добавить задачу:
2) 
 - todo add --desc "Текст задачи"
  
3) Показать список задач:

 - todo list --filter all
  
 ~ режим вывода:

 all — все задачи

 done — только выполненные

 pending — только невыполненные

4) Отметить задачу выполненой:
   
 - todo complete --id 1
 ~ --id — идентификатор задачи (обязательно > 0))

5) Удалить задачу:
 
 - todo delete --id 1

6) Экспортировать задачи и файл:
 
 - todo export --format --out backup.json
 - todo export --format csv --out backup.csv
  ~ --format (обязательный) — json или csv
  ~ --out (обязательный) — путь к выходному файлу

7) Импортировать задачи из файла
  
  - todo load --format json --file backup.json
  - todo load --format csv --file backup.csv

   ~--format (обязательный) — json или csv

   ~--file (обязательный) — путь к входному файлу
# Golang project work - Мини-приложение на Go
Todo-app
