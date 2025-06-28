package mapper

import (
	"First-Go-Project/internal/dto"
	"First-Go-Project/internal/entity"
)

func ToNoteEntity(input dto.NoteRequest) entity.Note {
	return entity.NewNote(entity.UnsetNoteId, input.Title, input.Body)
}

func ToNoteDTO(note entity.Note) dto.NoteResponse {
	return dto.NoteResponse{
		Id:    note.Id(),
		Title: note.Title,
		Body:  note.Body,
	}
}

func ToNoteDTOList(notes []entity.Note) []dto.NoteResponse {
	resp := make([]dto.NoteResponse, 0, len(notes))
	for _, note := range notes {
		resp = append(resp, ToNoteDTO(note))
	}
	return resp
}
