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
	Part1(lines[0])
}

func Part1(input string) {
	var InputMap map[int]int
	InputMap = make(map[int]int)

	for _, num := range strings.Split(input, " ") {
		numInt, _ := strconv.Atoi(num)
		InputMap[numInt] = InputMap[numInt] + 1
	}

	fmt.Println(InputMap)

	for i := 0; i < 75; i++ {
		fmt.Println(i)
		InputMap = Process(InputMap)
	}
	total := 0
	for _, value := range InputMap {
		total += int(value)
	}
	fmt.Println(total)

}

func Process(input map[int]int) (output map[int]int) {

	output = make(map[int]int)

	for inputNum, value := range input {
		numString := strconv.Itoa(inputNum)
		if inputNum == 0 {
			output[1] += value
		} else if len(numString)%2 == 0 {
			halfWay := len(numString) / 2
			left := numString[:halfWay]
			right := numString[halfWay:]

			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)

			output[leftInt] += value
			output[rightInt] += value

		} else {
			output[inputNum*2024] += value
		}
	}
	return output
}
