package main

import (
	_ "embed"
	"fmt"
	"sort"
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

var total int
var ribbon_length_total int

func main() {
	// Part 1
	for _, value := range strings.Split(input, "\n") {
		value_s := string(value)
		present_list := strings.Split(value_s, "x")
		length, _ := strconv.Atoi(present_list[0])
		width, _ := strconv.Atoi(present_list[1])
		height, _ := strconv.Atoi(present_list[2])

		var sides []int
		all_sides := []int{length, width, height}
		sort.Ints(all_sides)

		ribbon_length_total += all_sides[0] + all_sides[0] + all_sides[1] + all_sides[1]
		ribbon_length_total += length * width * height

		sides = append(sides, 2*length*width)
		sides = append(sides, 2*width*height)
		sides = append(sides, 2*height*length)

		sitesTotal := 0
		smallestNumber := sides[0]

		for _, element := range sides {
			sitesTotal += element
			if element < smallestNumber {
				smallestNumber = element
			}
		}

		total += sitesTotal
		total += smallestNumber / 2

	}
	fmt.Println("Part 1:")
	fmt.Println(total)
	fmt.Println("Part 2:")
	fmt.Println(ribbon_length_total)

}
