package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var substrings []string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func isnice(input string) bool {

	var rule1 bool
	rule1 = false
	input_len := len(input)
	for i, _ := range input {
		if i+3 < input_len {
			first_string := input[i : i+2]
			second_string := input[i+2:]
			if strings.Contains(second_string, first_string) {
				rule1 = true
			}
		}
	}

	var rule2 bool
	rule2 = false
	for i, _ := range input {
		if i+2 < input_len {
			first_string := input[i]
			second_string := input[i+2]
			if first_string == second_string {
				rule2 = true
			}
		}
	}

	if rule1 == true && rule2 == true {
		return true
	}
	return false
}

func main() {
	// Part 1

	//input = "qjhvhtzxzqqjkmpb"

	count := 0
	for _, word := range strings.Split(input, "\n") {
		// fmt.Println(string(word))
		result := isnice(string(word))

		if result == true {
			count += 1
		}
	}

	fmt.Println(count)

}
