package service

import (
	"forum/intenal/repository"
)

type Service struct {
	User
}

type User interface {
	CreateUser() error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{User: NewAuthService(repo)}
}
