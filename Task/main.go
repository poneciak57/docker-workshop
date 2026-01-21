package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&Task{})
}

func indexHandler(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Tasks": tasks,
	})
}

func taskGetHandler(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(200, gin.H{"tasks": tasks})
}

func taskPostHandler(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&newTask); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created", "task": newTask})
}

func taskDeleteHandler(c *gin.Context) {
	id := c.Param("id")

	// Delete logic with Gorm
	if result := db.Delete(&Task{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func main() {
	initDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)

	// API Routes
	router.GET("/api/v1/tasks", taskGetHandler)
	router.POST("/api/v1/tasks", taskPostHandler)
	router.DELETE("/api/v1/tasks/:id", taskDeleteHandler)

	router.Run(":8080")

}
