package main

import (
	"fmt"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, false)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

type Mirror struct {
	char               rune
	directionsForInput map[int][]int
}

func NewMirror(char rune, directions map[int][]int) Mirror {
	return Mirror{char: char, directionsForInput: directions}
}

var possibleMirrors = map[rune]Mirror{
	'|': NewMirror('|', map[int][]int{
		NORTH: {SOUTH},
		SOUTH: {NORTH},
		EAST:  {NORTH, SOUTH},
		WEST:  {NORTH, SOUTH},
	}),

	'-': NewMirror('-', map[int][]int{
		NORTH: {EAST, WEST},
		SOUTH: {EAST, WEST},
		EAST:  {WEST},
		WEST:  {EAST},
	}),

	'\\': NewMirror('\\', map[int][]int{
		NORTH: {EAST},
		SOUTH: {WEST},
		EAST:  {NORTH},
		WEST:  {SOUTH},
	}),

	'/': NewMirror('/', map[int][]int{
		NORTH: {WEST},
		SOUTH: {EAST},
		WEST:  {NORTH},
		EAST:  {SOUTH},
	}),

	'.': NewMirror('.', map[int][]int{
		NORTH: {SOUTH},
		SOUTH: {NORTH},
		EAST:  {WEST},
		WEST:  {EAST},
	}),
}

type Node struct {
	mirror      Mirror
	x, y        int
	visitedFrom map[int]struct{}
}

func NewNode(x, y int, mirror Mirror) *Node {
	return &Node{x: x, y: y, mirror: mirror, visitedFrom: make(map[int]struct{})}
}
func solve(lines []string, exploreBoundaries bool) (int, error) {
	nodeMap := make([][]*Node, len(lines))
	energised := make([][]bool, len(lines))
	for y, line := range lines {
		nodeMap[y] = make([]*Node, len(line))
		for x, char := range line {
			nodeMap[y][x] = NewNode(x, y, possibleMirrors[char])
		}
		energised[y] = make([]bool, len(line))
	}

	var explore func(int, int, int)
	explore = func(x, y int, direction int) {
		if x < 0 || y < 0 || x >= len(nodeMap[0]) || y >= len(nodeMap) {
			return
		}
		node := nodeMap[y][x]
		energised[y][x] = true

		var nextInputDir int
		switch direction {
		case NORTH:
			nextInputDir = SOUTH
		case SOUTH:
			nextInputDir = NORTH
		case EAST:
			nextInputDir = WEST
		case WEST:
			nextInputDir = EAST
		}
		if _, ok := node.visitedFrom[direction]; ok {
			return
		}
		node.visitedFrom[direction] = struct{}{}

		nextDirections, ok := node.mirror.directionsForInput[nextInputDir]
		if !ok {
			return
		}

		for _, dir := range nextDirections {
			switch dir {
			case NORTH:
				explore(x, y-1, NORTH)
			case SOUTH:
				explore(x, y+1, SOUTH)
			case EAST:
				explore(x+1, y, EAST)
			case WEST:
				explore(x-1, y, WEST)
			}
		}
	}

	var maxEnergised int
	check := func(x, y, dir int) {
		explore(x, y, dir)

		var count int
		for y, row := range energised {
			for x := range row {
				if energised[y][x] {
					count++
					energised[y][x] = false
				}
				nodeMap[y][x].visitedFrom = make(map[int]struct{})
			}
		}
		if count > maxEnergised {
			maxEnergised = count
		}

	}

	if exploreBoundaries {
		for y := 0; y < len(lines); y++ {
			check(0, y, EAST)
			check(len(lines[y])-1, y, WEST)
		}

		for x := 0; x < len(lines[0]); x++ {
			check(x, 0, SOUTH)
			check(x, len(lines), NORTH)
		}
	} else {
		check(0, 0, EAST)
	}
	return maxEnergised, nil
}
