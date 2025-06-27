package service

import (
	"First-Go-Project/internal/entity"
	"First-Go-Project/internal/repository"
)

type NoteItem interface {
	Create(note entity.Note) (int, error)
	GetById(id int) (entity.Note, error)
	GetAll() ([]entity.Note, error)
	Update(note entity.Note) error
	Delete(id int) error
}

type Service struct {
	NoteItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NoteItem: NewNoteService(repo),
	}
}
