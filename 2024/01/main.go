package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	total []int
)

func main() {
	Left := []int{}
	Right := []int{}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, v := range lines {
		splitted := strings.Split(v, " ")

		first := splitted[0]
		last := splitted[len(splitted)-1]

		s_first, _ := strconv.Atoi(string(first))
		s_last, _ := strconv.Atoi(string(last))

		Left = append(Left, s_first)
		Right = append(Right, s_last)
	}

	slices.Sort(Right)
	slices.Sort(Left)

	similarity_total := 0
	distance_total := 0

	for i, v := range Left {
		distance := v - Right[i]
		if distance < 0 {
			distance = distance * -1
		}
		distance_total += distance

		occurences := 0

		for _, right := range Right {
			if v == right {
				occurences++
			}
		}
		similarity_total += v * occurences

	}
	fmt.Println("Part 1", distance_total)
	fmt.Println("Part 2", similarity_total)

}
