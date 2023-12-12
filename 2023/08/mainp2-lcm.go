package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var (
	total         int
	Mapping       map[string]Node
	StartingNodes []Node
)

type Node struct {
	Name string

	Left  string
	Right string
}

func (n *Node) Next(letter string) string {
	if letter == "L" {
		return n.Left
	}
	return n.Right

}

func main() {
	Mapping = make(map[string]Node)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	Moves := lines[0]
	for _, line := range lines[2:] {
		split := strings.Split(line, " ")
		Name := split[0]
		Left := split[2][1:4]
		Right := split[3][:3]
		node := Node{Left: Left, Right: Right, Name: Name}
		Mapping[Name] = node
		if string(Name[2]) == "A" {
			StartingNodes = append(StartingNodes, node)
		}
	}

	var counts []int

	for _, node := range StartingNodes {
		count := 1
		i := 0
		for {
			if i == len(Moves) {
				i = 0
			}
			nextNodeString := node.Next(string(Moves[i]))
			if string(nextNodeString[2]) == "Z" {
				break
			}
			node = Mapping[nextNodeString]
			count++
			i++
		}
		counts = append(counts, count)
	}

	a := 1
	for _, num := range counts {
		b := num
		a = calculateLCM(a, b)
	}

	fmt.Println("Total: ", a)
}

// These functions stolen from ChatGPT
func calculateLCM(a, b int) int {
	// Use the GCD function from the math package
	gcd := int(calculateGCD(a, b))

	// Calculate LCM using the formula
	lcm := (a * b) / gcd

	return lcm
}

func calculateGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
