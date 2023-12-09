package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-23/pkg/parser"
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
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, false)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
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
	return checkRank(countPerCard)
}

func (h *Hand) handTypeWithWildcard() int {
	countPerCard := make(map[rune]int)
	for _, card := range h.cards {
		countPerCard[card]++
	}
	jokerCount := countPerCard['J']
	delete(countPerCard, 'J')

	type CardInfo struct {
		card  rune
		count int
	}
	var cardInfos []CardInfo
	for card, count := range countPerCard {
		cardInfos = append(cardInfos, CardInfo{card, count})
	}
	sort.Slice(cardInfos, func(i, j int) bool {
		return cardInfos[i].count > cardInfos[j].count
	})

	if jokerCount == 5 {
		return FIVE_KIND
	}
	countPerCard[cardInfos[0].card] += jokerCount
	return checkRank(countPerCard)
}

func (h Hand) beats(h2 Hand, useWildcards bool) bool {
	var h1Type, h2Type int
	if useWildcards {
		cardValues['J'] = 0
		h1Type, h2Type = h.handTypeWithWildcard(), h2.handTypeWithWildcard()
	} else {
		cardValues['J'] = 11
		h1Type, h2Type = h.handType(), h2.handType()
	}

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

func checkRank(countPerCard map[rune]int) int {
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

func solve(lines []string, useWildcards bool) (int, error) {
	var hands []Hand
	for _, line := range lines {
		data := strings.Split(line, " ")
		bid, err := strconv.Atoi(data[1])
		if err != nil {
			return 0, err
		}
		hands = append(hands, Hand{cards: data[0], bid: bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].beats(hands[j], useWildcards)
	})

	var result int
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}

	return result, nil
}
