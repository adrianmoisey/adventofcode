package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var total int
var totalp1 int

func main() {
	Hexacecimal := "0123456789abcdef"
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		// Part 1
		strippedLine := line[1 : len(line)-1]

		for strings.Index(strippedLine, "\\") != -1 {
			Index := strings.Index(strippedLine, "\\")
			if Index+1 < len(strippedLine) && string(strippedLine[Index+1]) == "x" {
				strippedLine = strippedLine[:Index] + strippedLine[Index+3:]

			} else {
				strippedLine = strippedLine[:Index] + "_" + strippedLine[Index+2:]
			}
		}
		totalp1 += len(line) - len(strippedLine)

		// Part 2
		var newLine []string
		var skip int
		newLine = append(newLine, "\"")

		for i, character := range line {
			if 64 <= character && character <= 90 || 97 <= character && character <= 122 || skip > 0 {
				skip--
				newLine = append(newLine, string(character))
			} else if character == 92 { // 92 is a \
				if line[i+1] == 120 && strings.Contains(Hexacecimal, string(line[i+2])) && strings.Contains(Hexacecimal, string(line[i+3])) { // 120 is a lowercase x
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
		fmt.Println()
		fmt.Println(line)
		fmt.Println(line, len(strippedLine))
		fmt.Println(strings.Join(newLine, ""), len(strings.Join(newLine, "")))

		total += len(strings.Join(newLine, "")) - len(line)
	}
	// 2116 too low
	// 3487 too high
	fmt.Println(totalp1)
	fmt.Println(total)
}
