package service

import (
	"First-Go-Project/internal/entity"
	"First-Go-Project/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateNoteCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")

		id, err := s.Create(note)

		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("fail, invalid note", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "", "test body")

		id, err := s.Create(note)

		assert.Error(t, err)
		assert.Equal(t, 0, id)
	})
}

func TestNoteService_GetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := s.Create(note)
		assert.NoError(t, err)

		got, err := s.GetById(id)

		assert.NoError(t, err)
		assert.Equal(t, "test title", got.Title)
		assert.Equal(t, "test body", got.Body)
	})
}

func TestNoteService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		for i := 0; i < 5; i++ {
			note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
			_, err := s.Create(note)
			assert.NoError(t, err)
		}

		got, err := s.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, 5, len(got))

		for i, note := range got {
			assert.Equal(t, i+1, note.Id())
		}
	})

	t.Run("success, empty list", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)

		got, err := s.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(got))
	})
}

func TestNoteService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := s.Create(note)
		assert.NoError(t, err)

		updateNote := entity.NewNote(id, "updated title", "updated body")
		err = s.Update(updateNote)

		assert.NoError(t, err)

		got, err := s.GetById(id)
		assert.NoError(t, err)
		assert.Equal(t, "updated title", got.Title)
		assert.Equal(t, "updated body", got.Body)
	})

	t.Run("fail, invalid note id", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		_, err := s.Create(note)
		assert.NoError(t, err)

		updateNote := entity.NewNote(0, "updated title", "updated body")
		err = s.Update(updateNote)

		assert.Error(t, err)

		_, err = s.GetById(0)
		assert.Error(t, err)
	})

	t.Run("fail, invalid note title", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := s.Create(note)
		assert.NoError(t, err)

		updateNote := entity.NewNote(id, "", "updated body")
		err = s.Update(updateNote)

		assert.Error(t, err)

		got, err := s.GetById(id)
		assert.NoError(t, err)
		assert.Equal(t, "test title", got.Title)
		assert.Equal(t, "test body", got.Body)
	})
}

func TestNoteService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := repository.NewRepository()
		s := NewNoteService(repo)
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := s.Create(note)
		assert.NoError(t, err)

		err = s.Delete(id)

		assert.NoError(t, err)

		_, err = s.GetById(id)
		assert.ErrorIs(t, err, repository.ErrNoteNotFound)
	})
}
