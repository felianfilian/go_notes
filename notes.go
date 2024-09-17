package main

import (
	"fmt"
	"notes/note"
)

func main() {
	title, content := getNoteData()

	myNote, err := note.NewNote(title, content)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println(myNote)
	fmt.Println(content)
}

func getNoteData() (string, string) {
	fmt.Println("Go Notes")
	title := getUserInput("note title: ")
	content := getUserInput("note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	return value
}