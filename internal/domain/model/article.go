package model

type Article struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Content string `json:"content"`
}
