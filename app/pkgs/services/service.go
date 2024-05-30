package services

import "github.com/asliddinberdiev/task-manager/pkgs/repositories"

type User interface {
}

type Service struct {
	User
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
