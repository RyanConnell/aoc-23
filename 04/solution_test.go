package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

// Note: Part 1 and 2 are running the same code here, so timings wont really differ.

func TestDay04(t *testing.T) {
	tester.TimeAndCheck(4, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (int, error) {
				r, _, err := solve(lines)
				return r, err
			},
			Expected: 13,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution: func(lines []string) (int, error) {
				r, _, err := solve(lines)
				return r, err
			},
			Expected: 24848,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution: func(lines []string) (int, error) {
				_, r, err := solve(lines)
				return r, err
			},
			Expected: 30,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution: func(lines []string) (int, error) {
				_, r, err := solve(lines)
				return r, err
			},
			Expected: 7258152,
		},
	})
}
