package dto

type NoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
