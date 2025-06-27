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
	h.registerNoteRoutes(e)
}

func (h *Handler) registerNoteRoutes(e *echo.Echo) {
	notes := e.Group("/notes")
	notes.POST("", h.createNote)
	notes.GET("", h.getAllNotes)
	notes.GET("/:id", h.getNoteById)
	notes.PUT("/:id", h.updateNote)
	notes.DELETE("/:id", h.deleteNote)
}
