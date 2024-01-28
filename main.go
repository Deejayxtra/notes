package main

import (
	"fmt"
	"notes/notes"
	"os"
)

func main() {
	notes.UserInterface("foobar")
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./notestool <collection>")
		os.Exit(1)
	}

	collection := os.Args[1]
	notes.UserInterface(collection)
	displayMenu(collection)
}

func displayMenu(collection string) {
	var choice int

	for {
		fmt.Println("1. Show Notes")
		fmt.Println("2. Add Note")
		fmt.Println("3. Delete Note")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice (1-4): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			notes.ShowNotes(collection)
		case 2:
			fmt.Print("Enter the note text: ")
			var noteText string
			fmt.Scanln(&noteText)
			notes.AddNote(collection, noteText)
		case 3:
			fmt.Print("Enter the note ID to delete: ")
			var noteID int
			fmt.Scanln(&noteID)
			notes.DeleteNote(collection, noteID)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
