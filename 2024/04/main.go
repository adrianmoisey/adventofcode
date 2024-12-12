package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	Part2(lines)
}

func Part2(input []string) {

	var verticalTotal int
	rowLength := len(input[0]) - 1
	columnLength := len(input) - 1

	for y, line := range input {
		for x, character := range line {
			if character == 'A' {
				// Up Left
				if y-1 >= 0 && x-1 >= 0 && y+1 <= columnLength && x-1 >= 0 && y+1 <= columnLength && x+1 <= rowLength {
					upLeft := string(input[y-1][x-1])
					upRight := string(input[y-1][x+1])
					downLeft := string(input[y+1][x-1])
					downRight := string(input[y+1][x+1])

					if upLeft == "M" && downRight == "S" || upLeft == "S" && downRight == "M" {
						if upRight == "M" && downLeft == "S" {
							verticalTotal++
						} else if upRight == "S" && downLeft == "M" {
							verticalTotal++
						}
					}

					fmt.Println(upLeft, upRight, downLeft, downRight)
				}
			}
		}
	}

	fmt.Println("Part 1", verticalTotal)

}
