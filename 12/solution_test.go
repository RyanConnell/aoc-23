package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay12(t *testing.T) {
	tester.TimeAndCheck(12, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    21,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    8075,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    525152,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    4232520187524,
		},
	})
}
