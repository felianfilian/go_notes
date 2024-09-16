package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if(err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println(title)
	fmt.Println(content)
}

func getNoteData() (string, string, error) {
	fmt.Println("Go Notes")
	title, err := getUserInput("note title: ")

	if(err != nil){
		fmt.Println(err)
		return "","", err
	}

	content, err := getUserInput("note content: ")

	if(err != nil){
		fmt.Println(err)
		return "","", err
	}

	return title, content, err
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if(value == "") {
		return "", errors.New("invalid input")
	}
	return value, nil
}