package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay15(t *testing.T) {
	tester.TimeAndCheck(15, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    1320,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    497373,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    145,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    259356,
		},
	})
}
