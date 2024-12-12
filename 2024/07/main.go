package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	Part1(lines)
}

func Part1(input []string) {
	var total int
	for _, line := range input {
		splitLine := strings.Split(line, ": ")
		var nums []int
		for _, num := range strings.Split(splitLine[1], " ") {
			numi, _ := strconv.Atoi(string(num))
			nums = append(nums, numi)
		}
		sum, _ := strconv.Atoi(splitLine[0])
		fmt.Println(sum, nums)

		n := len(nums) - 1
		totalCombinations := 1
		base := 3 // Possible values: 0, 1, 2

		// Calculate the total number of combinations
		for i := 0; i < n; i++ {
			totalCombinations *= base
		}

		results := [][]int{}

		for i := 0; i < totalCombinations; i++ {
			combination := make([]int, n)
			num := i

			// Generate the combination for this number
			for j := n - 1; j >= 0; j-- {
				combination[j] = num % base // Extract the digit in base 3
				num /= base                 // Move to the next digit
			}

			results = append(results, combination)
		}

		// Print all combinations
		for _, operators := range results {
			// fmt.Println(operators)
			if SumA(sum, nums, operators) {
				total += sum
				break
			}
		}

	}
	fmt.Println(total)
}

func SumA(sum int, nums []int, operators []int) bool {
	var total int

	for i, num := range nums {
		if i == 0 {
			total = num
		} else {
			if operators[i-1] == 0 {
				total = total + num

			} else if operators[i-1] == 1 {
				total = total * num
			} else if operators[i-1] == 2 {
				concat := strconv.Itoa(total) + strconv.Itoa(num)

				total, _ = strconv.Atoi(concat)
			}
		}
	}
	if total == sum {
		return true
	}

	return false
}
