package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var (
	total     int
	cardTally map[int]int
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
	cardTally = make(map[int]int)
	lines := strings.Split(input, "\n")
	for index, line := range lines {
		index++

		if line == "" {
			break
		}

		// Add our card
		cardTally[index] = cardTally[index] + 1

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

		matches := 0
		for _, number := range MyNumbers {
			if contains(WinningNumbers, number) {
				matches++
			}
		}
		fmt.Println(index+1, matches)
		for i := index + 1; i < matches+index+1; i++ {
			cardTally[i] = cardTally[i] + (1 * cardTally[index])
			fmt.Println(i)
		}

		// 1 of card 1
		// 2 of card 2
		// 4 of card 3
		// 8 of card 4
		// 14 of card 5
		// 1 of card 6

		//fmt.Println(lines[index-1])
	}

	fmt.Println(cardTally)
	for index := range cardTally {
		total += cardTally[index]
	}
	fmt.Println("Total: ", total)

}
