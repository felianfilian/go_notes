package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string `json:"note_title"`
	content   string `json:"note_content"`
	createdAt time.Time `json:"note_created"`
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