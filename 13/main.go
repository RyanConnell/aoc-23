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

type Pattern struct {
	lines           []string
	hasHiddenMirror bool
}

func (p *Pattern) lineOfSymmetry() (int, int) {
	horizontal := findLineOfSymmetry(p.lines, p.hasHiddenMirror)
	rotated := rotate(p.lines)
	vertical := findLineOfSymmetry(rotated, p.hasHiddenMirror)
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

func findLineOfSymmetry(lines []string, allowOffByOne bool) int {
	// Vertical symmetry will always have a situation where the first line or the last line
	// are at the edge of the symmetry.
	positions := [][]int{
		{}, // Represents the indexes that match line[0]
		{}, // Represents the indexes that match line[-1]
	}

	var maxDiff int
	if allowOffByOne {
		maxDiff = 1
	}

	for i, line := range lines {
		if i != 0 {
			if lineDifferences(line, lines[0]) <= maxDiff {
				positions[0] = append(positions[0], i)
			}
		}
		if i != len(lines)-1 {
			if lineDifferences(line, lines[len(lines)-1]) <= maxDiff {
				positions[1] = append(positions[1], i)
			}
		}
	}

	// For each potential match check if the symmetry works out.
	for _, pos := range positions[0] {
		if checkSymmetry(lines, pos/2, allowOffByOne) {
			return pos / 2
		}
	}
	for _, pos := range positions[1] {
		midpoint := pos + ((len(lines) - 1 - pos) / 2)
		if checkSymmetry(lines, midpoint, allowOffByOne) {
			return midpoint
		}
	}
	return -1
}

// Midpoint always refers to the N where we suspect (lines[N] == lines[N+1])
func checkSymmetry(lines []string, midpoint int, allowOffByOne bool) bool {
	var maxDiff int
	if allowOffByOne {
		maxDiff = 1
	}
	for i := 0; i <= midpoint; i++ {
		if midpoint-i < 0 || midpoint+1+i >= len(lines) {
			break
		}
		maxDiff -= lineDifferences(lines[midpoint-i], lines[midpoint+1+i])
		if maxDiff < 0 {
			return false
		}
	}
	return maxDiff == 0
}

func parsePatterns(lines []string, hasHiddenMirror bool) []*Pattern {
	var lastPatternBreak int
	var patterns []*Pattern
	for i, line := range lines {
		if line != "" {
			continue
		}
		patterns = append(patterns, &Pattern{lines[lastPatternBreak:i], hasHiddenMirror})
		lastPatternBreak = i + 1
	}
	return append(patterns, &Pattern{lines[lastPatternBreak:], hasHiddenMirror})
}

func solve(lines []string, hasSmudge bool) (int, error) {
	patterns := parsePatterns(lines, hasSmudge)
	var sum int
	for _, pattern := range patterns {
		verticalPos, horizontalPos := pattern.lineOfSymmetry()
		if verticalPos != -1 {
			sum += verticalPos + 1
		} else if horizontalPos != -1 {
			sum += 100 * (horizontalPos + 1)
		}
	}
	return sum, nil
}

func lineDifferences(line1, line2 string) int {
	var count int
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			count++
		}
	}
	return count
}
