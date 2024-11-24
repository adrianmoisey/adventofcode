package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var input string
var Grid map[int]map[int]int

func main() {
	linesInput := strings.Split(strings.TrimSpace(input), "\n")
	Max := [2]int{len(linesInput), len(linesInput[1])}

	// Populate the Map
	Grid = make(map[int]map[int]int)
	for x, line := range linesInput {
		Grid[x] = make(map[int]int)
		for y, character := range line {
			characterI, _ := strconv.Atoi(string(character))
			Grid[x][y] = characterI
		}
	}
	Goal := [2]int{len(Grid), len(Grid[0])}
	Start := [2]int{0, 0}

	frontier := make(PriorityQueue, 0)

	heap.Init(&frontier)
	heap.Push(&frontier, &Node{value: Start, priority: 0})

	came_from := make(map[[2]int][2]int)
	cost_so_far := make(map[[2]int]int)

	came_from[Start] = Start
	cost_so_far[Start] = 0

	for frontier.Len() != 0 {
		NextNode := heap.Pop(&frontier).(*Node)
		//NextNodeCasted := NextNode.(Node)
		current := NextNode.value

		if current == Goal {
			break
		}

		for _, next := range Neighbours(current, Max) {
			new_cost := cost_so_far[current] + Grid[next[0]][next[1]]

			_, present := cost_so_far[next]
			if !present || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				frontier.Push(&Node{value: next, priority: new_cost})
				came_from[next] = current
			}
		}
	}
	fmt.Println((len(Grid) + 1) * (len(Grid[0]) + 1))
	fmt.Println(len(came_from))
	fmt.Println(came_from)
	fmt.Println(cost_so_far[Goal])
	fmt.Println(Grid)

}

func Neighbours(current [2]int, Max [2]int) (neighbours [][2]int) {
	N := [2]int{current[0] - 1, current[1]}
	S := [2]int{current[0] + 1, current[1]}
	E := [2]int{current[0], current[1] - 1}
	W := [2]int{current[0], current[1] + 1}

	for _, node := range [][2]int{N, S, E, W} {
		if (node[0] >= 0 && node[0] <= Max[0]) && (node[1] >= 0 && node[1] <= Max[1]) {
			neighbours = append(neighbours, node)
		}
	}
	return neighbours
}

// An Item is something we manage in a priority queue.
type Node struct {
	value    [2]int
	priority int
	index    int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(node *Node, value [2]int, priority int) {
	node.value = value
	node.priority = priority
	heap.Fix(pq, node.index)
}
