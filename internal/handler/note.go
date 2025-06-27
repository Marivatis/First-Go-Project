package handler

import (
	"First-Go-Project/internal/dto"
	"First-Go-Project/internal/mapper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func (h *Handler) createNote(c echo.Context) error {
	var input dto.NoteRequestCreate
	if err := c.Bind(&input); err != nil {
		log.Printf("bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	note := mapper.ToNoteEntityCreate(input)

	id, err := h.services.Create(note)
	if err != nil {
		log.Printf("note create error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("note created: %v", id)
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Note created",
		"id":      id,
	})
}

func (h *Handler) getNoteById(c echo.Context) error {
	id, err := parseIdParam(c)
	if err != nil {
		return err
	}

	note, err := h.services.GetById(id)
	if err != nil {
		log.Printf("note not found error: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	resp := mapper.ToNoteDTO(note)

	log.Printf("got note: %v", resp)
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) getAllNotes(c echo.Context) error {
	notes, err := h.services.GetAll()
	if err != nil {
		log.Printf("notes not found error: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	resp := mapper.ToNoteDTOList(notes)

	log.Printf("got notes: %v", resp)
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) updateNote(c echo.Context) error {
	var input dto.NoteRequestUpdate
	if err := c.Bind(&input); err != nil {
		log.Printf("bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	note, err := mapper.ToNoteEntityUpdate(input)
	if err != nil {
		log.Printf("invalid note format error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.services.Update(note)
	if err != nil {
		log.Printf("note update error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("note updated: %v", note)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Note updated",
		"note":    note,
	})
}

func (h *Handler) deleteNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("invalid id format error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	err = h.services.Delete(id)
	if err != nil {
		log.Printf("note delete error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("note deleted: %v", id)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Note deleted",
		"id":      id,
	})
}

func parseIdParam(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("invalid id format error: %v", err)
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	return id, nil
}
