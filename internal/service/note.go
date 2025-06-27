package service

import (
	"First-Go-Project/internal/entity"
	"First-Go-Project/internal/repository"
	"First-Go-Project/internal/validator"
)

type NoteService struct {
	repo *repository.Repository
}

func NewNoteService(repo *repository.Repository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(note entity.Note) (int, error) {
	err := validator.ValidateNoteCreate(note)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(note)
}

func (s *NoteService) GetById(id int) (entity.Note, error) {
	return s.repo.GetById(id)
}

func (s *NoteService) GetAll() ([]entity.Note, error) {
	notes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	entity.SortNotesById(notes)
	return notes, nil
}

func (s *NoteService) Update(note entity.Note) error {
	err := validator.ValidateNoteUpdate(note)
	if err != nil {
		return err
	}
	return s.repo.Update(note)
}

func (s *NoteService) Delete(id int) error {
	return s.repo.Delete(id)
}
