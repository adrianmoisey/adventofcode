package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	Grid       = make(map[int]string)
	total      int
	GridWidth  int
	GridLength int
)

func GridLookerUpper(indexStart, indexEnd, row int) (symbol bool) {

	symbol = false

	var characters []string
	fmt.Println("Block start")

	for c := indexStart; c <= indexEnd; c++ {
		fmt.Println(string(Grid[row][c]))
	}

	left := indexStart - 1
	right := indexEnd + 1
	up := row - 1
	middle := row
	down := row + 1

	for pos := left; pos <= right; pos++ {
		if pos < 0 || up < 0 {
			continue
		}
		characters = append(characters, string(Grid[up][pos]))
	}

	for pos := left; pos <= right; pos++ {
		if pos < 0 || down > GridLength-1 {
			continue
		}
		characters = append(characters, string(Grid[down][pos]))
	}

	if left >= 0 {
		characters = append(characters, string(Grid[middle][left]))
	}
	if right <= GridWidth {
		characters = append(characters, string(Grid[middle][right]))
	}

	for _, character := range characters {
		fmt.Println(character)
		if _, err := strconv.Atoi(character); err != nil {
			if character != "." {
				symbol = true
			}
		}

	}

	return symbol
}

func main() {
	lines := strings.Split(input, "\n")
	for index, line := range lines {
		if line == "" {
			break
		}
		Grid[index] = line
	}

	GridLength = len(Grid)
	GridWidth = len(Grid[0])

	// Loop over grid and find numbers
	for row := 0; row < GridLength; row++ {

		CurrentNumber := ""
		IndexStart := -1

		for characterIndex := 0; characterIndex < GridWidth; characterIndex++ {
			ChatacterString := string(Grid[row][characterIndex])

			_, err := strconv.Atoi(ChatacterString)
			if err == nil {
				CurrentNumber += ChatacterString
				if IndexStart == -1 {
					IndexStart = characterIndex
				}

				if characterIndex+1 <= GridWidth {
					// This is a bug. I made the grid 1 wider to account for it
					ChatacterStringNext := string(Grid[row][characterIndex+1])
					_, err := strconv.Atoi(ChatacterStringNext)
					if err != nil {
						symbol := GridLookerUpper(IndexStart, characterIndex, row)
						fmt.Println("Symbol? ", symbol, CurrentNumber)
						if symbol {
							value, _ := strconv.Atoi(CurrentNumber)
							total += value
						}

					}
				}
			} else {
				CurrentNumber = ""
				IndexStart = -1
			}
		}
	}
	fmt.Println(total)
	// 411409 too low
	// 526341 too low
	// 528819 - right - hacked by making grid 1 wider
	fmt.Println("end")
}
