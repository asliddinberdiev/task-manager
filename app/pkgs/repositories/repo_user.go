package repositories

import (
	"github.com/asliddinberdiev/task-manager/pkgs/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) error {
	query := "INSERT INTO " + userTable + " (id, username, first_name, last_name, admin, phone, action) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, user.ID, user.Username, user.FirstName, user.LastName, user.Admin, user.Phone, user.Action)

	return err
}

func (r *UserRepository) GetById(userId string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM " + userTable + " WHERE id = $1"
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM " + userTable
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserRepository) Update(user models.User) error {
	query := "UPDATE " + userTable + " SET username=$1, first_name=$2, last_name=$3, admin=$4, phone=$5, action=$6 WHERE id=$7"
	_, err := r.db.Exec(query, user.Username, user.FirstName, user.LastName, user.Admin, user.Phone, user.Action, user.ID)
	return err
}

func (r *UserRepository) Delete(userId string) error {
	query := "DELETE FROM " + userTable + " WHERE id = $1"
	_, err := r.db.Exec(query, userId)

	return err
}
