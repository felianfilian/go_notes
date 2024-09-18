package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	myNote, err := note.NewNote(title, content)

	if(err != nil) {
		fmt.Println(err)
		return
	}

	myNote.Display();

	//fmt.Println(myNote)
	//fmt.Println(content)
}

func getNoteData() (string, string) {
	fmt.Println("Go Notes")
	title := getUserInput("note title: ")
	content := getUserInput("note content: ")

	return title, content
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