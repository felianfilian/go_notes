package note

import "time"

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func NewNote(title, content  string) Note {
	return Note {
		title: title,
		content: content,
		createdAt: time.Now(),
	}
}