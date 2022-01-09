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

	contains := 0
	for _, s := range input {
		if strings.Contains(string(s), "a") {
			contains += 1
		}
		if strings.Contains(string(s), "e") {
			contains += 1
		}
		if strings.Contains(string(s), "i") {
			contains += 1
		}
		if strings.Contains(string(s), "o") {
			contains += 1
		}
		if strings.Contains(string(s), "u") {
			contains += 1
		}
	}
	if contains < 3 {
		return false
	}

	var rule2 bool
	rule2 = false
	input_len := len(input)
	for i, s := range input {
		if i+1 < input_len {
			if input[i+1] == byte(s) {
				rule2 = true
			}
		}
	}

	// Rule 3
	substrings = []string{"ab", "cd", "pq", "xy"}
	for _, s := range substrings {
		if strings.Contains(input, s) {
			return false
		}
	}

	return rule2
}

func main() {
	// Part 1

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
