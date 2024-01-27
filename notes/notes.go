package notes

import (
	"fmt"
	"strconv"
	"strings"
)

func ShowNotes(collection string) {
	strings := Utilities.FileGetStrings(collection)
	for _, s := range strings {
		fmt.Println(s)
	}
	return
}

func AddNote(collection, text string) {
	file := Utilities.FileGetStrings(collection)

	lastFileNumber := 0
	if len(file) > 0 {
		lastNote := file[len(file)-1]
		lastFileNumber, _ = strconv.Atoi(strings.SplitN(lastNote, " - ", 2)[0])
	}

	nextFileNumber := lastFileNumber + 1

	newNote := fmt.Sprintf("%03d - %s", nextFileNumber, text)
	file = append(file, newNote)

	Utilities.FileCreate(collection)
	Utilities.FileAppendStrings(collection, file)
}

func DeleteNote(collection string, id int) {

}
