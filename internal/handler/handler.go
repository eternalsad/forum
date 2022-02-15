package handler

import (
	"forum/internal/service"
	"html/template"
)

// Handler struct to connect with service struct
type Handler struct {
	// service
	// TODO service that will validate
	// data given taken from handler
	service *service.Service
	t       *template.Template
}

// NewHandler creates new Handler entity
func NewHandler(service *service.Service, t *template.Template) *Handler {
	return &Handler{service, t}
}
