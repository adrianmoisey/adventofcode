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
var GalaxyPositions []string
var Total int
var ExpansionSize = 1000000 - 1

func main() {
	var HorizontalExpansions []int
	var VerticalExpansions []int

	linesInput := strings.Split(strings.TrimSpace(input), "\n")
	for i, line := range linesInput {
		if !strings.Contains(line, "#") {
			HorizontalExpansions = append(HorizontalExpansions, i)
		}
	}

	for i := 0; i < len(linesInput); i++ {
		column := make([]string, 0)
		for _, line := range linesInput {
			column = append(column, string(line[i]))
		}
		if !slices.Contains(column, "#") {
			VerticalExpansions = append(VerticalExpansions, i)
		}

	}

	fmt.Println(HorizontalExpansions, VerticalExpansions)

	for iline, line := range linesInput {
		for icharacter, character := range line {
			if string(character) == "#" {
				fmt.Println("Pos: ", iline, icharacter)

				MultiplierX := 0
				for _, EX := range HorizontalExpansions {
					if iline >= EX {
						MultiplierX += ExpansionSize
					}
				}
				fmt.Println("MX: ", MultiplierX)

				MultiplierY := 0
				for _, EY := range VerticalExpansions {
					if icharacter >= EY {
						MultiplierY += ExpansionSize
					}
				}
				fmt.Println("MY: ", MultiplierY)
				fmt.Println("New Pos: ", iline+MultiplierX, icharacter+MultiplierY)

				X := strconv.Itoa(iline + MultiplierX)
				Y := strconv.Itoa(icharacter + MultiplierY)

				Position := X + "," + Y
				GalaxyPositions = append(GalaxyPositions, Position)
			}
		}
	}
	fmt.Println("Galaxy Positions")
	fmt.Println(GalaxyPositions)

	for i := 0; i <= len(GalaxyPositions); i++ {
		for j := i + 1; j <= len(GalaxyPositions)-1; j++ {
			fmt.Println(GalaxyPositions[i], GalaxyPositions[j])
			GalaxyA := strings.Split(GalaxyPositions[i], ",")
			GalaxyB := strings.Split(GalaxyPositions[j], ",")

			GalaxyAX, _ := strconv.Atoi(GalaxyA[0])
			GalaxyAY, _ := strconv.Atoi(GalaxyA[1])

			GalaxyBX, _ := strconv.Atoi(GalaxyB[0])
			GalaxyBY, _ := strconv.Atoi(GalaxyB[1])

			GalaxyXSum := Abs(GalaxyAX - GalaxyBX)
			GalaxyYSum := Abs(GalaxyAY - GalaxyBY)

			fmt.Println(Abs(GalaxyXSum + GalaxyYSum))
			Total += Abs(GalaxyXSum + GalaxyYSum)
		}

	}
	fmt.Println(Total)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
