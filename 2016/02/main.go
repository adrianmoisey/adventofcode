package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x int32
	y int32
}

func main() {
	// Open file by line
	readFile, _ := os.Open("input.txt")
	//readFile, _ := os.Open("sample.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Configure grid and starting position
	grid := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	grid2 := [5][5]int{{0, 0, 1, 0, 0}, {0, 2, 3, 4, 0}, {5, 6, 7, 8, 9}, {0, 10, 11, 12, 0}, {0, 0, 13, 0, 0}}
	position2 := coords{0, 2}
	position := coords{1, 1}
	code1 := []string{}
	code2 := []string{}

	for fileScanner.Scan() {
		count := 0
		line := fileScanner.Text()
		for _, v := range line {
			switch string(v) {
			case "U":
				if position.y != 0 {
					position.y -= 1
				}
				if position2.y != 0 && grid2[position2.y-1][position2.x] != 0 {
					position2.y -= 1
				}

			case "D":
				if position.y != 2 {
					position.y += 1
				}
				if position2.y != 4 && grid2[position2.y+1][position2.x] != 0 {
					position2.y += 1
				}

			case "L":
				if position.x != 0 {
					position.x -= 1
				}
				if position2.x != 0 && grid2[position2.y][position2.x-1] != 0 {
					position2.x -= 1
				}

			case "R":
				if position.x != 2 {
					position.x += 1
				}
				if position2.x != 4 && grid2[position2.y][position2.x+1] != 0 {
					position2.x += 1
				}
			}
		}

		pin1 := grid[position.y][position.x]
		pin2 := grid2[position2.y][position2.x]

		var pin2_str string
		if pin2 == 10 {
			pin2_str = "A"
		} else if pin2 == 11 {
			pin2_str = "B"
		} else if pin2 == 12 {
			pin2_str = "C"
		} else if pin2 == 13 {
			pin2_str = "D"
		} else {
			pin2_str = fmt.Sprint(pin2)
		}

		code1 = append(code1, fmt.Sprint(pin1))
		code2 = append(code2, pin2_str)

		count += 1
	}
	readFile.Close()
	fmt.Println("Part 1: ", code1)
	fmt.Println("Part 2: ", code2)

}
