package main

import (
	"fmt"
	"forum/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := database.DBInit()
	if err != nil {
		fmt.Println(err)
		return
	}

}
