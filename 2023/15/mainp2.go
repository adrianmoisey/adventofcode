package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var input string
var total int
var Boxes map[int][]string

func Hash(input string) (value int) {
	currentValue := 0
	for _, c := range input {
		currentValue += int(c)
		currentValue = currentValue * 17
		currentValue = currentValue % 256
	}
	return currentValue
}

func main() {
	Boxes = make(map[int][]string)

	linesInput := strings.Split(strings.TrimSpace(input), "\n")
	steps := strings.Split(linesInput[0], ",")
	for _, step := range steps {
		if strings.Contains(step, "=") {
			index := strings.Index(step, "=")
			box := Hash(step[:index])
			letter := step[:index]
			Box := Boxes[box]
			var NewLine []string

			inserted := false
			for _, line := range Box {
				//fmt.Println(step, letter+"=")

				if strings.Contains(line, letter+"=") {
					NewLine = append(NewLine, step)
					inserted = true
				} else {
					NewLine = append(NewLine, line)
				}
			}
			if !inserted {
				NewLine = append(NewLine, step)
			}
			Boxes[box] = NewLine

		} else {
			index := strings.Index(step, "-")
			box := Hash(step[:index])
			letter := step[:index]
			Box := Boxes[box]
			var NewStrings []string
			for _, line := range Box {
				//fmt.Println(letter + "=")
				if !strings.Contains(line, letter+"=") {
					NewStrings = append(NewStrings, line)
				}
			}
			Boxes[box] = NewStrings

		}

	}

	fmt.Println(len(Boxes))

	for i := 0; i <= 257; i++ {
		fmt.Println(Boxes[i])
		for slot, lens := range Boxes[i] {
			//fmt.Println(i+1, lens)
			index := strings.Index(lens, "=")
			//fmt.Println(lens[index+1:])
			LensValue, _ := strconv.Atoi(lens[index+1:])

			Sum := (i + 1) * (slot + 1) * LensValue
			fmt.Println(Sum)
			total += Sum
			//fmt.Println("Sum", Sum)
		}
	}
	fmt.Println("total", total)
}
