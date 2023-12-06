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
	grid map[int]string
)

func main() {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}

		cubes := strings.Split(line, ":")[1]

		cubeTurns := strings.Split(cubes, ";")

		turnCubeMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, turn := range cubeTurns {

			cube := strings.Split(turn, ",")
			for _, singeCube := range cube {
				cleanCube := strings.Trim(singeCube, " ")
				splitCube := strings.Split(cleanCube, " ")
				cubeQuantity := splitCube[0]
				cubeColour := splitCube[1]
				cubeQuantityInt, _ := strconv.Atoi(string(cubeQuantity))
				if cubeQuantityInt > turnCubeMap[cubeColour] {
					turnCubeMap[cubeColour] = cubeQuantityInt
				}
			}

		}
		power := turnCubeMap["blue"] * turnCubeMap["red"] * turnCubeMap["green"]
		total += power
	}
	fmt.Println(total)
}
