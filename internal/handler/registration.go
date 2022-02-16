package handler

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"
)

// RenderRegister handles register endpoint
func (handler *Handler) RenderRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler.t.Execute(w, "aboba")
		err := handler.t.ExecuteTemplate(w, "main", "aboba")
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Registration goes brrr...
func (handler *Handler) Registration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(r.Form)
		userData, err := models.NewUser(r.Form)
		if err != nil {
			log.Println(err)
			//should redirect user to a error page
			return
		}
		err = handler.service.CreateUser(userData)
		if err != nil {
			// should redirect user to a error page
			// handler.t.ExecuteTemplate(w, "authentication", "sorry email has already been taken")
			handler.ErrorHandler(http.StatusForbidden, config.ErrorEmailTaken, "authentication", w)
			log.Println(err)
			return
		}
		// otherwise what should be done???
		// should redirect user to handler (handler *Handler)succesfull()
		handler.t.ExecuteTemplate(w, "authentication", "success")
	}
}
