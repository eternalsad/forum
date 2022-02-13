package repository

import "database/sql"

type User interface {
	CreateUser() error
}

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{User: NewAuthRepository(db)}
}
