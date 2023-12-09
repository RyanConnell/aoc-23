package tester

import (
	"fmt"
	"time"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

type TestCase[V comparable] struct {
	Description string
	File        string
	Solution    func(lines []string) (V, error)
	Expected    V
}

func TimeAndCheck[V comparable](day int, cases []TestCase[V]) {
	for _, tc := range cases {
		lines := parser.MustReadFile(fmt.Sprintf("input/%s", tc.File))

		start := time.Now()
		result, err := tc.Solution(lines)
		duration := time.Now().Sub(start)

		var message string
		if err != nil {
			message = fmt.Sprintf("unexpected failure: %v", err)
		} else if result != tc.Expected {
			message = fmt.Sprintf("incorrect result; want %v got %v", tc.Expected, result)
		} else {
			message = fmt.Sprintf("took %s", duration)
		}

		fmt.Printf("| Day %02d | %-20s| %-20s |\n", day, tc.Description, message)
	}
}
