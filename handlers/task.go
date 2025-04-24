package handlers

import (
	"net/http"
	"todo-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []models.Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	}
}

func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&task)
		c.JSON(http.StatusCreated, task)
	}
}

func DeleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Task{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	}
}
