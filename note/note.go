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
	Title     string `json:"note_title"`
	Content   string `json:"note_content"`
	CreatedAt time.Time `json:"note_created"`
}

func (note Note) Display() {
	println("\n--The Note--")
	fmt.Printf("\nNote Title: %v\nNote Content: %v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

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
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}