package services

import "github.com/asliddinberdiev/task-manager/pkgs/repositories"

type UserService struct {
	repo repositories.User
}

func NewUserService(repo repositories.User) *UserService {
	return &UserService{repo: repo}
}
