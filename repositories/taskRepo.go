package repositories

import (
	"database/sql"
	"todo-backend/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) Create(task models.Task) error {
	_, err := r.DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	return err
}

func (r *TaskRepository) Update(id int, task models.Task) error {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?", task.Title, task.Description, task.Completed, id)
	return err
}

func (r *TaskRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
