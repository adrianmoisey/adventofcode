package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

var Grid map[string][]Antenna
var AntiGrid map[Antenna]bool

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	Part1(lines)
}

type Antenna struct {
	X int // Left/Right
	Y int // Down/Up
}

func Part1(input []string) {
	Grid = make(map[string][]Antenna)
	AntiGrid = make(map[Antenna]bool)

	for x, line := range input {
		for y, c := range line {
			if string(c) != "." {
				Grid[string(c)] = append(Grid[string(c)], Antenna{x, y})
				AntiGrid[Antenna{x, y}] = true

			}
		}
	}

	fmt.Println(Grid)
	for C, locations := range Grid {
		for _, location := range locations {
			for _, otherlocation := range locations {
				if location == otherlocation {
					continue
				}

				Primary := location
				Secondary := otherlocation
				for {
					fmt.Println(C)
					Difference := Antenna{Secondary.X - Primary.X, Secondary.Y - Primary.Y}
					Anti := Antenna{Secondary.X + Difference.X, Secondary.Y + Difference.Y}
					fmt.Println("Primary", Primary)
					fmt.Println("Secondary", Secondary)
					fmt.Println("Difference", Difference)
					fmt.Println("Anti", Anti)

					if Anti.X < 0 || Anti.Y < 0 || Anti.X >= len(input) || Anti.Y >= len(input[0]) {
						fmt.Println("Break on", Anti)
						break
					}
					AntiGrid[Anti] = true

					Primary = Secondary
					Secondary = Anti
				}
			}
		}
	}
	fmt.Println(AntiGrid)
	// for k, _ := range AntiGrid {
	// 	fmt.Println("A", k)
	// }
	fmt.Println(len(AntiGrid))

	for x, line := range input {
		for y, c := range line {
			Location := Antenna{x, y}
			if AntiGrid[Location] {
				fmt.Print("#")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}

}
