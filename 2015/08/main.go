package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string
var total int

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var newLine []string
		var skip int
		newLine = append(newLine, "\"")

		for i, character := range line {
			if 64 <= character && character <= 90 || 97 <= character && character <= 122 || skip > 0 {
				skip--
				newLine = append(newLine, string(character))
			} else if character == 92 { // 92 -- \
				if line[i+1] == 120 {
					skip = 3
				}
				characterNew := "\\" + string(character)
				newLine = append(newLine, characterNew)

			} else {
				characterNew := "\\" + string(character)
				newLine = append(newLine, characterNew)
			}
		}

		newLine = append(newLine, "\"")

		total += len(strings.Join(newLine, ""))
	}
	fmt.Println(total)
}
