package models

type Blog struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BlogDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
