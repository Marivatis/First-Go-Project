package entity

import "sort"

const UnsetNoteId = -1

type Note struct {
	id    int
	Title string
	Body  string
}

func NewNote(id int, title, body string) Note {
	return Note{
		id:    id,
		Title: title,
		Body:  body,
	}
}

func (n *Note) Id() int {
	return n.id
}

func SortNotesById(notes []Note) {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Id() < notes[j].Id()
	})
}
