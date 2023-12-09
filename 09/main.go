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

type Sequence struct {
	values []int
}

func (s *Sequence) extrapolate() [][]int {
	results := [][]int{s.values}
	for {
		lastRow := len(results) - 1

		allZeros := true
		result := make([]int, len(results[lastRow])-1)
		for i := 0; i < len(results[lastRow])-1; i++ {
			result[i] = results[lastRow][i+1] - results[lastRow][i]
			if result[i] != 0 {
				allZeros = false
			}
		}
		results = append(results, result)
		if allZeros {
			break
		}
	}
	return results
}

func (s *Sequence) prev() int {
	increases := s.extrapolate()
	var offset int
	for i := len(increases) - 2; i > 0; i-- {
		offset = increases[i][0] - offset
	}
	return s.values[0] - offset
}

func (s *Sequence) next() int {
	increases := s.extrapolate()
	var offset int
	for i := len(increases) - 2; i > 0; i-- {
		offset = increases[i][len(increases[i])-1] + offset
	}
	return s.values[len(s.values)-1] + offset

}

func asInts(values []string) ([]int, error) {
	var results []int
	for _, val := range values {
		result, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func solve(lines []string, includePrevious bool) (int, error) {
	sequences := make([]*Sequence, len(lines))
	for i, line := range lines {
		values, err := asInts(strings.Split(line, " "))
		if err != nil {
			return 0, err
		}
		sequences[i] = &Sequence{values}
	}

	var sum int
	for _, seq := range sequences {
		var next int
		if includePrevious {
			next = seq.prev()
		} else {
			next = seq.next()
		}
		sum += next
	}

	return sum, nil
}
