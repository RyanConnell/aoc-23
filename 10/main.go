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
	'X': NewPipe('X', nil),
}

func generateAndExploreMap(lines []string) [][]*Node {
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

	return nodeMap
}

func solve(lines []string) (int, error) {
	nodeMap := generateAndExploreMap(lines)

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
	nodeMap := generateAndExploreMap(lines)

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

	isolatedMap := createIsolatedMap(furthestNode, nodeMap)
	zoomedMap := createZoomedMap(isolatedMap)
	floodZoomedMap(zoomedMap)
	squashedMap := squashZoomedMap(zoomedMap)

	// Count remaining tiles in our squashed map
	var count int
	for _, row := range squashedMap {
		for _, solid := range row {
			if !solid {
				count++
			}
		}
	}
	return count, nil
}

func floodZoomedMap(zoomedMap [][]bool) {
	var flood func(int, int) // Declared without := so we can reference it from within itself.

	flood = func(x, y int) {
		if zoomedMap[y][x] {
			return
		}

		zoomedMap[y][x] = true

		if x+1 < len(zoomedMap[y]) {
			flood(x+1, y)
		}
		if x-1 >= 0 {
			flood(x-1, y)
		}
		if y+1 < len(zoomedMap) {
			flood(x, y+1)
		}
		if y-1 >= 0 {
			flood(x, y-1)
		}
	}

	// Flood inwards from the outskirts of the map
	for y := 0; y < len(zoomedMap); y++ {
		flood(0, y)                   // Left edge
		flood(len(zoomedMap[y])-1, y) // Right edge
	}
	for x := 0; x < len(zoomedMap); x++ {
		flood(x, 0)                // Top edge
		flood(x, len(zoomedMap)-1) // Bottom edge
	}
}

func createIsolatedMap(node *Node, nodeMap [][]*Node) [][]*Node {
	isolatedMap := make([][]*Node, len(nodeMap))
	for y, row := range nodeMap {
		isolatedMap[y] = make([]*Node, len(nodeMap[y]))
		for x := range row {
			isolatedMap[y][x] = NewNode(x, y, possiblePipes['.'])
		}
	}

	for _, from := range node.pipe.directions {
		nextX, nextY := node.x, node.y
		for {
			isolatedMap[nextY][nextX] = nodeMap[nextY][nextX]
			if nodeMap[nextY][nextX].pipe.char == 'S' {
				break
			}

			nextDir := nodeMap[nextY][nextX].pipe.output(from)
			switch nextDir {
			case NORTH:
				nextY--
				from = SOUTH
			case SOUTH:
				nextY++
				from = NORTH
			case EAST:
				nextX++
				from = WEST
			case WEST:
				nextX--
				from = EAST
			}
		}
	}

	return isolatedMap
}

// The 'zoomed' map converts every element in a traditional 'nodeMap' into a 3x3 grid that we
// can traverse.
func createZoomedMap(nodeMap [][]*Node) [][]bool {
	zoomedMap := make([][]bool, len(nodeMap)*3)
	for y, row := range nodeMap {
		zoomedMap[(y * 3)] = make([]bool, len(nodeMap[y])*3)
		zoomedMap[(y*3)+1] = make([]bool, len(nodeMap[y])*3)
		zoomedMap[(y*3)+2] = make([]bool, len(nodeMap[y])*3)
		for x, node := range row {
			// Generate the expanded grid.
			if len(node.pipe.directions) == 0 {
				continue
			}
			zoomedMap[(y*3)+1][(x*3)+1] = true // Add center point
			for _, dir := range node.pipe.directions {
				switch dir {
				case NORTH:
					zoomedMap[(y * 3)][(x*3)+1] = true
				case SOUTH:
					zoomedMap[(y*3)+2][(x*3)+1] = true
				case EAST:
					zoomedMap[(y*3)+1][(x*3)+2] = true
				case WEST:
					zoomedMap[(y*3)+1][(x * 3)] = true
				}
			}
		}
	}
	return zoomedMap
}

// Squashes are '3x zoomed' map into a 1x zoom map - we intentionally drop detail here to make it
// easier to count nodes later.
func squashZoomedMap(zoomedMap [][]bool) [][]bool {
	squashedMap := make([][]bool, len(zoomedMap)/3)
	for zy, row := range zoomedMap {
		if squashedMap[zy/3] == nil {
			squashedMap[zy/3] = make([]bool, len(zoomedMap[zy])/3)
		}
		for zx, solid := range row {
			if solid {
				squashedMap[zy/3][zx/3] = true
			}
		}
	}
	return squashedMap
}
