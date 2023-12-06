package main

import (
	_ "embed"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

var (
	total int
)

func contains(s []string, elem string) bool {
	for _, a := range s {
		if a == elem {
			return true
		}
	}
	return false
}

func main() {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		splits := strings.Split(line[8:], " |")

		winningNumbers := splits[0]
		myNumbers := splits[1]

		WinningSplits := strings.Split(winningNumbers, " ")
		myNumbersSplits := strings.Split(myNumbers, " ")

		var WinningNumbers []string
		for _, number := range WinningSplits {
			if number == "" {
				continue
			}
			WinningNumbers = append(WinningNumbers, number)
		}

		var MyNumbers []string
		for _, number := range myNumbersSplits {
			if number == "" {
				continue
			}

			MyNumbers = append(MyNumbers, number)
		}

		fmt.Println(WinningNumbers, MyNumbers)

		prize := 0
		for _, number := range MyNumbers {
			if slices.Contains(WinningNumbers, number) {
				//if contains(WinningNumbers, number) {
				if prize == 0 {
					prize += 1
				} else {
					prize += prize
				}
			}
		}
		fmt.Println(prize)
		total += prize
	}

	fmt.Println("Total: ", total)
}
