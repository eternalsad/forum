package handler

import (
	"forum/intenal/service"
)

// Handler struct to connect with service struct
type Handler struct {
	// service
	// TODO service that will validate
	// data given taken from handler
	service *service.Service
}
