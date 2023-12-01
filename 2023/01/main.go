package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var total int
	lines := strings.Split(input, "\n")

	for _, line := range lines {

		var firstInt string
		var lastInt string

		for _, character := range line {
			if _, err := strconv.Atoi(string(character)); err == nil {
				if firstInt == "" {
					firstInt = string(character)
				}
				lastInt = string(character)
			}
		}
		value := firstInt + lastInt
		intValue, _ := strconv.Atoi(value)

		total += intValue
	}
	fmt.Println("Part1: ")
	fmt.Println(total)
}
