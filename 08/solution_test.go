package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay08(t *testing.T) {
	tester.TimeAndCheck(8, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample-1.txt",
			Solution:    solve,
			Expected:    6,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    24253,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample-2.txt",
			Solution:    solvePart2,
			Expected:    6,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    12357789728873,
		},
	})
}
