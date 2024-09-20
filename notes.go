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

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	myNote, err := note.NewNote(title, content)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	myTodo, err := todo.NewTodo(todoText)
	fmt.Println(myTodo.Text)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	err = outputData(myNote)
	if(err != nil) {
		return
	}
	
	err = outputData(myTodo)
	if(err != nil) {
		return
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
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
	// save note
	err := data.Save()
	if(err != nil) {
		fmt.Println("saving note failed")
		return err
	}

	fmt.Println("Save Note Successful")	
	return nil
}