package dto

type NoteRequestCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteRequestUpdate struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
