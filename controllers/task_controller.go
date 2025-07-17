package controllers

import (
	"net/http"
	// "strconv"
	"task_management/data"
	"task_management/models"
	// "task_management/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks,err := data.GetTasks()
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,  gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	
	task, err := data.GetTaskByID(id)
	if err !=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := data.AddTask(newTask)
	if err !=nil  {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "task already exits"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	

	
	err := data.DeleteTaskByID(id)
	if err !=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})

}
func UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	
	

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTaskByID(id, updatedTask)
	if err !=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTask)

}
