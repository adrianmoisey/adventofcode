package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int32
	y int32
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var pointing string = "N"
	var x, y int32
	var part2_answer int32

	visited := make(map[coords]bool)

	N := map[string]string{
		"L": "W",
		"R": "E",
	}
	E := map[string]string{
		"L": "N",
		"R": "S",
	}
	S := map[string]string{
		"L": "E",
		"R": "W",
	}
	W := map[string]string{
		"L": "S",
		"R": "N",
	}

	input, _ := os.ReadFile("input.txt")
	//input, _ := os.ReadFile("samplea.txt")
	inputslice := bytes.Split(input, []byte{','})

	for _, value := range inputslice {
		stripped_value := strings.TrimSpace(string(value))
		direction := string(stripped_value[0])
		distance_8, _ := strconv.Atoi(string(stripped_value[1:]))

		distance := int32(distance_8)

		switch pointing {
		case "N":
			direction = N[direction]
		case "E":
			direction = E[direction]
		case "S":
			direction = S[direction]
		case "W":
			direction = W[direction]
		}
		pointing = direction

		switch pointing {
		case "N":
			for i := x + 1; i <= x+distance; i++ {
				pos := coords{i, y}
				recordCoord(pos, visited, &part2_answer)
			}
			x += distance

		case "E":
			for i := y + 1; i <= y+distance; i++ {
				pos := coords{x, i}
				recordCoord(pos, visited, &part2_answer)
			}
			y += distance
		case "S":
			for i := x - 1; i >= x-distance; i-- {
				pos := coords{i, y}
				recordCoord(pos, visited, &part2_answer)
			}
			x -= distance
		case "W":
			for i := y - 1; i >= y-distance; i-- {
				pos := coords{x, i}
				recordCoord(pos, visited, &part2_answer)
			}
			y -= distance
		}
	}

	fmt.Printf("Part 1: %d\n", Abs(x)+Abs(y))
	fmt.Printf("Part 2: %d\n", part2_answer)
}

func recordCoord(position coords, visited map[coords]bool, part2_answer *int32) {
	if visited[position] {
		if *part2_answer == 0 {
			*part2_answer = Abs(position.x) + Abs(position.y)
		}
	}
	visited[position] = true
}
