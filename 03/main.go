package main

import (
	"fmt"
	"strconv"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input.txt")

	sum, gearSum, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Sum: %d; GearSum: %d\n", sum, gearSum)

}

func solve(lines []string) (int, int, error) {
	var sum int
	adjacencies := make(map[string][]int)

	// Find and replace all numbers
	parseAndSum := func(row, startCol, endCol int) error {
		num, err := strconv.Atoi(string(lines[row][startCol : endCol+1]))
		if err != nil {
			return err
		}

		// Check around the number to see if there are any symbols
		for y := row - 1; y <= row+1; y++ {
			if y < 0 || y >= len(lines)-1 {
				continue
			}
			for x := startCol - 1; x <= endCol+1; x++ {
				if x < 0 || x >= len(lines[y]) {
					continue
				}
				if lines[y][x] == '.' || (lines[y][x] >= '0' && lines[y][x] <= '9') {
					continue
				}

				if lines[y][x] == '*' {
					key := fmt.Sprintf("%d/%d", x, y)
					if _, ok := adjacencies[key]; !ok {
						adjacencies[key] = []int{num}
					} else {
						adjacencies[key] = append(adjacencies[key], num)
					}
				}

				sum += num
			}
		}

		return nil
	}

	for y := 0; y < len(lines); y++ {
		numStart := -1
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] < '0' || lines[y][x] > '9' {
				if numStart != -1 {
					if err := parseAndSum(y, numStart, x-1); err != nil {
						return -1, -1, err
					}
					numStart = -1
				}
				continue
			}
			if numStart == -1 {
				numStart = x
			}
		}

		if numStart != -1 {
			if err := parseAndSum(y, numStart, len(lines[y])-1); err != nil {
				return -1, -1, err
			}
		}
	}

	// Sum up all gear adjacencies.
	var gearSum int
	for _, vals := range adjacencies {
		if len(vals) != 2 {
			continue
		}
		gearSum += vals[0] * vals[1]
	}

	return sum, gearSum, nil
}
