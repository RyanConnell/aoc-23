package main

import (
	"fmt"
	"strconv"
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
	var sum int
	for _, line := range strings.Split(lines[0], ",") {
		sum += hash(line)
	}
	return sum, nil
}

func hash(line string) int {
	var sum int
	for _, c := range line {
		sum += int(c)
		sum *= 17
		sum %= 256
	}
	return sum
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	hashMap := make([][]string, 256)
	for _, line := range strings.Split(lines[0], ",") {
		var sum, operatorIdx int
		for i, c := range line {
			if c == '=' || c == '-' {
				operatorIdx = i
				break
			}
			sum += int(c)
			sum *= 17
			sum %= 256
		}

		operator := line[operatorIdx]
		line = strings.Trim(fmt.Sprintf("%s %s", line[:operatorIdx], line[operatorIdx+1:]), " ")

		// Remove
		if operator == '-' {
			for i, storedLine := range hashMap[sum] {
				stored, _, _ := strings.Cut(storedLine, " ")
				if line == stored {
					hashMap[sum] = append(hashMap[sum][:i], hashMap[sum][i+1:]...)
					break
				}
			}
		}

		// Append
		if operator == '=' {
			var added bool
			for i, storedLine := range hashMap[sum] {
				label, _, _ := strings.Cut(storedLine, " ")
				if label == line[:operatorIdx] {
					hashMap[sum][i] = line
					added = true
					break
				}
			}
			if !added {
				hashMap[sum] = append(hashMap[sum], line)
			}
		}
	}

	var sum int
	for i, entries := range hashMap {
		for j, entry := range entries {
			_, focalLengthStr, _ := strings.Cut(entry, " ")
			focalLength, err := strconv.Atoi(focalLengthStr)
			if err != nil {
				return 0, err
			}
			sum += (i + 1) * (j + 1) * focalLength
		}
	}
	return sum, nil
}
