package repository

import (
	"database/sql"
	"forum/models"
)

type User interface {
	CreateUser(*models.User) error
	UserExists(*models.User) (bool, error)
}

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{User: NewAuthRepository(db)}
}
