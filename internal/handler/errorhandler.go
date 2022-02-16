package handler

import "net/http"

func (handler *Handler) ErrorHandler(code int, message string, templateName string, w http.ResponseWriter) {
	w.WriteHeader(code)
	handler.t.ExecuteTemplate(w, templateName, message)
}
