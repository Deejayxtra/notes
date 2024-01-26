package notes

import (
	"fmt"
	"strconv"
	"strings"
)

type implStrings struct{}

var Utilities implStrings

func (implStrings) GetStringFromCmd() string {
	var input string
	fmt.Scanln(&input)
	input = strings.Trim(input, " ")
	return input
}

func (implStrings) GetIntFromCmd() int {
	var input string
	fmt.Scanln(&input)
	input = strings.Trim(input, " ")
	userinput, err := strconv.Atoi(input)
	if err == nil {
		return -1
	}
	return userinput
}
