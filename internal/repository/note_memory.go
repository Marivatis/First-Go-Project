package repository

import (
	"First-Go-Project/internal/entity"
	"errors"
	"sync"
)

type NoteMemory struct {
	mu     sync.RWMutex
	data   map[int]entity.Note
	lastId int
}

func NewNoteMemory() *NoteMemory {
	return &NoteMemory{
		data: make(map[int]entity.Note),
	}
}

func (r *NoteMemory) Create(note entity.Note) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastId++
	note.Id = r.lastId
	r.data[note.Id] = note
	return note.Id, nil
}

func (r *NoteMemory) GetById(id int) (entity.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, ok := r.data[id]
	if !ok {
		return entity.Note{}, errors.New("note not found")
	}
	return note, nil
}
