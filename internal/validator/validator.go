package validator

import (
	"First-Go-Project/internal/entity"
	"errors"
	"strings"
)

func ValidateNoteCreate(note entity.Note) error {
	if strings.TrimSpace(note.Title) == "" {
		return errors.New("title is required")
	}
	if len(note.Title) > 100 {
		return errors.New("title too long")
	}
	return nil
}

func ValidateNoteUpdate(note entity.Note) error {
	if note.Id() < 0 {
		return errors.New("invalid note id")
	}
	if strings.TrimSpace(note.Title) == "" {
		return errors.New("title is required")
	}
	if len(note.Title) > 100 {
		return errors.New("title too long")
	}
	return nil
}
