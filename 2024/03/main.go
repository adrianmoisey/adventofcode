package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var enabled bool

func main() {
	total := 0
	enabled = true

	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, mul := range matches {
			if mul == "don't()" {
				enabled = false
				continue
			} else if mul == "do()" {
				enabled = true
				continue
			}
			if enabled {
				total += Multiply(mul)
			}
		}
	}
	fmt.Println(total)
}

func Multiply(input string) int {
	input = input[4 : len(input)-1]
	digits := strings.Split(input, ",")

	intArray := make([]int, len(digits))

	for i, s := range digits {
		// Convert string to int
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", s, err)
			continue
		}
		intArray[i] = val
	}

	return intArray[0] * intArray[1]
}
