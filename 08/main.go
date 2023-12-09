package main

import (
	"fmt"
	"strings"

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

func solve(lines []string) (int, error) {
	path := lines[0]

	tree := make(map[string][]string)
	for _, line := range lines[2:] {
		data := strings.Split(line, " = ")
		tree[data[0]] = strings.Split(data[1][1:len(data[1])-1], ", ")
	}

	var steps int
	next := "AAA"
	for {
		if next == "ZZZ" {
			break
		}

		direction := path[steps%len(path)]
		if direction == 'L' {
			next = tree[next][0]
		} else {
			next = tree[next][1]
		}
		steps++
	}
	return steps, nil
}

/// Part 2 \\\

type Traveller struct {
	start      string
	current    string
	zLocations []int
}

func (t *Traveller) zDeltas() []int {
	deltas := make([]int, len(t.zLocations))
	var last int
	for i, zLoc := range t.zLocations {
		deltas[i] = zLoc - last
		last = zLoc
	}
	return deltas
}

func solvePart2(lines []string) (int, error) {
	path := lines[0]

	var travellers []*Traveller
	tree := make(map[string][]string)
	pathsToNode := make(map[string]int)
	for _, line := range lines[2:] {
		data := strings.Split(line, " = ")
		tree[data[0]] = strings.Split(data[1][1:len(data[1])-1], ", ")

		if data[0][2] == 'A' {
			travellers = append(travellers, &Traveller{data[0], data[0], nil})
		}

		pathsToNode[tree[data[0]][0]] += 1
		pathsToNode[tree[data[0]][1]] += 1
	}

	var steps int
	for {
		allSatisfied := true
		for _, traveller := range travellers {
			if traveller.current[2] == 'Z' {
				traveller.zLocations = append(traveller.zLocations, steps)
			}
			if traveller.current[2] != 'Z' {
				allSatisfied = allSatisfied && len(traveller.zLocations) > 5
				if allSatisfied {
					continue
				}
			}
		}
		if allSatisfied {
			break
		}

		direction := path[steps%len(path)]
		for _, traveller := range travellers {
			if direction == 'L' {
				traveller.current = tree[traveller.current][0]
			} else {
				traveller.current = tree[traveller.current][1]
			}
		}
		steps++
	}

	stepsPerLoop := make([]int, len(travellers))
	for i, traveller := range travellers {
		stepsPerLoop[i] = traveller.zLocations[1] - traveller.zLocations[0]
	}

	result := 1
	for _, steps := range stepsPerLoop {
		result = lcm(result, steps)
	}

	return result, nil
}

// Some euclidean algorithm for doing this - far too much maths for my liking...
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
