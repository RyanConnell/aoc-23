package main

import (
	"fmt"

	"github.com/RyanConnell/aoc-23/pkg/parser"
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

type Galaxy struct {
	x, y int
}

func (g *Galaxy) distanceTo(target *Galaxy) int {
	var distance int
	if g.x > target.x {
		distance += g.x - target.x
	} else {
		distance += target.x - g.x
	}

	if g.y > target.y {
		distance += g.y - target.y
	} else {
		distance += target.y - g.y
	}

	return distance
}

func solve(lines []string, massiveExpansion bool) (int, error) {
	// Find the paths between galaxies
	galaxies := findGalaxies(lines, massiveExpansion)
	var sum int
	for i, source := range galaxies {
		for _, target := range galaxies[i:] {
			sum += source.distanceTo(target)
		}
	}
	return sum, nil
}

func findGalaxies(galaxyMap []string, massiveExpansion bool) []*Galaxy {
	yHasGalaxy := make([]bool, len(galaxyMap))
	xHasGalaxy := make([]bool, len(galaxyMap[0]))
	var galaxies []*Galaxy
	for y, row := range galaxyMap {
		var yFound bool
		for x, char := range row {
			if char == '#' {
				galaxies = append(galaxies, &Galaxy{x, y})
				xHasGalaxy[x] = xHasGalaxy[x] || true
				yFound = true
			}
		}
		yHasGalaxy[y] = yFound
	}

	offsets := func(hasGalaxy []bool) []int {
		result := make([]int, len(galaxyMap))
		for y := range yHasGalaxy {
			if y != 0 {
				result[y] = result[y-1]
			}
			if !hasGalaxy[y] {
				if !massiveExpansion {
					result[y]++
				} else {
					result[y] += 999999
				}
			}
		}
		return result
	}
	xOffsets := offsets(xHasGalaxy)
	yOffsets := offsets(yHasGalaxy)

	for _, galaxy := range galaxies {
		galaxy.x += xOffsets[galaxy.x]
		galaxy.y += yOffsets[galaxy.y]
	}

	return galaxies
}
