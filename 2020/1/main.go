// Advent of code 2020 - Day 1
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	input_list := bytes.Split(input, []byte("\n"))

	for index, value1 := range input_list {
		for _, value2 := range input_list[index:] {
			value1_int, _ := strconv.Atoi(string(value1))
			value2_int, _ := strconv.Atoi(string(value2))
			if value1_int+value2_int == 2020 {
				fmt.Println("Part 1: ")
				fmt.Println(value1_int * value2_int)
			}

			for _, value3 := range input_list[index:] {
				value3_int, _ := strconv.Atoi(string(value3))
				if value1_int+value2_int+value3_int == 2020 {
					fmt.Println("Part 2: ")
					fmt.Println(value1_int * value2_int * value3_int)
				}
			}

		}
	}
}
