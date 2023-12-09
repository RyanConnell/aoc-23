package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

// Note: Part 1 and 2 are running the same code here, so timings wont really differ.

func TestDay03(t *testing.T) {
	tester.TimeAndCheck(3, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (int, error) {
				r, _, err := solve(lines)
				return r, err
			},
			Expected: 4361,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution: func(lines []string) (int, error) {
				r, _, err := solve(lines)
				return r, err
			},
			Expected: 528819,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (int, error) {
				_, r, err := solve(lines)
				return r, err
			},
			Expected: 467835,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution: func(lines []string) (int, error) {
				_, r, err := solve(lines)
				return r, err
			},
			Expected: 80403602,
		},
	})
}
