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

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	Part2(lines)
}

type Rule struct {
	Before int
	After  int
}

func Part2(input []string) {

	var rules []Rule
	var updates [][]int
	// var total int

	for _, line := range input {
		if strings.Contains(line, "|") {
			lineSplit := strings.Split(line, "|")
			firstRule, _ := strconv.Atoi(lineSplit[0])
			secondRule, _ := strconv.Atoi(lineSplit[1])
			rule := Rule{Before: firstRule, After: secondRule}
			rules = append(rules, rule)
		}

		if strings.Contains(line, ",") {
			lineSplit := strings.Split(line, ",")
			update := make([]int, len(lineSplit))
			for i, value := range lineSplit {
				valueInt, _ := strconv.Atoi(value)
				update[i] = valueInt
			}
			updates = append(updates, update)
		}
	}

	invalidUpdates := [][]int{}
	for _, update := range updates {
		var valid bool
		valid = true

		valid = ValidateRule(rules, update)
		if !valid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fmt.Println(invalidUpdates)

	total := 0
	for _, update := range invalidUpdates {
		total += ValidateRuleFix(rules, update)
	}
	fmt.Println(total)
}

func ValidateRuleFix(rules []Rule, update []int) int {
	for {
		if !ValidateRule(rules, update) {
			for _, rule := range rules {

				left := slices.Index(update, rule.Before)
				right := slices.Index(update, rule.After)
				if left == -1 || right == -1 {
					continue
				}
				if left > right {
					update[left], update[right] = update[right], update[left]
				}
			}
		} else {
			middle := len(update) - 1
			middle = middle / 2

			return update[middle]
		}

	}
}

func ValidateRule(rules []Rule, update []int) bool {
	for _, rule := range rules {

		left := slices.Index(update, rule.Before)
		right := slices.Index(update, rule.After)
		if left == -1 || right == -1 {
			continue
		}
		if left > right {
			return false
		}
	}
	return true
}
