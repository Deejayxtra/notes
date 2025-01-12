package notes

import (
	"bufio"
	"fmt"
	"os"
)

func displayMenu() int {
	fmt.Println("\vSelect operation:")
	fmt.Println("1. Show notes.")
	fmt.Println("2. Add a note.")
	fmt.Println("3. Delete a note.")
	fmt.Println("4. Exit")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0 // Palautetaan 0, jotta pääohjelman silmukka pysähtyy ja käyttäjä näkee virheilmoituksen
	}
	return choice
}

func showNotes(collection string) {
	fmt.Println("\vNotes:")
	ShowNotes(collection)
	// Kutsutaan kokoelman käsittelijän funktiota muistiinpanojen hakemiseksi
	// ja tulostetaan muistiinpanot näytölle
}

func addNote(collection string) {
	fmt.Println("\vEnter the note text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		noteText := scanner.Text()
		// Kutsutaan kokoelman käsittelijän funktiota uuden muistiinpanon lisäämiseksi
		// Huomaa, että tämä on vain esimerkki, oikea kutsu riippuu kokoelman toteutuksesta
		//		addNoteToCollection(noteText)
		AddNote(collection, noteText)
		fmt.Println("Note added successfully.")
	}
}

func deleteNote(collection string) {
	fmt.Println("\vEnter the number of note to remove or 0 to cancel:")
	var noteID int
	fmt.Scanln(&noteID)
	// Kutsutaan kokoelman käsittelijän funktiota muistiinpanon poistamiseksi
	// Huomaa, että tämä on vain esimerkki, oikea kutsu riippuu kokoelman toteutuksesta
	//	deleteNoteFromCollection(noteID)
	DeleteNote(collection, noteID)
}

func UserInterface(collection string) {
	for {
		choice := displayMenu()

		switch choice {
		case 1:
			showNotes(collection)
		case 2:
			addNote(collection)
		case 3:
			deleteNote(collection)
		case 4:
			fmt.Println("\vExiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
