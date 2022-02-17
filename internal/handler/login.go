package handler

import "net/http"

func (handler *Handler) RenderLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		handler.t.ExecuteTemplate(w, "authentication", "")
	}
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {

}
