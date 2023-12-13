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

type Pattern struct {
	lines []string
}

func (p *Pattern) lineOfSymmetry() (int, int) {
	horizontal := findLineOfSymmetry(p.lines)
	rotated := rotate(p.lines)
	vertical := findLineOfSymmetry(rotated)
	return vertical, horizontal
}

func rotate(lines []string) []string {
	rotated := make([]string, len(lines[0]))
	for _, line := range lines {
		for j, char := range line {
			rotated[j] += string(char)
		}
	}
	return rotated
}

func findLineOfSymmetry(lines []string) int {
	// Vertical symmetry will always have a situation where the first line or the last line
	// are at the edge of the symmetry.
	positions := [][]int{
		{}, // Represents the indexes that match line[0]
		{}, // Represents the indexes that match line[-1]
	}
	for i, line := range lines {
		if i != 0 && line == lines[0] {
			positions[0] = append(positions[0], i)
		}
		if i != len(lines)-1 && line == lines[len(lines)-1] {
			positions[1] = append(positions[1], i)
		}
	}

	// For each potential match check if the symmetry works out.
	for _, pos := range positions[0] {
		if checkSymmetry(lines, pos/2) {
			return pos / 2
		}
	}
	for _, pos := range positions[1] {
		midpoint := pos + ((len(lines) - 1 - pos) / 2)
		if checkSymmetry(lines, midpoint) {
			return midpoint
		}
	}
	return -1
}

// Midpoint always refers to the N where we suspect (lines[N] == lines[N+1])
func checkSymmetry(lines []string, midpoint int) bool {
	fmt.Printf("CheckSymmetry: %d: ", midpoint)
	for i := 0; i < midpoint; i++ {
		if midpoint-i < 0 || midpoint+1+i >= len(lines) {
			break
		}
		if lines[midpoint-i] != lines[midpoint+1+i] {
			fmt.Printf("false\n")
			return false
		}
	}
	fmt.Printf("true\n")
	return true
}

/// Part 1 \\\

func parsePatterns(lines []string) []*Pattern {
	var lastPatternBreak int
	var patterns []*Pattern
	for i, line := range lines {
		if line != "" {
			continue
		}
		patterns = append(patterns, &Pattern{lines[lastPatternBreak:i]})
		lastPatternBreak = i + 1
	}
	return append(patterns, &Pattern{lines[lastPatternBreak:]})
}

func solve(lines []string) (int, error) {
	patterns := parsePatterns(lines)
	var sum int
	for i, pattern := range patterns {
		fmt.Printf("Pattern Length: %d\n", len(pattern.lines))
		verticalPos, horizontalPos := pattern.lineOfSymmetry()
		if verticalPos != -1 {
			sum += verticalPos + 1
			fmt.Printf("Pattern %d: vertical @ %d\n", i, verticalPos+1)
		} else if horizontalPos != -1 {
			sum += 100 * (horizontalPos + 1)
			fmt.Printf("Pattern %d: horizontal @ %d\n", i, horizontalPos+1)
		} else {
			fmt.Printf("WARNING: Pattern %d had no horizontal or vertical symmetry\n", i)
		}
	}
	return sum, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
