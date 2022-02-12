package database

import "database/sql"

// DBIsEmpty checks whether users table
// is empty or not
func DBIsEmpty(db *sql.DB) bool {
	_, ok := db.Query("SELECT * FROM users;")
	if ok != nil {
		return false
	}
	return true
}
