package service

import (
	"forum/intenal/repository"
)

type AuthService struct {
	// repo connection to repository
	// repo repository.User
	repo repository.User
}

// what the fuck should i recieve here???
func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (auth *AuthService) CreateUser() error {
	return nil
}
