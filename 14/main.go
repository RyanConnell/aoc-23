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

func solve(lines []string) (int, error) {
	layout := make([][]rune, len(lines))
	for i, line := range lines {
		layout[i] = []rune(line)
	}

	tiltHorizontal(layout, -1) // North

	var sum int
	for y, line := range layout {
		for _, char := range line {
			if char == 'O' {
				sum += len(layout) - y
			}
		}
	}

	return sum, nil
}

func solvePart2(lines []string) (int, error) {
	layout := make([][]rune, len(lines))
	lastLayout := make([][]rune, len(lines))
	for i, line := range lines {
		layout[i] = []rune(line)
		lastLayout[i] = []rune(line)
	}

	seen := map[string]struct{}{}
	maxLen := 1000000000
	loopStart, loopEnd := -1, -1
	for i := 0; i < maxLen; i++ {
		// Tilt in each direction
		tiltHorizontal(layout, -1) // North
		tiltVertical(layout, -1)   // West
		tiltHorizontal(layout, 1)  // South
		tiltVertical(layout, 1)    // East

		key := layoutKey(layout)
		if _, ok := seen[key]; ok {
			if loopStart == -1 {
				loopStart = i
			} else if loopEnd == -1 {
				loopEnd = i

				loopSize := loopEnd - loopStart
				numDivisions := (maxLen - loopStart) / loopSize
				endpoint := maxLen - (loopSize * numDivisions) + loopSize
				maxLen = endpoint
			}
			seen = map[string]struct{}{}
		}
		seen[key] = struct{}{}
	}

	var sum int
	for y, line := range layout {
		for _, char := range line {
			if char == 'O' {
				sum += len(layout) - y
			}
		}
	}
	return sum, nil
}

func layoutKey(layout [][]rune) string {
	var key string
	for i := 0; i < len(layout); i++ {
		key += string(layout[i])
	}
	return key
}

// direction must be set to +1 (south) or -1 (north)
func tiltVertical(layout [][]rune, direction int) {
	for _, line := range layout {
		var start int
		end := len(line) - 1
		if direction == -1 {
			start = len(line) - 1
			end = 0
		}

		var consecutive int

		for i := start; i >= 0 && i < len(line); i += direction {
			if line[i] == 'O' {
				consecutive++
				line[i] = '.'
			}

			if consecutive == 0 {
				continue
			}

			if line[i] == '#' || i == end {
				var startOffset int
				if line[i] == '#' {
					startOffset = 1
				}
				for offset := startOffset; offset < consecutive+startOffset; offset++ {
					realOffset := offset
					if direction == 1 {
						realOffset *= -1
					}
					if line[i+realOffset] == '.' {
						line[i+realOffset] = 'O'
					}
				}
				consecutive = 0
			}

		}
	}
}

// direction must be set to +1 (south) or -1 (north)
func tiltHorizontal(layout [][]rune, direction int) {
	// Squash upwards
	consecutive := make([]int, len(layout[0]))

	do := func(i, end int) {
		for j, char := range layout[i] {
			if char == 'O' {
				consecutive[j]++
				layout[i][j] = '.'
			}

			if consecutive[j] == 0 {
				continue
			}

			if char == '#' || i == end {
				var startOffset int
				if char == '#' {
					startOffset = 1
				}
				for offset := startOffset; offset < consecutive[j]+startOffset; offset++ {
					realOffset := offset
					if direction == 1 {
						realOffset *= -1
					}
					if layout[i+realOffset][j] == '.' {
						layout[i+realOffset][j] = 'O'
					}
				}
				consecutive[j] = 0
			}

		}

	}
	if direction == -1 {
		for i := len(layout) - 1; i >= 0; i-- {
			do(i, 0)
		}
	} else {
		for i := 0; i < len(layout); i++ {
			do(i, len(layout)-1)
		}
	}

}
