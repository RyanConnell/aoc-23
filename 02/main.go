package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")
	sum, powerSum := solve(lines)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("PowerSum: %d\n", powerSum)
}

func solve(lines []string) (int, int) {
	validCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	var sum, powerSum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		game, err := parseLine(line)
		if err != nil {
			continue
		}

		valid := game.valid(validCubes)
		if valid {
			sum += game.id
		}

		power := game.power()
		powerSum += power
	}

	return sum, powerSum
}

type Game struct {
	id     int
	rounds []map[string]int
}

func (g *Game) valid(cubeCounts map[string]int) bool {
	for _, round := range g.rounds {
		for color, count := range round {
			maxCount, ok := cubeCounts[color]
			if !ok || maxCount < count {
				return false
			}
		}
	}
	return true
}

func (g *Game) power() int {
	maxCubes := make(map[string]int)
	for _, round := range g.rounds {
		for color, count := range round {
			maxCount, ok := maxCubes[color]
			if !ok || maxCount < count {
				maxCubes[color] = count
			}
		}
	}

	power := 1
	for _, count := range maxCubes {
		power *= count
	}
	return power
}

func parseLine(str string) (*Game, error) {
	segments := strings.Split(str, ": ")
	gameID, err := strconv.Atoi(strings.Split(segments[0], " ")[1])
	if err != nil {
		return nil, err
	}

	var rounds []map[string]int
	for id, line := range strings.Split(segments[1], "; ") {
		colors := strings.Split(line, ", ")
		rounds = append(rounds, make(map[string]int))
		for _, color := range colors {
			data := strings.Split(color, " ")
			count, err := strconv.Atoi(data[0])
			if err != nil {
				return nil, err
			}
			rounds[id][data[1]] = count
		}
	}
	return &Game{id: gameID, rounds: rounds}, nil
}
