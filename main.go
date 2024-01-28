package main

import (
	"fmt"
	"notes/notes"

)

func main() {
	fmt.Println(notes.ShowNotes("file"))

	collection := "file" 
	var fileText string
	notes.AddNote(collection, fileText)

}

