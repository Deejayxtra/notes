package main

import (
	"fmt"
	"notes/notes"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./notestool <collection>")
		os.Exit(1)
	}

	collection := os.Args[1]
	notes.UserInterface(collection)
}
