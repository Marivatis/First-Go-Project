package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

func (h *Handler) getNoteById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get note by Id",
		"id":      id,
	})
}

func (h *Handler) createNote(c echo.Context) error {
	type NoteInput struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	var input NoteInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Note created",
		"note":    input,
	})
}
