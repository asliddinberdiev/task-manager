package models

import (
    "time"
    "github.com/go-playground/validator"
)

type Task struct {
    ID          string    `json:"id" db:"id"`
    UserID      int       `json:"user_id" db:"user_id"`
    CategoryID  string    `json:"category_id" db:"category_id"`
    Title       string    `json:"title" db:"title" validate:"required"`
    Description string    `json:"description" db:"description"`
    Priority    int       `json:"priority" db:"priority" validate:"oneof=1 2 3"`
    Status      int       `json:"status" db:"status" validate:"oneof=1 2 3"`
    Deadline    time.Time `json:"deadline" db:"deadline"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// ValidateTask validates the Task struct
func ValidateTask(t *Task) error {
    validate := validator.New()
    validate.RegisterValidation("future", func(fl validator.FieldLevel) bool {
        deadline := fl.Field().Interface().(time.Time)
        return deadline.After(time.Now())
    })
    
    validate.RegisterValidation("not_zero_time", func(fl validator.FieldLevel) bool {
        return !fl.Field().Interface().(time.Time).IsZero()
    })

    return validate.Struct(t)
}
