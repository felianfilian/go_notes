package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Title     string `json:"note_title"`
	Text   string `json:"note_content"`
	CreatedAt time.Time `json:"note_created"`
}

func (todo Todo) Display() {
	println("\n--The Note--")
	fmt.Printf("\nNote Title: %v\nNote Content: %v\n\n", todo.Title, todo.Text)
}

func (todo Todo) Save() error {
	fileName := strings.ReplaceAll(todo.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)

	if(err != nil) {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func NewNote(title, content  string) (Todo, error) {
	if(title == "" || content == "") {
		return Todo{}, errors.New("invalid input")
	}

	return Todo {
		Title: title,
		Text: content,
		CreatedAt: time.Now(),
	}, nil
}