package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var counter int

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"status"`
}

func connectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}

func showAllAvailableTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

func showCompletedTasks(c *gin.Context) {
	var tasks []Task
	db.Where("is_done = ?", true).Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

func showActiveTasks(c *gin.Context) {
	var tasks []Task
	db.Where("is_done = ?", false).Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found, impossible to delete"})
		return
	}

	db.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"task": "successfully removed"})
}

func getTaskById(id string) (*Task, error) {
	var tasks []Task
	db.Find(&tasks)
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("task not found")
}

func taskById(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	counter++

	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query param"})
		return
	}

	task, err := getTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	fmt.Println(counter)
	fmt.Println(task.IsDone)

	if counter%2 != 0 {
		task.IsDone = true
		if err := db.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}
	} else {
		task.IsDone = false
		if err := db.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, task)
}

func main() {
	db = connectDb()
	newRouter := gin.Default()
	newRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "server active, db connected"})
	})
	newRouter.GET("/tasks", showAllAvailableTasks)
	newRouter.GET("/tasks/active", showActiveTasks)
	newRouter.GET("/tasks/completed", showCompletedTasks)
	newRouter.GET("/tasks/:id", taskById)
	newRouter.POST("/tasks", createTask)
	newRouter.PATCH("/update", updateTask)
	newRouter.DELETE("/tasks/del/:id", deleteTask)
	newRouter.Run()
}
