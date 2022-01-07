package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type coord struct {
	x int
	y int
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	// Part 1
	x, y := 0, 0
	positions := make(map[coord]int)

	for _, value := range input {
		value_s := string(value)
		fmt.Println(value_s)
		if value_s == "^" {
			x += 1
		} else if value_s == "v" {
			x -= 1
		} else if value_s == "<" {
			y -= 1
		} else if value_s == ">" {
			y += 1
		}
		a := coord{x, y}
		current_value := positions[a]
		current_value += 1
		positions[a] = current_value
	}
	fmt.Println("Part 1")
	fmt.Println(len(positions))

	// Part 2
	santa_x, santa_y := 0, 0
	robo_x, robo_y := 0, 0

	part2_positions := make(map[coord]int)

	santa_turn := true

	for _, value := range input {
		value_s := string(value)
		if santa_turn == true {
			if value_s == "^" {
				santa_x += 1
			} else if value_s == "v" {
				santa_x -= 1
			} else if value_s == "<" {
				santa_y -= 1
			} else if value_s == ">" {
				santa_y += 1
			}

			santa_a := coord{santa_x, santa_y}
			current_value := part2_positions[santa_a]
			current_value += 1
			part2_positions[santa_a] = current_value

			santa_turn = false
		} else if santa_turn == false {
			if value_s == "^" {
				robo_x += 1
			} else if value_s == "v" {
				robo_x -= 1
			} else if value_s == "<" {
				robo_y -= 1
			} else if value_s == ">" {
				robo_y += 1
			}

			robo_a := coord{robo_x, robo_y}
			current_value := part2_positions[robo_a]
			current_value += 1
			part2_positions[robo_a] = current_value

			santa_turn = true
		}

	}
	fmt.Println("Part 2")
	fmt.Println(len(part2_positions))

}
