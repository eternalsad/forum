package service

import (
	"forum/internal/repository"
	"forum/models"
)

type Service struct {
	User
}

type User interface {
	CreateUser(*models.User) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{User: NewAuthService(repo)}
}
