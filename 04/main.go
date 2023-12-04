package main

import (
	"bufio"
	"fmt"
	"math"
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

	sum, totalCards, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Sum: %d; Total Cards: %d\n", sum, totalCards)

}

func solve(lines []string) (int, int, error) {
	var sum int
	winnersPerCard := make(map[int]int)

	for i, line := range lines {
		round := strings.Split(line, ": ")[1]
		data := strings.Split(round, " | ")

		winners := make(map[string]struct{})
		for _, num := range strings.Split(data[0], " ") {
			if num != "" {
				winners[num] = struct{}{}
			}
		}

		var count int
		for _, num := range strings.Split(data[1], " ") {
			if _, ok := winners[num]; ok {
				count++
			}
		}
		winnersPerCard[i] = count

		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
		}
	}

	var cardTotal int
	cache := make([]int, len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		cache[i] = 1
		for j := 1; j <= winnersPerCard[i]; j++ {
			if i+j < len(cache) {
				cache[i] += cache[i+j]
			}
		}
		cardTotal += cache[i]
	}

	return sum, cardTotal, nil
}
