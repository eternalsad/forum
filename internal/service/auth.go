package service

import (
	"fmt"
	"forum/internal/repository"
	"forum/models"
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

func (auth *AuthService) CreateUser(userData *models.User) error {
	exists, err := auth.repo.UserExists(userData)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("such email has already been taken")
	}
	err = auth.repo.CreateUser(userData)
	return err
}
