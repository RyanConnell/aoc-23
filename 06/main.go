package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	solutionPart1, err := solve(lines, false)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1: Combinations: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2: Combinations: %d\n", solutionPart2)
}

func solve(lines []string, squashInput bool) (int, error) {
	var timeStrs, distanceStrs []string
	if squashInput {
		timeStrs = []string{strings.Replace(strings.Split(lines[0], ":")[1], " ", "", -1)}
		distanceStrs = []string{strings.Replace(strings.Split(lines[1], ":")[1], " ", "", -1)}
	} else {
		timeStrs = strings.Split(strings.Split(lines[0], ":")[1], " ")
		distanceStrs = strings.Split(strings.Split(lines[1], ":")[1], " ")
	}

	times, err := asInts(timeStrs)
	if err != nil {
		return 0, err
	}
	distances, err := asInts(distanceStrs)
	if err != nil {
		return 0, err
	}

	sum := 1
	for i := 0; i < len(times); i++ {
		val := combinations(times[i], distances[i])
		sum *= val
	}

	return sum, nil
}

func asInts(strs []string) ([]int, error) {
	var ints []int
	for _, str := range strs {
		if str == "" {
			continue
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, val)
	}
	return ints, nil
}

func combinations(raceTime, expectedDistance int) int {
	var count int
	for speed := 1; speed < raceTime; speed++ {
		if speed*(raceTime-speed) > expectedDistance {
			count++
		}
	}
	return count
}
