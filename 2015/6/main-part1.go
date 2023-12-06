package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type Coord struct {
	x int
	y int
}

func on(from_x int, from_y int, to_x int, to_y int, grid map[Coord]bool) map[Coord]bool {

	for x := from_x; x <= to_x; x++ {
		for y := from_y; y <= to_y; y++ {
			// fmt.Println(x, y)
			grid[Coord{x, y}] = true
		}
	}
	return grid
}

func off(from_x int, from_y int, to_x int, to_y int, grid map[Coord]bool) map[Coord]bool {

	for x := from_x; x <= to_x; x++ {
		for y := from_y; y <= to_y; y++ {
			// fmt.Println(x, y)
			grid[Coord{x, y}] = false
		}
	}
	return grid
}

func toggle(from_x int, from_y int, to_x int, to_y int, grid map[Coord]bool) map[Coord]bool {

	for x := from_x; x <= to_x; x++ {
		for y := from_y; y <= to_y; y++ {
			// fmt.Println(x, y)

			current_position := grid[Coord{x, y}]
			if current_position == true {
				grid[Coord{x, y}] = false
			} else {
				grid[Coord{x, y}] = true
			}
		}
	}
	return grid
}

func main() {

	grid := make(map[Coord]bool)

	for _, word := range strings.Split(input, "\n") {

		var action, from, to string

		input_split := strings.Split(word, " ")

		if input_split[0] == "toggle" {
			action, from, to = input_split[0], input_split[1], input_split[3]
		} else {
			action, from, to = input_split[1], input_split[2], input_split[4]
		}

		from_split := strings.Split(from, ",")
		to_split := strings.Split(to, ",")

		from_x_i, _ := strconv.Atoi(from_split[0])
		from_y_i, _ := strconv.Atoi(from_split[1])

		to_x_i, _ := strconv.Atoi(to_split[0])
		to_y_i, _ := strconv.Atoi(to_split[1])

		if action == "on" {
			grid = on(from_x_i, from_y_i, to_x_i, to_y_i, grid)
		} else if action == "off" {
			grid = off(from_x_i, from_y_i, to_x_i, to_y_i, grid)
		} else if action == "toggle" {
			grid = toggle(from_x_i, from_y_i, to_x_i, to_y_i, grid)
		}
	}

	var counter int
	for _, value := range grid {
		if value == true {
			counter += 1
		}
	}

	fmt.Println("Part 1")
	fmt.Println(counter)
}
