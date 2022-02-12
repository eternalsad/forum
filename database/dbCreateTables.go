package database

import (
	"database/sql"
	"io/ioutil"
)

// CreateTables creates tables inside db
func CreateTables(db *sql.DB) error {
	query, err := ioutil.ReadFile("./database/schema/up/users.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}
	return nil
}
