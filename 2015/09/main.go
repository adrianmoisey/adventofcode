package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

//go:embed sample.txt
var input string
var totalp1 int
var grid map[string]Node

type Tuple struct {
	name     string
	distance int
}

type Node struct {
	Name       string
	Neighbours []Neighbour
}

func (n *Node) addNeighbour(name string, distance int) {
	newNeighbour := Neighbour{Name: name, Distance: distance}
	n.Neighbours = append(n.Neighbours, newNeighbour)
}

type Neighbour struct {
	Name     string
	Distance int
}

func part1(input string) {
	grid = make(map[string]Node)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		splitter := strings.Split(line, " = ")
		Citys := strings.Split(splitter[0], " to ")
		Distance, _ := strconv.Atoi(splitter[1])
		Source := Citys[0]
		Destination := Citys[1]

		fmt.Println("Distance", Distance, "Source", Source, "Destination", Destination)

		node, exists := grid[Source]
		if !exists {
			node = Node{Name: Source}
			grid[Source] = node
		}
		node.addNeighbour(Destination, Distance)

		grid[node.Name] = node
	}

	unvisited := mapset.NewSet[Tuple]()

	for key := range grid {
		fmt.Println(key)
		unvisited.Add(Tuple{key, 99999999999})
	}
	unvisited.Add(Tuple{"London", 0})
	fmt.Println(unvisited)

	fmt.Println(totalp1)

}

func main() {
	part1(input)
}
