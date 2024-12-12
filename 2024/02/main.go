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
	report := [][]int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		line_s := strings.Split(line, " ")
		line_i := []int{}
		for _, num := range line_s {
			num_int, _ := strconv.Atoi(string(num))
			line_i = append(line_i, num_int)
		}
		report = append(report, line_i)
	}

	safeReports := 0

	for _, line := range report {
		if IsSafe(line) {
			safeReports++
		} else {
			fmt.Println("Fresh line", line)
			for i := range line {
				new_line := make([]int, len(line))

				_ = copy(new_line, line)

				new_line = slices.Delete(new_line, i, i+1)
				fmt.Println("Modified line", new_line)
				if IsSafe(new_line) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println(safeReports)
}

func IsSafe(line []int) bool {
	if line[0] > line[1] {
		slices.Reverse(line)
	}
	for i, num := range line {

		if i+1 == len(line) {
			return true
		}

		next := line[i+1]
		if num < next && next <= num+3 {
			continue
		} else {
			break
		}
	}
	return false
}
