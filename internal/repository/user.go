package repository

import (
	"database/sql"
	"fmt"
	"forum/models"
	"log"
)

type AuthRepository struct {
	db *sql.DB
}

func (auth *AuthRepository) CreateUser(userData *models.User) error {
	stmt, err := auth.db.Prepare("INSERT INTO users (username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf(err.Error() + "statement error")
	}
	_, err = stmt.Exec(userData.Username, userData.Email, userData.Password)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("db: ", err.Error(), " username or email has already been taken")
	}
	return err
}

func (auth *AuthRepository) UserExists(userData *models.User) (bool, error) {
	// stmt := fmt.Sprintf("SELECT email from users WHERE email='%v';", userData.Email)
	// res, err := auth.db.Exec(stmt)
	// if err != nil {
	// 	return false, err
	// }
	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	return false, err
	// }

	// if rows != 0 {
	// 	return true, nil
	// }
	// return false, nil
	var emailFromDB string
	rowNew := auth.db.QueryRow(`SELECT email FROM users WHERE email = ?;`, userData.Email)
	err := rowNew.Scan(&emailFromDB)
	fmt.Printf("emailFromDB: %v\n", emailFromDB)
	if err != nil {
		fmt.Println("huyhurma", err.Error())
		return false, err
	}
	// return true, nil
	if &emailFromDB == nil {
		return false, nil
	}
	return true, nil
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}
