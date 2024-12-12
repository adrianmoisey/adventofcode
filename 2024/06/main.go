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

	Part1(lines)
}

type Location struct {
	width  int
	height int
}

func Part1(input []string) {
	var guard Location
	var guard_direction string
	guard_direction = "^"

	for i, line := range input {
		for j, char := range line {
			if string(char) == guard_direction {
				guard.height = i
				guard.width = j
			}
		}
	}
	// fmt.Println(guard, guard_direction)

	//visitedLocations, ends := PathFind(input, guard, guard_direction, nil)
	var Locations []Location

	for i := range input {
		for j := range input[0] {
			location := Location{i, j}
			Locations = append(Locations, location)

		}
	}

	var total int
	for _, BlockLocations := range Locations {
		_, ends := PathFind(input, guard, guard_direction, &BlockLocations)
		if ends == false {
			fmt.Println(BlockLocations)
			total++
		}
	}
	fmt.Println("Total", total)
}

func PathFind(input []string, guard Location, guard_direction string, block *Location) (map[Location]bool, bool) {
	var next Location
	var visited = make(map[Location]bool)
	var visitedCorners = make(map[Location]int)

	var doesEnd bool
	doesEnd = false
	var og string

	if block != nil {
		guardRow := input[block.height]
		// fmt.Println("Original Row", guardRow)
		left := guardRow[:block.width]
		og = string(guardRow[block.width])
		if og == "#" || og == "^" {
			return visited, true
		}
		right := guardRow[block.width+1:]
		newline := left + "#" + right
		input[block.height] = newline
		// fmt.Println("New Row", newline, block.width)
		// fmt.Println("Left", left)
		// fmt.Println("Right", right)
	}

	for {
		next = guard
		if guard_direction == "^" {
			next.height -= 1
		} else if guard_direction == "v" {
			next.height += 1
		} else if guard_direction == ">" {
			next.width += 1
		} else if guard_direction == "<" {
			next.width -= 1
		}
		// fmt.Println(next.height, next.width)
		if next.height > len(input)-1 || next.width > len(input[0])-1 || next.height < 0 || next.width < 0 {
			doesEnd = true
			break
		}
		next_block := string(input[next.height][next.width])

		loopValue := 10

		if next_block == "#" {
			if guard_direction == "^" {
				guard_direction = ">"
				if visitedCorners[guard] > loopValue {
					// fmt.Println("Loop")
					break
				}
				visitedCorners[guard]++
			} else if guard_direction == ">" {
				guard_direction = "v"
				if visitedCorners[guard] > loopValue {
					// fmt.Println("Loop")
					break
				}

				visitedCorners[guard]++

			} else if guard_direction == "v" {
				guard_direction = "<"
				if visitedCorners[guard] > loopValue {
					// fmt.Println("Loop")
					break
				}

				visitedCorners[guard]++

			} else if guard_direction == "<" {
				guard_direction = "^"
				if visitedCorners[guard] > loopValue {
					// fmt.Println("Loop")
					break
				}

				visitedCorners[guard]++
			}
		} else {
			guard = next
			visited[guard] = true
		}
	}

	// if doesEnd == false {
	// 	fmt.Println(next)
	// 	fmt.Println(input)
	// }

	guardRow := input[block.height]
	left := guardRow[:block.width]
	right := guardRow[block.width+1:]
	newline := left + og + right
	input[block.height] = newline

	return visited, doesEnd

}

// 2023 too high
