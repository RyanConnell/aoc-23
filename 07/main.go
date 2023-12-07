package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FIVE_KIND = iota
	FOUR_KIND
	FULL_HOUSE
	THREE_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

var cardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

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

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) handType() int {
	countPerCard := make(map[rune]int)
	for _, card := range h.cards {
		countPerCard[card]++
	}
	countPerSize := make(map[int]int)
	for _, count := range countPerCard {
		countPerSize[count]++
	}

	if countPerSize[5] == 1 {
		return FIVE_KIND
	}
	if countPerSize[4] == 1 {
		return FOUR_KIND
	}
	if countPerSize[3] == 1 && countPerSize[2] == 1 {
		return FULL_HOUSE
	}
	if countPerSize[3] == 1 {
		return THREE_KIND
	}
	if countPerSize[2] == 2 {
		return TWO_PAIR
	}
	if countPerSize[2] == 1 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func (h Hand) beats(h2 Hand) bool {
	h1Type, h2Type := h.handType(), h2.handType()
	if h1Type == h2Type {
		for i := 0; i < len(h.cards); i++ {
			if h.cards[i] == h2.cards[i] {
				continue
			}
			return cardValues[rune(h.cards[i])] > cardValues[rune(h2.cards[i])]
		}
	}
	return h1Type < h2Type
}

func solve(lines []string) (int, error) {
	var hands []Hand
	for _, line := range lines {
		data := strings.Split(line, " ")
		bid, err := strconv.Atoi(data[1])
		if err != nil {
			return 0, err
		}
		hands = append(hands, Hand{cards: data[0], bid: bid})
	}

	fmt.Printf("Unsorted:\n")
	for _, hand := range hands {
		fmt.Printf("\tHand: %+v is of type %d\n", hand, hand.handType())
	}

	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].beats(hands[j])
	})

	fmt.Printf("Sorted:\n")
	var result int
	for i, hand := range hands {
		fmt.Printf("\tHand: %+v is of type %d\n", hand, hand.handType())
		result += hand.bid * (i + 1)
	}

	return result, nil
}
