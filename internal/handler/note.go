package handler

import (
	"First-Go-Project/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	note, err := h.services.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, note)
}

func (h *Handler) createNote(c echo.Context) error {
	var input entity.Note
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, err := h.services.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Note created",
		"id":      id,
	})
}
