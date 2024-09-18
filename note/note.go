package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
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

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)

	if(err != nil) {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
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