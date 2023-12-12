package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

func main() {

	linesInput := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range linesInput {
		splitString := strings.Split(line, " ")
		SpringCondition := splitString[0]
		Mapping := splitString[1]
		fmt.Println(SpringCondition, Mapping)
	}
}

func BuildCombos(length int, combo string)
