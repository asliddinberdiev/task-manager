package repositories

import (
	"github.com/asliddinberdiev/task-manager/pkgs/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(categoryName string) error {
	query := "INSERT INTO " + categoryTable + " (name) VALUES ($1)"
	_, err := r.db.Exec(query, categoryName)

	return err
}

func (r *CategoryRepository) GetById(categoryId string) (models.Category, error) {
	var user models.Category
	query := "SELECT * FROM " + categoryTable + " WHERE id = $1"
	err := r.db.Get(&user, query, categoryId)

	return user, err
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT * FROM " + categoryTable
	err := r.db.Select(&categories, query)

	return categories, err
}

func (r *CategoryRepository) Update(category models.Category) error {
	query := "UPDATE " + categoryTable + " SET name=$1 WHERE id=$2"
	_, err := r.db.Exec(query, category.Name, category.ID)
	return err
}

func (r *CategoryRepository) Delete(categoryId string) error {
	query := "DELETE FROM " + categoryTable + " WHERE id = $1"
	_, err := r.db.Exec(query, categoryId)

	return err
}
