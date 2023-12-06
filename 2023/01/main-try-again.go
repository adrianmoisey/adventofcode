package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input      string
	totalPart1 int
	totalPart2 int
)

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

	for _, line := range strings.Split(input, "\n") {
		var part1Ints []string
		var part2Ints []string

		for index, character := range line {
			// If a number
			_, err := strconv.Atoi(string(character))
			if err == nil {
				part1Ints = append(part1Ints, string(character))
				part2Ints = append(part2Ints, string(character))

			}
			// if a word
			for k := range lookup {
				if strings.HasPrefix(line[index:], k) {
					part2Ints = append(part2Ints, lookup[k])
				}
			}

		}

		if len(part1Ints) > 0 {
			total, _ := strconv.Atoi(part1Ints[0] + part1Ints[len(part1Ints)-1])
			totalPart1 += total
		}
		if len(part2Ints) > 0 {
			total, _ := strconv.Atoi(part2Ints[0] + part2Ints[len(part2Ints)-1])
			totalPart2 += total
		}

	}
	fmt.Println(totalPart1)
	fmt.Println(totalPart2)

}
