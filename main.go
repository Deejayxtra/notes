package main

import (
	"fmt"
	"notes/notes"
	"os"
)

func main() {
	fmt.Println("Welcome to the notes tool!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./notestool <collection>")
		os.Exit(1)
	}
	collection := os.Args[1]
	_, err := os.Stat(collection)
	if os.IsNotExist(err) {
		notes.Utilities.FileCreate(collection)
	}
	notes.UserInterface(collection)
}
