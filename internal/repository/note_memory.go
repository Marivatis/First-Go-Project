package repository

import (
	"First-Go-Project/internal/entity"
	"errors"
	"sync"
)

var ErrNoteNotFound = errors.New("note not found")

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
	n := entity.NewNote(r.lastId, note.Title, note.Body)

	r.data[n.Id()] = n
	return n.Id(), nil
}

func (r *NoteMemory) GetById(id int) (entity.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, ok := r.data[id]
	if !ok {
		return entity.Note{}, ErrNoteNotFound
	}
	return note, nil
}

func (r *NoteMemory) GetAll() ([]entity.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	notes := make([]entity.Note, 0, len(r.data))
	for _, note := range r.data {
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *NoteMemory) Update(note entity.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[note.Id()]
	if !ok {
		return ErrNoteNotFound
	}
	r.data[note.Id()] = note
	return nil
}

func (r *NoteMemory) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[id]
	if !ok {
		return ErrNoteNotFound
	}
	delete(r.data, id)
	return nil
}
