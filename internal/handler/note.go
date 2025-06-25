package handler

import (
	"First-Go-Project/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getNoteById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get note by Id",
		"id":      id,
	})
}

func (h *Handler) createNote(c echo.Context) error {
	var input entity.Note
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Note created",
		"note":    input,
	})
}
