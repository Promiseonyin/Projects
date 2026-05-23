package models

import (
	"fmt"
	"strings"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNoteInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (i CreateNoteInput) Validate() error {
	if strings.TrimSpace(i.Title) == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}
