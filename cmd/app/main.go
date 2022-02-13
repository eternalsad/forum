package main

import (
	"fmt"
	"forum/database"
	"forum/intenal/handler"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := database.DBInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	router := http.NewServeMux()
	router.HandleFunc("/register", handler.Register)
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
