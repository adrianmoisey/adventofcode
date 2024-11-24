package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var Grid [][2]int

// X, Y
var Current [2]int = [2]int{0, 0}

func main() {
	linesInput := strings.Split(strings.TrimSpace(input), "\n")
	Grid = append(Grid, Current)

	border := 0

	// R 6 (#70c710)

	for _, line := range linesInput {
		SplitLine := strings.Split(line, " ")
		Hex := SplitLine[2]

		HexDistance := Hex[2:7]
		Distance64, _ := strconv.ParseInt(HexDistance, 16, 32)
		Direction := Hex[7:8]

		// Fixme
		//DistanceS := SplitLine[1]
		//Distance, _ := strconv.Atoi(DistanceS)
		Distance := int(Distance64)

		border += Distance

		if Direction == "3" {
			NewPos := [2]int{Current[0] - Distance, Current[1]}
			Grid = append(Grid, NewPos)
			Current = NewPos
		} else if Direction == "1" {
			NewPos := [2]int{Current[0] + Distance, Current[1]}
			Grid = append(Grid, NewPos)
			Current = NewPos

		} else if Direction == "2" {
			NewPos := [2]int{Current[0], Current[1] - Distance}
			Grid = append(Grid, NewPos)
			Current = NewPos

		} else if Direction == "0" {
			NewPos := [2]int{Current[0], Current[1] + Distance}
			Grid = append(Grid, NewPos)
			Current = NewPos
		}
	}

	Area := 0

	// fmt.Println(Grid)

	// Shoelace
	for i := range Grid {
		if i == len(Grid)-1 {
			break
		}
		//fmt.Println(Grid[i])
		//fmt.Println(Grid[i+1])

		Green := Grid[i][0] * Grid[i+1][1]
		Red := Grid[i][1] * Grid[i+1][0]

		Area += (Green - Red)
	}

	if Area < 0 {
		Area = -Area
	}
	Area = Area / 2

	fmt.Println(Area, border)

	// internal = Area - (border/2) + 1
	// internal = 42 - (38/2) + 1
	// Answer is 62

	fmt.Println(Area - (border / 2) + 1 + border)

}
