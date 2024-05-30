package models

import (
    "time"
    "github.com/go-playground/validator"
)

type User struct {
    ID          int        `json:"id" db:"id"`
    Username    string     `json:"username" db:"username"`
    FirstName   string     `json:"first_name" validate:"required" db:"first_name"`
    LastName    string     `json:"last_name" db:"last_name"`
    Admin       bool       `json:"admin" db:"admin"`
    Phone       string     `json:"phone" validate:"len=12" db:"phone"`
    Action      string     `json:"action" db:"action"`
    CreatedAt   time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}


// ValidateUser validates the User struct
func ValidateUser(u *User) error {
    validate := validator.New()
    return validate.Struct(u)
}
