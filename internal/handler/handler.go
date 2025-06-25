package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/notes/:id", h.getNoteById)
	e.POST("/notes", h.createNote)
}
