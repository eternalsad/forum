package main

import (
	"fmt"
	"forum/database"
	"forum/internal/handler"
	"forum/internal/repository"
	"forum/internal/service"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.DBInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	files := []string{
		"./templates/html/main.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("cannot parse templates")
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, ts)

	router := http.NewServeMux()
	router.HandleFunc("/register", handler.RenderRegister())
	router.HandleFunc("/registration", handler.Registration())
	log.Print("starting server on: 8080")
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
