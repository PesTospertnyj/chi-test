package dto

type Book struct {
	ID     int64  `json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
