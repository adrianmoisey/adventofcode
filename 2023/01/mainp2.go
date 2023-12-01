package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lookup := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var total int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var firstInt string
		var lastInt string

		ogLine := line
		startLine := line
		endLine := line

		for {
			var lookupIndex string
			var matchIndex int
			var allError []int
			matchIndex = 999999999

			for k := range lookup {
				index := strings.Index(startLine, k)

				if index < matchIndex && index != -1 {
					matchIndex = index
					lookupIndex = k
				}
				if index == -1 {
					allError = append(allError, -1)
				}
			}

			startLine = strings.ReplaceAll(startLine, lookupIndex, lookup[lookupIndex])

			break
		}

		for _, character := range startLine {

			if _, err := strconv.Atoi(string(character)); err == nil {
				if firstInt == "" {
					firstInt = string(character)
				}
			}
		}

		for {
			var lookupIndex string
			var matchIndex int
			var allError []int
			matchIndex = -1

			for k := range lookup {
				index := strings.LastIndex(endLine, k)

				if index > matchIndex && index != -1 {
					matchIndex = index
					lookupIndex = k
				}
				if index == -1 {
					allError = append(allError, -1)
				}
			}
			endLine = strings.ReplaceAll(endLine, lookupIndex, lookup[lookupIndex])

			break
		}

		for _, character := range endLine {

			if _, err := strconv.Atoi(string(character)); err == nil {
				lastInt = string(character)
			}
		}
		value := firstInt + lastInt
		fmt.Println(value, ogLine, startLine, endLine)
		intValue, _ := strconv.Atoi(value)

		total += intValue
	}
	fmt.Println("Part2: ")
	fmt.Println(total)
}
