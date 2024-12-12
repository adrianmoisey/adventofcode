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
		fmt.Println("Start", sum, nums)
		if IsValid(sum, nums) {
			total += sum
		}
	}
	fmt.Println(total)
}

func IsValid(target int, nums []int) bool {
	if len(nums) == 1 {
		if target == nums[0] {
			return true
		} else {
			return false
		}
	}

	A := []int{nums[0] + nums[1]}
	if len(nums) > 2 {
		A = append(A, nums[2:]...)
	}

	B := []int{nums[0] * nums[1]}
	if len(nums) > 2 {
		B = append(B, nums[2:]...)
	}

	concat := strconv.Itoa(nums[0]) + strconv.Itoa(nums[1])
	num_concat, _ := strconv.Atoi(concat)
	C := []int{num_concat}
	if len(nums) > 2 {
		C = append(C, nums[2:]...)
	}

	if IsValid(target, A) {
		return true
	}
	if IsValid(target, B) {
		return true
	}
	if IsValid(target, C) {
		return true
	}

	return false

}
