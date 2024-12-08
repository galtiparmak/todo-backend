package controllers

import (
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Service *services.TaskService
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.Service.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.Service.CreateTask(task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.Service.UpdateTask(id, task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Service.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
