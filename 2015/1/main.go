package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

// 41 is )
// 40 is (

var count int
var iterations int

func main() {
	// Part 1
	for _, value := range input {
		if value == 40 {
			count = count + 1
		} else if value == 41 {
			count = count - 1
		}
	}
	fmt.Println("Part 1:")
	fmt.Println(count)

	// Part 2
	count = 0
	iterations = 0
	for _, value := range input {
		iterations += 1
		if value == 40 {
			count = count + 1
		} else if value == 41 {
			count = count - 1
		}
		if count == -1 {
			break
		}
	}
	fmt.Println("Part 2:")
	fmt.Println(iterations)
}
