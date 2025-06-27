package handler

import (
	"First-Go-Project/internal/dto"
	"First-Go-Project/internal/mapper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) createNote(c echo.Context) error {
	var input dto.NoteRequestCreate
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	note := mapper.ToNoteEntityCreate(input)

	id, err := h.services.Create(note)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Note created",
		"id":      id,
	})
}

func (h *Handler) getNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	note, err := h.services.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	resp := mapper.ToNoteDTO(note)

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) getAllNotes(c echo.Context) error {
	notes, err := h.services.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, notes)
}

func (h *Handler) updateNote(c echo.Context) error {
	var input dto.NoteRequestUpdate
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	note, err := mapper.ToNoteEntityUpdate(input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.services.Update(note)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Note updated",
		"note":    note,
	})
}

func (h *Handler) deleteNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	err = h.services.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Note deleted",
		"id":      id,
	})
}
