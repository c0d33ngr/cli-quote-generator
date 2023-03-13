package models

import "fmt"

type QuoteResponse struct {
	ID           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func (q QuoteResponse) TextOutput() string {
	qut := fmt.Sprintf("%s\n\n \t%s\n \t%s\n", q.Content, q.Author, q.Tags)
	return qut
}
