package models

import (
	"time"

	"github.com/go-playground/validator"
)

type Category struct {
	ID        string    `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name" validate:"required,min=3,max=25"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ValidateCategory validates the Category struct
func ValidateCategory(c *Category) error {
	validate := validator.New()
	return validate.Struct(c)
}
