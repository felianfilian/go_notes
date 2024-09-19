package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	myNote, err := note.NewNote(title, content)
	myTodo, err := todo.NewTodo(todoText)
	fmt.Println(myTodo.Text)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	myNote.Display();

	// save note
	err = myNote.Save()
	if(err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println("Save Note Successful")

	// save todo
	err = myTodo.Save()
	if(err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println("Save Todo Successful")
}



func getNoteData() (string, string) {
	fmt.Println("Go Notes")
	title := getUserInput("note title: ")
	content := getUserInput("note content: ")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo: ")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var value string
	// fmt.Scan(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	
	if( err != nil) {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	data.Save()
	return nil
}