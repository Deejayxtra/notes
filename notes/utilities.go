package notes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type implStrings struct{}

var Utilities implStrings

func (implStrings) CmdGetString() string {
	var input string
	fmt.Scanln(&input)
	input = strings.Trim(input, " ")
	return input
}

// returns -1 if userinput is not int
func (implStrings) CmdGetInt() int {
	var input string
	fmt.Scanln(&input)
	input = strings.Trim(input, " ")
	userinput, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return -1
	}
	return userinput
}

func (implStrings) FileGetStrings(filename string) []string {
	result := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return []string{""}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func (implStrings) FileMakeEmpty(filename string) {
	err := os.Truncate(filename, 0)
	if err != nil {
		log.Println(err)
	}
}

// won't do automatic \n so can be added before/after as needed
func (implStrings) FileAppendString(filename, text string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		log.Println(err)
	}
}

// won't do automatic \n so can be added before/after as needed
// will still do \n between the strings
func (implStrings) FileAppendStrings(filename string, strings []string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	for i, text := range strings {
		if i > 0 {
			text = "\n" + text
		}
		_, err = file.WriteString(text)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
