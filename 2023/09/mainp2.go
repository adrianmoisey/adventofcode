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
var Line []int
var Total int

func differences(a []int) int {
	if slices.Max(a) == 0 && slices.Min(a) == 0 {
		return 0
	}

	var differenceInt []int

	for i := range a {
		if len(a) == i+1 {
			break
		}
		num := a[i+1] - a[i]
		differenceInt = append(differenceInt, num)
	}
	return a[len(a)-1] + differences(differenceInt)
}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var intLine []int

		for _, num := range numbers {
			numI, _ := strconv.Atoi(num)
			intLine = append(intLine, numI)
		}

		//slices.Reverse(intLine)

		Total += differences(intLine)
	}
	fmt.Println(Total)
}
