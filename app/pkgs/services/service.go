package services

import "github.com/asliddinberdiev/task-manager/pkgs/repositories"

type Service struct {
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{}
}
