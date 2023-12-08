package main

import (
	"bufio"
	"fmt"
	"os"
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

	solutionPart1, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)
}

func solve(lines []string) (int, error) {
	path := lines[0]

	tree := make(map[string][]string)
	for _, line := range lines[2:] {
		data := strings.Split(line, " = ")
		tree[data[0]] = strings.Split(data[1][1:len(data[1])-1], ", ")
	}

	for node, paths := range tree {
		fmt.Printf("%q: %v\n", node, paths)
	}

	var steps int
	next := "AAA"
	for {
		if next == "ZZZ" {
			break
		}

		direction := path[steps%len(path)]
		if direction == 'L' {
			next = tree[next][0]
		} else {
			next = tree[next][1]
		}
		steps++
	}
	return steps, nil
}
