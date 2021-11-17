package dto

type CreateNotesRequest struct {
	Title string `json:"title"`
	Note string `json:"note"`
}
