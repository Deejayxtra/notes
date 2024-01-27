package notes

import (
	"bufio"
	"fmt"
	"os"
)

func displayMenu() int {
	fmt.Println("Select operation:")
	fmt.Println("1. Show notes.")
	fmt.Println("2. Add a note.")
	fmt.Println("3. Delete a note.")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scanln(&choice)
	return choice
}

func showNotes() {
	fmt.Println("Notes:")
	// Kutsutaan kokoelman käsittelijän funktiota muistiinpanojen hakemiseksi
	// ja tulostetaan muistiinpanot näytölle
}

func addNote() {
	fmt.Println("Enter note:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		noteText := scanner.Text()
		// Kutsutaan kokoelman käsittelijän funktiota uuden muistiinpanon lisäämiseksi
		// Huomaa, että tämä on vain esimerkki, oikea kutsu riippuu kokoelman toteutuksesta
		//		addNoteToCollection(noteText)
		fmt.Println("Note added successfully.")
		fmt.Println(noteText)
	}
}

func deleteNote() {
	fmt.Println("Enter note ID to delete:")
	var noteID int
	fmt.Scanln(&noteID)
	// Kutsutaan kokoelman käsittelijän funktiota muistiinpanon poistamiseksi
	// Huomaa, että tämä on vain esimerkki, oikea kutsu riippuu kokoelman toteutuksesta
	//	deleteNoteFromCollection(noteID)
	fmt.Println("Note deleted successfully.")
}

func UserInterface(collection string) {
	for {
		choice := displayMenu()

		switch choice {
		case 1:
			showNotes()
		case 2:
			addNote()
		case 3:
			deleteNote()
		case 4:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
