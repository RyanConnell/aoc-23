package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

var validNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	lines := parser.MustReadFile("input/input.txt")
	result, _ := solve(lines)
	fmt.Printf("Sum: %d\n", result)
}

func solve(lines []string) (int, error) {
	var sum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		val, err := asCalibrationValue(line)
		if err != nil {
			continue
		}
		sum += val
	}
	return sum, nil
}

func hasNumStrPrefix(str string) (int, bool) {
	for numStr, val := range validNumbers {
		if strings.HasPrefix(str, numStr) {
			return val, true
		}
	}
	return 0, false
}

func asCalibrationValue(str string) (int, error) {
	check := func(idx int) (int, error) {
		if str[idx] >= '0' && str[idx] <= '9' {
			first, err := strconv.Atoi(string(str[idx]))
			if err != nil {
				return -1, err
			}
			return first, nil
		}

		val, ok := hasNumStrPrefix(str[idx:])
		if !ok {
			return -1, nil
		}
		return val, nil
	}

	var err error
	first, last := -1, -1
	for i := 0; i < len(str); i++ {
		if first != -1 && last != -1 {
			break
		}
		if first == -1 {
			first, err = check(i)
			if err != nil {
				return -1, err
			}
		}
		if last == -1 {
			last, err = check(len(str) - 1 - i)
			if err != nil {
				return -1, err
			}
		}
	}

	return (first * 10) + last, nil
}
