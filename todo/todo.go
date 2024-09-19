package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text   string `json:"note_content"`

}

func (todo Todo) Display() {
	println(todo.Text)

}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if(err != nil) {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func NewTodo(content string) (Todo, error) {
	if(content == "") {
		return Todo{}, errors.New("invalid input")
	}

	return Todo {
		Text: content,
	}, nil
}