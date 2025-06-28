package repository

import (
	"First-Go-Project/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoteMemory_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := NewNoteMemory()
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := repo.Create(note)
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("multiple creates", func(t *testing.T) {
		repo := NewNoteMemory()
		for i := 0; i < 5; i++ {
			note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
			id, err := repo.Create(note)
			assert.NoError(t, err)
			assert.Equal(t, i+1, id)
		}
	})
}

func TestNoteMemory_GetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := NewNoteMemory()
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, err := repo.Create(note)
		assert.NoError(t, err)

		got, err := repo.GetById(id)
		assert.NoError(t, err)
		assert.Equal(t, "test title", got.Title)
		assert.Equal(t, "test body", got.Body)
		assert.Equal(t, id, got.Id())
	})

	t.Run("not found", func(t *testing.T) {
		repo := NewNoteMemory()
		_, err := repo.GetById(999)
		assert.Error(t, err)
	})
}

func TestNoteMemory_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := NewNoteMemory()
		for i := 0; i < 5; i++ {
			note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
			_, err := repo.Create(note)
			assert.NoError(t, err)
		}

		got, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, got, 5)

		ids := make(map[int]bool)
		for _, note := range got {
			id := note.Id()
			assert.Equal(t, "test title", note.Title)
			assert.Equal(t, "test body", note.Body)
			assert.False(t, ids[id], "duplicate id found")
			ids[id] = true
		}
	})
}

func TestNoteMemory_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := NewNoteMemory()
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, _ := repo.Create(note)

		updateNote := entity.NewNote(id, "updated title", "updated body")
		err := repo.Update(updateNote)
		assert.NoError(t, err)

		got, err := repo.GetById(id)
		assert.NoError(t, err)
		assert.Equal(t, "updated title", got.Title)
		assert.Equal(t, "updated body", got.Body)
	})

	t.Run("not found", func(t *testing.T) {
		repo := NewNoteMemory()
		note := entity.NewNote(42, "some title", "some body")

		err := repo.Update(note)
		assert.ErrorIs(t, err, ErrNoteNotFound)
	})
}

func TestNoteMemory_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := NewNoteMemory()
		note := entity.NewNote(entity.UnsetNoteId, "test title", "test body")
		id, _ := repo.Create(note)

		err := repo.Delete(id)
		assert.NoError(t, err)

		_, err = repo.GetById(id)
		assert.ErrorIs(t, err, ErrNoteNotFound)
	})

	t.Run("not found", func(t *testing.T) {
		repo := NewNoteMemory()

		err := repo.Delete(999)
		assert.ErrorIs(t, err, ErrNoteNotFound)
	})
}
