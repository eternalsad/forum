package repository

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

func (auth *AuthRepository) CreateUser() error {
	return nil
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}
