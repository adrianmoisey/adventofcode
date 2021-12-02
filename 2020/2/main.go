package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input() []string {
	input, _ := os.ReadFile("input.txt")
	s := string(input)
	split := strings.Split(s, "\n")
	return split[:len(split)-1]
}

func main() {
	input := read_input()

	var counter int
	for _, line := range input {

		split_line := strings.Split(line, " ")

		policy := split_line[0]
		letter := split_line[1][0:1]
		password := split_line[2]

		policy_split := strings.Split(policy, "-")
		policy_min, _ := strconv.Atoi(policy_split[0])
		policy_max, _ := strconv.Atoi(policy_split[1])

		count := strings.Count(password, letter)

		// fmt.Printf("%s, %s, %s, %d, %d, %d\n", policy, letter, password, policy_min, policy_max, count)
		// fmt.Println("")

		if count >= policy_min && count <= policy_max {
			counter += 1
		}
	}
	fmt.Printf("Part 1: %d\n", counter)

	counter = 0
	for _, line := range input {

		split_line := strings.Split(line, " ")

		policy := split_line[0]
		letter := split_line[1][0:1]

		password := split_line[2]
		policy_split := strings.Split(policy, "-")
		policy_min, _ := strconv.Atoi(policy_split[0])
		policy_max, _ := strconv.Atoi(policy_split[1])

		// fmt.Printf("%s, %s, %s, %d, %d, %d\n", policy, letter, password, policy_min, policy_max, count)
		// fmt.Println("")

		begin := string(password[policy_min-1]) == letter
		end := string(password[policy_max-1]) == letter

		/* if string(password[policy_min-1]) == letter && string(password[policy_max-1]) == letter {
			continue
		}
		if string(password[policy_min-1]) == letter || string(password[policy_max-1]) == letter {
			counter += 1
		} */
		if begin != end {
			counter += 1
		}

	}
	fmt.Printf("Part 2: %d\n", counter)
}
