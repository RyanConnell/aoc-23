package main

import (
	"fmt"
	"strconv"
	"strings"

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

func solve(lines []string, multiplier bool) (int, error) {
	var sum int
	for _, line := range lines {
		data := strings.Split(line, " ")
		groups, err := asInts(strings.Split(data[1], ","))
		if err != nil {
			return 0, err
		}

		groupStr := data[0]
		groupCounts := append([]int{}, groups...)
		if multiplier {
			// Part 2 requires that we join data by 5x with ? as a separator
			for i := 0; i < 4; i++ {
				groupStr += fmt.Sprintf("?%s", data[0])
				groupCounts = append(groupCounts, groups...)
			}
		} else {
			groupStr = data[0]
			groupCounts = groups
		}

		count := countArangements(groupStr, groupCounts)
		sum += count
	}
	return sum, nil
}

func unique[T any](ints []T) []T {
	set := make(map[any]struct{})
	var list []T
	for _, i := range ints {
		if _, ok := set[i]; ok {
			continue
		}
		set[i] = struct{}{}
		list = append(list, i)
	}
	return list
}

func countArangements(group string, counts []int) int {
	uniqueCounts := unique(counts)

	var lastSpring int
	for i, char := range group {
		if char == '#' && i > lastSpring {
			lastSpring = i
		}
	}

	cache := make(map[string][]int)
	var explore func(int) []int
	explore = func(idx int) []int {
		if idx >= len(group) {
			return nil
		}
		if idx < 0 {
			return cache[group]
		}
		if val, ok := cache[group[idx:]]; ok {
			return val
		}

		matches := make([]int, len(counts)+1)
		if group[idx] != '.' {
			for _, count := range uniqueCounts {
				if idx+count > len(group) {
					continue
				}
				if strings.Contains(group[idx:idx+count], ".") {
					continue
				}

				// If the character _after_ the count is a '#' then count is not valid.
				if idx+count < len(group) && group[idx+count] == '#' {
					continue
				}

				if count == counts[len(counts)-1] {
					if idx+count >= len(group) || lastSpring < idx+count+1 {
						matches[1]++
					}
				}
				for subMatch, matchCount := range explore(idx + count + 1) {
					if matchCount == 0 {
						continue
					}
					if subMatch < len(counts) && count == counts[len(counts)-1-subMatch] {
						matches[subMatch+1] += matchCount
					}
				}
			}
		}

		if group[idx] != '#' {
			for i, match := range explore(idx + 1) {
				matches[i] += match
			}
		}

		cache[group[idx:]] = append([]int{}, matches...)
		return explore(idx - 1)
	}

	return explore(len(group) - 1)[len(counts)]
}

func asInts(input []string) ([]int, error) {
	var results []int
	for _, in := range input {
		if in == "" {
			continue
		}
		result, err := strconv.Atoi(in)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
