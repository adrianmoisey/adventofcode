package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	cubeMap = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	total int
)

func main() {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		fmt.Println()
		fmt.Println(line)
		game := strings.Split(line, ":")[0]
		gamdId := strings.Split(game, " ")[1]

		cubes := strings.Split(line, ":")[1]

		cubeTurns := strings.Split(cubes, ";")

		doesItFit := true

		for _, turn := range cubeTurns {

			fmt.Println("Cube Turn...")
			fmt.Println(turn)
			cube := strings.Split(turn, ",")
			for _, singeCube := range cube {
				cleanCube := strings.Trim(singeCube, " ")
				splitCube := strings.Split(cleanCube, " ")
				cubeQuantity := splitCube[0]
				cubeColour := splitCube[1]
				fmt.Println("Single Cube: ", cubeQuantity, cubeColour)
				cubeQuantityInt, _ := strconv.Atoi(string(cubeQuantity))
				if cubeQuantityInt > cubeMap[cubeColour] {
					doesItFit = false
				}
				fmt.Println(cubeMap[cubeColour])
			}
		}

		if doesItFit == true {
			value, _ := strconv.Atoi(string(gamdId))
			total += value
		}

		fmt.Println(gamdId)
		fmt.Println(cubeTurns)

	}
	fmt.Println(total)

}
