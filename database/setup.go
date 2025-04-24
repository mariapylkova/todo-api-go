package database

import (
	"log"

	"github.com/mariapylkova/todo-api-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	db.AutoMigrate(&models.Task{})

	return db
}
