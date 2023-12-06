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
	Gears      = make(map[int]map[int]string)
	total      int
	GridWidth  int
	GridLength int
)

func GridLookerUpper(indexStart, indexEnd, row int, CurrentNumber string) (symbol bool) {

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
		if string(Grid[up][pos]) == "*" {
			fmt.Println("HERE")
			fmt.Println(Gears[up])
			fmt.Println(Gears[1])
			val := Gears[up][pos]

			Gears[up][pos] = val + "," + CurrentNumber
		}
		characters = append(characters, string(Grid[up][pos]))
	}

	for pos := left; pos <= right; pos++ {
		if pos < 0 || down > GridLength-1 {
			continue
		}
		if string(Grid[down][pos]) == "*" {
			Gears[down][pos] = Gears[down][pos] + "," + CurrentNumber
		}

		characters = append(characters, string(Grid[down][pos]))
	}
	if left >= 0 {
		if string(Grid[middle][left]) == "*" {
			Gears[middle][left] = Gears[middle][left] + "," + CurrentNumber
		}

		characters = append(characters, string(Grid[middle][left]))
	}
	if right <= GridWidth {
		if string(Grid[middle][right]) == "*" {
			Gears[middle][right] = Gears[middle][right] + "," + CurrentNumber
		}

		characters = append(characters, string(Grid[middle][right]))
	}
	fmt.Println(Gears[1])

	return symbol
}

func main() {

	//Gears := map[int]map[int]string{}
	lines := strings.Split(input, "\n")
	for index, line := range lines {
		if line == "" {
			break
		}
		Grid[index] = line
	}

	GridLength = len(Grid)
	GridWidth = len(Grid[0])
	// Popular Gears
	for row := 0; row < GridLength; row++ {
		Gears[row] = map[int]string{}
		for width := 0; width < GridWidth; width++ {
			Gears[row][width] = ""
		}
	}

	// Loop over grid and find numbers
	for row := 0; row < GridLength; row++ {
		CurrentNumber := ""
		IndexStart := -1

		for characterIndex := 0; characterIndex < GridWidth; characterIndex++ {
			//g.Gears[row][characterIndex] = ""

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
						_ = GridLookerUpper(IndexStart, characterIndex, row, CurrentNumber)

					}
				}
			} else {
				CurrentNumber = ""
				IndexStart = -1
			}
		}
	}

	for _, row := range Gears {
		for _, col := range row {
			if col != "" {
				if strings.Count(col, ",") == 2 {
					fmt.Println(col)
					nums := strings.Split(col, ",")
					num_1, _ := strconv.Atoi(nums[1])
					num_2, _ := strconv.Atoi(nums[2])
					total += num_1 * num_2
					fmt.Println(num_1 * num_2)

				}
			}
		}
	}

	fmt.Println(total)
	fmt.Println("end")
}
