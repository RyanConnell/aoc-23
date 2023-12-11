package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay11(t *testing.T) {
	tester.TimeAndCheck(11, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    374,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    10077850,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    82000210,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    504715068438,
		},
	})
}
