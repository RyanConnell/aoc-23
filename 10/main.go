package main

import (
	"fmt"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

type Node struct {
	pipe        Pipe
	x, y        int
	distance    int
	visitedFrom map[int]struct{}
}

func NewNode(x, y int, pipe Pipe) *Node {
	return &Node{x: x, y: y, pipe: pipe, distance: -1, visitedFrom: make(map[int]struct{})}
}

type Status struct {
	prune bool
}

func walk(node *Node, nodeMap [][]*Node, steps, input int) *Status {
	if node == nil {
		return &Status{prune: true}
	}
	// Ensure that this pipe allows input from the given direction.
	if _, ok := node.pipe.directionSet[input]; !ok {
		return &Status{prune: true}
	}
	if node.pipe.char == 'S' {
		node.distance = 0
		return &Status{}
	}

	getNode := func(x, y int) *Node {
		if x < 0 || y < 0 || y >= len(nodeMap) || x >= len(nodeMap[y]) {
			return nil
		}
		return nodeMap[y][x]
	}

	if node.distance != -1 && node.distance < steps {
		return &Status{}
	}
	node.distance = steps
	nextDir := node.pipe.output(input)
	var status *Status
	switch nextDir {
	case NORTH:
		status = walk(getNode(node.x, node.y-1), nodeMap, steps+1, SOUTH)
	case SOUTH:
		status = walk(getNode(node.x, node.y+1), nodeMap, steps+1, NORTH)
	case EAST:
		status = walk(getNode(node.x+1, node.y), nodeMap, steps+1, WEST)
	case WEST:
		status = walk(getNode(node.x-1, node.y), nodeMap, steps+1, EAST)
	}

	if status.prune {
		node.pipe = possiblePipes['.']
		node.distance = -1
	}

	return status
}

type Pipe struct {
	char         rune
	directionSet map[int]struct{}
	directions   []int
}

func NewPipe(char rune, directions []int) Pipe {
	directionSet := make(map[int]struct{})
	for _, dir := range directions {
		directionSet[dir] = struct{}{}
	}
	return Pipe{char, directionSet, directions}
}

func (p *Pipe) output(input int) int {
	if input != p.directions[0] {
		return p.directions[0]
	}
	return p.directions[1]
}

var possiblePipes = map[rune]Pipe{
	'|': NewPipe('|', []int{NORTH, SOUTH}),
	'-': NewPipe('-', []int{EAST, WEST}),
	'L': NewPipe('L', []int{NORTH, EAST}),
	'J': NewPipe('J', []int{NORTH, WEST}),
	'7': NewPipe('7', []int{SOUTH, WEST}),
	'F': NewPipe('F', []int{SOUTH, EAST}),
	'S': NewPipe('S', []int{NORTH, EAST, SOUTH, WEST}),
	'.': NewPipe('.', nil),
}

func solve(lines []string) (int, error) {
	var startNode *Node
	nodeMap := make([][]*Node, len(lines))
	for y, line := range lines {
		nodeMap[y] = make([]*Node, len(line))
		for x, char := range line {
			nodeMap[y][x] = NewNode(x, y, possiblePipes[char])
			if char == 'S' {
				startNode = nodeMap[y][x]
			}
		}
	}

	getNode := func(x, y int) *Node {
		if x < 0 || y < 0 || y >= len(nodeMap) || x >= len(nodeMap[y]) {
			return nil
		}
		return nodeMap[y][x]
	}

	explore := func(node *Node, from int) {
		walk(node, nodeMap, 1, from)
	}

	// Maybe explore SOUTH
	explore(getNode(startNode.x, startNode.y+1), NORTH)

	// Maybe explore NORTH
	explore(getNode(startNode.x, startNode.y-1), SOUTH)

	// Maybe explore EAST
	explore(getNode(startNode.x+1, startNode.y), WEST)

	// Maybe explore WEST
	explore(getNode(startNode.x-1, startNode.y), EAST)

	var furthestNode *Node
	for _, row := range nodeMap {
		for _, node := range row {
			if node.distance == -1 {
				continue
			}
			if furthestNode == nil || furthestNode.distance < node.distance {
				furthestNode = node
			}
		}
	}

	return furthestNode.distance, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
