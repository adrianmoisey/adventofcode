package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	s_content := string(content)
	temp := strings.Split(s_content, "\n")

	var count int
	count = 0

	var totals []int

	for _, value := range temp {
		s_value := string(value)

		if s_value == "" {
			totals = append(totals, count)
			count = 0
		}

		int, _ := strconv.Atoi(s_value)
		count += int
	}
	sort.Ints(totals)

	lastElement := totals[len(totals)-1]
	topThree := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]

	fmt.Println("Part 1: ", lastElement)
	fmt.Println("Part 2: ", topThree)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
