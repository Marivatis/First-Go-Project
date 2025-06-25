package service

import (
	"First-Go-Project/internal/entity"
	"First-Go-Project/internal/repository"
)

type NoteService struct {
	repo *repository.Repository
}

func NewNoteService(repo *repository.Repository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(note entity.Note) (int, error) {
	return s.repo.Create(note)
}

func (s *NoteService) GetById(id int) (entity.Note, error) {
	return s.repo.GetById(id)
}
