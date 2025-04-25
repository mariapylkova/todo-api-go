#  todo-api

Простой REST API для управления задачами, написанный на Go с использованием фреймворка Gin и базы данных SQLite.

##  Возможности

- Создание задачи
- Получение списка задач
- Удаление задачи

##  Технологии

- Go
- Gin
- GORM
- SQLite
- Docker

##  Структура проекта

```
todo-api/ 
├── main.go 
├── go.mod 
├── models/ 
    │ └── task.go 
├── handlers/ 
    │ └── task.go
├── database/ 
    │ └── setup.go 
├── Dockerfile
```

## Установка и запуск

### Локально

```bash
git clone https://github.com/mariapylkova/todo-api.git
cd todo-api
go mod tidy
go run main.go
```
### С помощью Docker
```
docker build -t todo-api .
docker run -p 8080:8080 todo-api
```
##  Эндпоинты

| Метод   | URL            | Описание                    | Тело запроса (если нужно)     |
|---------|----------------|-----------------------------|-------------------------------|
| GET     | `/tasks`       | Получить список всех задач  | –                             |
| POST    | `/tasks`       | Создать новую задачу        | `{ "title": "New Task" }`     |
| DELETE  | `/tasks/:id`   | Удалить задачу по ID        | –                             |
