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
		"./templates/html/registration.html",
		"./templates/html/main.html",
	}
	// ts := template.Must(template.ParseGlob("./templates/html"))
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
	router.HandleFunc("/", handler.RenderMain())
	log.Print("starting server on: 8989")
	err = http.ListenAndServe("localhost:8989", router)
	if err != nil {
		log.Fatal(err)
	}
}
