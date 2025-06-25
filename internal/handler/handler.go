package handler

import (
	"First-Go-Project/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/notes/:id", h.getNoteById)
	e.POST("/notes", h.createNote)
}
