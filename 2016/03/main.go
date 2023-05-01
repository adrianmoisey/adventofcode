package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type triangle struct {
	x int64
	y int64
	z int64
}

func isValidTriangle(triangle triangle) (valid bool) {
	valid = triangle.x+triangle.y > triangle.z && triangle.x+triangle.z > triangle.y && triangle.z+triangle.y > triangle.x
	return
}

func main() {
	// Open file by line
	readFile, _ := os.Open("input.txt")
	//readFile, _ := os.Open("sample.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	part1 := 0
	part2 := 0
	part2_1 := triangle{}
	part2_2 := triangle{}
	part2_3 := triangle{}

	for i := 0; fileScanner.Scan(); i++ {
		// Part 1
		line := fileScanner.Text()
		words := strings.Fields(line)
		x, _ := strconv.ParseInt(words[0], 0, 64)
		y, _ := strconv.ParseInt(words[1], 0, 64)
		z, _ := strconv.ParseInt(words[2], 0, 64)
		if x+y > z && x+z > y && z+y > x {
			part1 += 1
		}

		// Part 2
		if i%3 == 0 {
			part2_1.x = x
			part2_2.x = y
			part2_3.x = z

		} else if i%3 == 1 {
			part2_1.y = x
			part2_2.y = y
			part2_3.y = z

		} else if i%3 == 2 {
			part2_1.z = x
			part2_2.z = y
			part2_3.z = z

			// Last part of triangle, do the calculation
			if isValidTriangle(part2_1) {
				part2 += 1
			}
			if isValidTriangle(part2_2) {
				part2 += 1
			}

			if isValidTriangle(part2_3) {
				part2 += 1
			}

		}

	}
	readFile.Close()
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)

}
