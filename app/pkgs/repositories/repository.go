package repositories

import (
	"github.com/asliddinberdiev/task-manager/pkgs/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) error
	GetById(userId string) (models.User, error)
	GetAll() ([]models.User, error)
	Update(user models.User) error
	Delete(userId string) error
}

type Category interface {
	Create(categoryName string) error
	GetById(categoryId string) (models.Category, error)
	GetAll() ([]models.Category, error)
	Update(category models.Category) error
	Delete(categoryId string) error
}

type Task interface {
	Create(task models.Task) error
	GetById(taskId string) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(task models.Task) error
	Delete(taskId string) error
}

type Repository struct {
	User
	Category
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Category: NewCategoryRepository(db),
		Task:     NewTaskRepository(db),
	}
}
