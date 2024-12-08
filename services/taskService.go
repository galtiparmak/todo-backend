package services

import (
	"todo-backend/models"
	"todo-backend/repositories"
)

type TaskService struct {
	Repo *repositories.TaskRepository
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) CreateTask(task models.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) UpdateTask(id int, task models.Task) error {
	return s.Repo.Update(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.Repo.Delete(id)
}
