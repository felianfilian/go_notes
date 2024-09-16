package main

import (
	"fmt"
)

func main() {
	title, content := getNoteData()

	

	fmt.Println(title)
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