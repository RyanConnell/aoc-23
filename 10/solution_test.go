package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay10(t *testing.T) {
	tester.TimeAndCheck(10, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample-1)",
			File:        "sample1.txt",
			Solution:    solve,
			Expected:    4,
		},
		{
			Description: "Part 1 (sample-2)",
			File:        "sample2.txt",
			Solution:    solve,
			Expected:    8,
		},
		{
			Description: "Part 1 (sample-3)",
			File:        "sample3.txt",
			Solution:    solve,
			Expected:    9,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    6778,
		},
		{
			Description: "Part 2 (sample-1)",
			File:        "p2-sample1.txt",
			Solution:    solvePart2,
			Expected:    4,
		},
		{
			Description: "Part 2 (sample-2)",
			File:        "p2-sample2.txt",
			Solution:    solvePart2,
			Expected:    4,
		},
		{
			Description: "Part 2 (sample-3)",
			File:        "p2-sample3.txt",
			Solution:    solvePart2,
			Expected:    8,
		},
		{
			Description: "Part 2 (sample-4)",
			File:        "p2-sample4.txt",
			Solution:    solvePart2,
			Expected:    10,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    433,
		},
	})
}
