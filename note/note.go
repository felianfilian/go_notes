package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string `json:"note_title"`
	content   string `json:"note_content"`
	createdAt time.Time `json:"note_created"`
}

func (note Note) Display() {
	println("\nThe Note")
	fmt.Printf("\nNote Title: %v\nNote Content: %v", note.title, note.content)
}

func NewNote(title, content  string) (Note, error) {
	if(title == "" || content == "") {
		return Note{}, errors.New("invalid input")
	}

	return Note {
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}