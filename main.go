package main

import (
	"todo-backend/config"
	"todo-backend/controllers"
	"todo-backend/repositories"
	"todo-backend/routes"
	"todo-backend/services"
)

func main() {
	config.ConnectDB()

	taskRepo := &repositories.TaskRepository{DB: config.DB}
	taskService := &services.TaskService{Repo: taskRepo}
	taskController := &controllers.TaskController{Service: taskService}

	router := routes.SetupRoutes(taskController)
	router.Run(":8080")
}
