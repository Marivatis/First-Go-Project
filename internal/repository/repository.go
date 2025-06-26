package repository

import "First-Go-Project/internal/entity"

type NoteItem interface {
	Create(note entity.Note) (int, error)
	GetById(id int) (entity.Note, error)
	GetAll() ([]entity.Note, error)
	Update(note entity.Note) error
	Delete(id int) error
}

type Repository struct {
	NoteItem
}

func NewRepository() *Repository {
	return &Repository{
		NoteItem: NewNoteMemory(),
	}
}
