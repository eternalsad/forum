package handler

import (
	"fmt"
	"net/http"
)

func (handler *Handler) RenderMain() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := handler.t.ExecuteTemplate(rw, "main", "aboba")
		if err != nil {
			fmt.Println(err)
		}
	}
}
