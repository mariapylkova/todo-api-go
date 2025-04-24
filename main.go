package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mariapylkova/todo-api-go/database"
	"github.com/mariapylkova/todo-api-go/handlers"
)

func main() {
	db := database.SetupDatabase()

	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks(db))
	r.POST("/tasks", handlers.CreateTask(db))
	r.DELETE("/tasks/:id", handlers.DeleteTask(db))

	r.Run()
}
