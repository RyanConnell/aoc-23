package main

import (
	"fmt"
	"testing"

	"github.com/RyanConnell/aoc-23/pkg/tester"
)

func TestDay12(t *testing.T) {
	tester.TimeAndCheck(12, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample-1)",
			File:        "sample1.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    6,
		},
		{
			Description: "Part 1 (sample-2)",
			File:        "sample2.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    21,
		},
		{
			Description: "Part 1 (sample-rob)",
			File:        "sample-rob.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    364,
		},
		{
			Description: "Part 1 (final)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    8075,
		},
		{
			Description: "Part 2 (sample-1)",
			File:        "sample2.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    525152,
		},
	})
}

func TestCountArangements(t *testing.T) {
	testCases := []struct {
		group    string
		counts   []int
		expected int
	}{
		{"?.##.###??.????#.", []int{1, 2, 4, 1, 1}, 3},
		{"???.?..##.?.?#?.??", []int{3, 1, 2, 1, 1, 1}, 2},
		{"????.?#?#???.#.?#?.", []int{2, 4, 2, 1, 2}, 6},
		{"##?.#???.#..?.????#?", []int{3, 1, 1, 1, 1, 4}, 7},
	}

	/*
		###.#.#..#..#...#### 3,1,1,1,1,4
		###.#.#..#..#..####. 3,1,1,1,1,4
		###.#.#..#....#.#### 3,1,1,1,1,4
		###.#..#.#..#...#### 3,1,1,1,1,4
		###.#..#.#..#..####. 3,1,1,1,1,4
		###.#..#.#....#.#### 3,1,1,1,1,4

		###.#..#.#..?.????#? 3,1,1,1,1,4
		##?.#???.#..?.????#? 3,1,1,1,1,4
		##?.#???.#..?.????#? 3,1,1,1,1,4
		##?.#???.#..?.????#? 3,1,1,1,1,4
	*/

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := countArangements(tc.group, tc.counts)
			if result != tc.expected {
				t.Fatalf("Unexpected result: want %d got %d\n", tc.expected, result)
			}
		})
	}
}
