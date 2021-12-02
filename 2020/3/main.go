package main

import (
	"fmt"
	"os"
	"strings"
)

func read_input() []string {
	input, _ := os.ReadFile("input.txt")
	s := string(input)
	split := strings.Split(s, "\n")
	return split[:len(split)-1]
}

func main() {

	var totals int
	totals = 1

	type Coordinates struct {
		X int
		Y int
	}

	input := read_input()
	var variants = []Coordinates{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	for _, coord := range variants {
		tree_count := 0
		row := 0
		column := 0
		max_file := len(input)
		max_line := len(input[0])

		for column < max_file {
			// fmt.Println(string(input[column][row]))
			if string(input[column][row]) == "#" {
				tree_count += 1
			}

			row += coord.X
			column += coord.Y
			if row >= max_line {
				row = row - max_line
			}
		}
		fmt.Println(tree_count)
		totals = totals * tree_count
	}

	fmt.Println(totals)
}
