package database

import (
	"database/sql"
)

// DBInit Initializes pointer to a database
// If database doesn't exist it creates
// database
func DBInit() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if !DBIsEmpty(db) {
		if err := CreateTables(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
