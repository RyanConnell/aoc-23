package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay09(t *testing.T) {
	tester.TimeAndCheck(9, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    114,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    1993300041,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    2,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    1038,
		},
	})
}
