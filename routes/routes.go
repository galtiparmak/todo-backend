package routes

import (
	"todo-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controller *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}
