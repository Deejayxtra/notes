package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowNotes(collection string) []string {
	return Utilities.FileGetStrings(collection)

}

func AddNote(collection, text string) {
	file := Utilities.FileGetStrings(collection)

	lastFileNumber := 0
	if len(file) > 0 {
		lastNote := file[len(file)-1]
		lastFileNumber, _ = strconv.Atoi(strings.SplitN(lastNote, " - ", 2)[0])
	}

	nextFileNumber := lastFileNumber + 1
	fmt.Print("Enter the note text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileText := scanner.Text()

	newNote := fmt.Sprintf("%03d - %s", nextFileNumber, fileText)
	file = append(file, newNote)

	Utilities.FileCreate(collection)
	Utilities.FileAppendStrings(collection, file)
}

func DeleteNote(collection string, id int) {
	
}
