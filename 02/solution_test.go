package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay02(t *testing.T) {
	tester.TimeAndCheck(2, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { r, _ := solve(lines); return r, nil },
			Expected:    8,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { r, _ := solve(lines); return r, nil },
			Expected:    2593,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { _, r := solve(lines); return r, nil },
			Expected:    2286,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { _, r := solve(lines); return r, nil },
			Expected:    54699,
		},
	})
}
