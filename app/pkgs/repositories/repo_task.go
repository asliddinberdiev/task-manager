package repositories

import (
	"github.com/asliddinberdiev/task-manager/pkgs/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task models.Task) error {
	query := "INSERT INTO " + taskTable + " (title, description, priority, status, deadline) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Priority, task.Status, task.Deadline)

	return err
}

func (r *TaskRepository) GetById(taskId string) (models.Task, error) {
	var task models.Task
	query := "SELECT * FROM " + taskTable + " WHERE id = $1"
	err := r.db.Get(&task, query, taskId)

	return task, err
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	query := "SELECT * FROM " + taskTable
	err := r.db.Select(&tasks, query)

	return tasks, err
}

func (r *TaskRepository) Update(task models.Task) error {
	query := "UPDATE " + taskTable + " SET title=$1, description=$2, priority=$3, status=$4, deadline=$5 WHERE id=$7"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Priority, task.Status, task.Deadline, task.ID)
	return err
}

func (r *TaskRepository) Delete(taskId string) error {
	query := "DELETE FROM " + taskTable + " WHERE id = $1"
	_, err := r.db.Exec(query, taskId)

	return err
}
