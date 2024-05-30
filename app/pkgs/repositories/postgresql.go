package repositories

import (
	"fmt"

	"github.com/asliddinberdiev/task-manager/pkgs/utils"
	"github.com/jmoiron/sqlx"
)

const (
	userTable     = "users"
	categoryTable = "categories"
	taskTable     = "tasks"
)

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Name,
		cfg.Password,
		cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	utils.Logger.Info("initialize postgres")

	return db, nil
}
