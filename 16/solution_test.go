package main

import (
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay16(t *testing.T) {
	tester.TimeAndCheck(16, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    46,
		},
	})
}
