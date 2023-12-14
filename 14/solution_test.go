package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay14(t *testing.T) {
	tester.TimeAndCheck(14, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    136,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    105982,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    64,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    85175,
		},
	})
}
