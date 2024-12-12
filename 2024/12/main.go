package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	Part1(lines)
}

func Part1(input []string) {
	// var TotalArea map[string]int
	// var TotalPerimeter map[string]int

	// TotalArea = make(map[string]int)
	// TotalPerimeter = make(map[string]int)
	var Grid [][]string
	Grid = make([][]string, len(input))

	for i := range Grid {
		Grid[i] = make([]string, len(input[0]))
	}

	for x, line := range input {
		for y, character := range line {
			Grid[x][y] = string(character)
		}
	}

	// for x, line := range Grid {
	// 	for y, character := range line {
	// 		TotalArea[string(character)] += 1
	// 		TotalPerimeter[string(character)] += Process(x, y, Grid)
	// 	}
	// }

	var Seen map[Coords]bool
	Seen = make(map[Coords]bool)
	total := 0

	for x, line := range Grid {
		for y, _ := range line {
			total += DFS(x, y, Grid, Seen)
		}
	}
	fmt.Println(total)

	// fmt.Println("Area", TotalArea)
	// fmt.Println("Peremiter", TotalPerimeter)
	// total := 0
	// for k, v := range TotalArea {
	// 	total += v * TotalPerimeter[k]
	// }
	// fmt.Println(total)
	// fmt.Println(Grid)
}

type Coords struct {
	X int
	Y int
}

func DFS(x, y int, grid [][]string, Seen map[Coords]bool) int {
	var PeremiterTotal int
	var AreaTotal int
	S := []Coords{}
	S = append(S, Coords{x, y})
	for {
		if len(S) == 0 {
			break
		}
		// fmt.Println("S", S)
		// fmt.Println("Grid", grid)
		// fmt.Println()
		V := S[len(S)-1]
		S = S[:len(S)-1]

		if Seen[V] == true {
			continue
		}
		Seen[V] = true
		AreaTotal++

		character := grid[V.X][V.Y]
		// grid[V.X][V.Y] = "."

		if V.X > 0 {
			if string(grid[V.X-1][V.Y]) != character {
				PeremiterTotal++
			} else if string(grid[V.X-1][V.Y]) == character {
				if !slices.Contains(S, Coords{V.X - 1, V.Y}) {
					S = append(S, Coords{V.X - 1, V.Y})
				}
			}
		} else {
			PeremiterTotal++
		}
		if V.X < len(grid)-1 {
			if string(grid[V.X+1][V.Y]) != character {
				PeremiterTotal++
			} else if string(grid[V.X+1][V.Y]) == character {
				if !slices.Contains(S, Coords{V.X + 1, V.Y}) {
					S = append(S, Coords{V.X + 1, V.Y})
				}
			}
		} else {
			PeremiterTotal++
		}

		if V.Y > 0 {
			if string(grid[V.X][V.Y-1]) != character {
				PeremiterTotal++
			} else if string(grid[V.X][V.Y-1]) == character {
				if !slices.Contains(S, Coords{x, y - 1}) {
					S = append(S, Coords{V.X, V.Y - 1})
				}
			}
		} else {
			PeremiterTotal++
		}
		if V.Y < len(grid)-1 {
			if string(grid[V.X][V.Y+1]) != character {
				PeremiterTotal++
			} else if string(grid[V.X][V.Y+1]) == character {
				if !slices.Contains(S, Coords{V.X, V.Y + 1}) {
					S = append(S, Coords{V.X, V.Y + 1})
				}
			}
		} else {
			PeremiterTotal++
		}
	}
	// fmt.Println(Seen)
	// fmt.Println(grid)
	// fmt.Println("Area, peremiter", AreaTotal*PeremiterTotal)
	return AreaTotal * PeremiterTotal
}
