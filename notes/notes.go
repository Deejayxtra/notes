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

	Utilities.FileMakeEmpty(collection)
	Utilities.FileAppendStrings(collection, file)
}

func DeleteNote(collection string, id int) {
	strArr := Utilities.FileGetStrings(collection)
	if id < 0 || len(strArr) <= id {
		fmt.Println("Invalid Id")
		return
	}
	for i := id - 1; i < len(strArr)-1; i++ {
		s := strArr[i+1]
		before, after, found := strings.Cut(s, " - ")
		if found {
			num, _ := strconv.Atoi(before)
			num--
			s = fmt.Sprintf("%03d - %s", num, after)
		}
		strArr[i] = s
	}
	strArr = strArr[:len(strArr)-1]
	Utilities.FileMakeEmpty(collection)
	Utilities.FileAppendStrings(collection, strArr)
}