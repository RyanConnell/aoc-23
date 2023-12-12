package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/aoc-23/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/sample2.txt")

	//	solutionPart1, err := solve(lines, false)
	//	if err != nil {
	//		fmt.Printf("Error: %v", err)
	//		return
	//	}
	//	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

func solve(lines []string, multiplier bool) (int, error) {
	var sum int
	for i, line := range lines {
		data := strings.Split(line, " ")
		groups, err := asInts(strings.Split(data[1], ","))
		if err != nil {
			return 0, err
		}

		groupStr := data[0]
		groupCounts := append([]int{}, groups...)
		if multiplier { // Part 2 requires that we join data by 5x with ? as a separator
			for i := 0; i < 4; i++ {
				groupStr += fmt.Sprintf("?%s", data[0])
				groupCounts = append(groupCounts, groups...)
			}
			//fmt.Printf("GroupStr: %q\n", groupStr)
			//fmt.Printf("GroupCounts: %v\n", groupCounts)
		} else {
			groupStr = data[0]
			groupCounts = groups
		}

		count := countArangements(groupStr, groupCounts)
		fmt.Printf("%4d/%-4d => %50s => %d ways\n", i+1, len(lines), line, count)
		sum += count
	}
	return sum, nil
}

func asInts(input []string) ([]int, error) {
	var results []int
	for _, in := range input {
		if in == "" {
			continue
		}
		result, err := strconv.Atoi(in)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func countArangements(line string, groupCounts []int) int {
	var groups []string
	for _, val := range strings.Split(line, ".") {
		if val == "" {
			continue
		}
		groups = append(groups, val)
	}

	validPos := validPositions(groups, groupCounts)
	//fmt.Printf("Valid Positions: %v\n", validPos)
	combosPerGroup := make([][][]int, len(groups))
	groupCache := make(map[string]int)
	for i, group := range groups {
		//fmt.Printf("\t\t- Computing %q\n", group)
		if cacheIdx, ok := groupCache[group]; ok {
			combosPerGroup[i] = combosPerGroup[cacheIdx]
		} else {
			groupCache[group] = i
			combosPerGroup[i] = positionCombinations(group, validPos[group], groupCounts)
		}
		//fmt.Printf("Combos per group %d: %v\n", i, combosPerGroup[i])
	}

	//fmt.Printf("\t\t- Calulating allCombinations\n")
	return allCombinations(combosPerGroup, groupCounts)
}

type Combination struct {
	values []int
	count  int
}

func NewCombo(values []int) *Combination {
	return &Combination{values: values, count: 1}
}

func (c *Combination) key() string {
	return fmt.Sprintf("%+v", c.values)
}

func product2(combos map[string]*Combination, b [][]int) map[string]*Combination {
	comboMap := make(map[string]*Combination)
	for _, c := range combos {
		for i := range b {
			vals := append([]int{}, c.values...)
			vals = append(vals, b[i]...)
			newC := NewCombo(vals)
			key := newC.key()
			if _, ok := comboMap[key]; ok {
				comboMap[key].count += c.count
			} else {
				comboMap[key] = newC
				newC.count = c.count
			}
		}
	}
	return comboMap
}

// Old version we used for part 1; Replaced by product2 above ^
func product(a [][]int, b [][]int) [][]int {
	var combos [][]int
	for i := range a {
		for j := range b {
			val := append([]int{}, a[i]...)
			val = append(val, b[j]...)
			combos = append(combos, val)
		}
	}
	return combos
}

func allCombinations(combosPerGroup [][][]int, groupCounts []int) int {
	allCombos := make(map[string]*Combination)
	for _, combo := range combosPerGroup[0] {
		if checkCombo(combo, groupCounts, true) {
			nc := NewCombo(combo)
			key := nc.key()
			if _, ok := allCombos[key]; ok {
				allCombos[key].count++
			} else {
				allCombos[key] = nc
			}
		}
	}
	for i := 1; i < len(combosPerGroup); i++ {
		//fmt.Printf("Attempting to multiply:\n\tA: %v\n\tB: %v\n", allCombos, combosPerGroup[i])
		allCombos = product2(allCombos, combosPerGroup[i])
		for key, combo := range allCombos {
			if !checkCombo(combo.values, groupCounts, true) {
				delete(allCombos, key)
			}
		}
		//fmt.Printf("newCombos: %v\n", allCombos)
	}

	var validCombos []*Combination
	for _, combo := range allCombos {
		if checkCombo(combo.values, groupCounts, false) {
			validCombos = append(validCombos, combo)
		}
	}

	var totalComboCount int
	for _, combination := range validCombos {
		totalComboCount += combination.count
	}
	return totalComboCount
}

func checkCombo(combo []int, groupCounts []int, allowPartial bool) bool {
	if !allowPartial && len(combo) != len(groupCounts) {
		return false
	}
	if len(combo) > len(groupCounts) {
		return false
	}
	for i := 0; i < len(combo); i++ {
		if combo[i] != groupCounts[i] {
			return false
		}
	}
	return true
}

func validPositions(groups []string, groupCounts []int) map[string]map[int][]int {
	validPos := make(map[string]map[int][]int)
	for _, group := range groups {
		validPos[group] = make(map[int][]int)
		for _, count := range groupCounts {
			// If we had a duplicate we can skip
			if _, ok := validPos[group][count]; ok {
				continue
			}

			// Don't bother looking at the group if it won't fit the range.
			if count > len(group) {
				continue
			}
			if count == len(group) {
				validPos[group][count] = append(validPos[group][count], 0)
				continue
			}

			// Figure out where in the group our ranges could sit.
			for i := range group[:len(group)-count+1] {
				//fmt.Printf("%d in %s => ", i, group[:len(group)-count])
				if i-1 >= 0 && group[i-1] == '#' {
					//fmt.Printf("# appears before\n")
					continue
				}
				if i+count < len(group) && group[i+count] == '#' {
					//fmt.Printf("# appears after\n")
					continue
				}
				//fmt.Printf("valid\n")
				validPos[group][count] = append(validPos[group][count], i)
			}
		}
	}
	return validPos
}

func positionCombinations(group string, positionsPerCount map[int][]int, groupCounts []int) [][]int {
	countPerPosition := make([][]int, len(group))
	for count, positions := range positionsPerCount {
		for _, position := range positions {
			countPerPosition[position] = append(countPerPosition[position], count)
		}
	}
	/*
		contains := func(str string, r rune) bool {
			for i := 0; i < len(str); i++ {
				if str[i] == byte(r) {
					return true
				}
			}
			return false
		}
	*/

	// Check that the sequence exists somewhere within the groupCounts
	contains := func(a []int, b []int) bool {
		for i := 0; i < len(a); i++ {
			var matchCount int
			for j := 0; j < len(b); j++ {
				if i+j >= len(a) || a[i+j] != b[j] {
					break
				}
				matchCount++
			}
			if matchCount == len(b) {
				return true
			}
		}
		return false
	}

	constraints := make(map[int]int)
	for _, count := range groupCounts {
		constraints[count]++
	}
	withinConstraints := func(combo []int) bool {
		matches := make(map[int]int)
		for _, count := range combo {
			matches[count]++
		}
		if len(matches) > len(constraints) {
			return false
		}
		for key, count := range matches {
			if count > constraints[key] {
				return false
			}
		}
		return contains(groupCounts, combo)
	}

	// Pre-compute the strings.Contains checks for sprints (#)
	var lastSpring int
	for i, char := range group {
		if char == '#' {
			lastSpring = i
		}
	}

	cachePerIdx := make(map[int][][]int)

	// NOTE: This was a failure - thought populating from the back of the list would help cut
	// down the branches a bit but it turns out it's till sooo fucking slow for large inputs
	var getCombosBackwards func(int) [][]int
	getCombosBackwards = func(idx int) [][]int {
		if idx < 0 {
			return cachePerIdx[0]
		}
		if idx >= len(group) {
			return nil
		}

		if val, ok := cachePerIdx[idx]; ok {
			return append([][]int{}, val...)
		}
		//fmt.Printf("getCombosBackwards('%d') ", idx)

		var combos [][]int
		for _, count := range countPerPosition[idx] {
			if idx+count+1 >= len(group) || lastSpring < idx+count+1 {
				combos = append(combos, []int{count})
			}

			nextCombos := getCombosBackwards(idx + count + 1)
			for _, c := range nextCombos {
				combo := append([]int{count}, c...)
				if withinConstraints(combo) {
					combos = append(combos, combo)
				}
			}
		}

		if group[idx] != '#' {
			for _, combo := range getCombosBackwards(idx + 1) {
				if withinConstraints(combo) {
					combos = append(combos, combo)
				}
			}
		}

		cachePerIdx[idx] = combos
		//fmt.Printf("=> combo len: %d => %v\n", len(combos), combos)
		return getCombosBackwards(idx - 1)
	}

	var getCombos func(start, maxLen int) [][]int
	getCombos = func(start, maxLen int) [][]int {
		if start >= len(group) || maxLen < 0 {
			return nil
		}

		if val, ok := cachePerIdx[start]; ok {
			return append([][]int{}, val...)
		}

		var combos [][]int
		for _, count := range countPerPosition[start] {
			if start+count+1 >= len(group) || lastSpring < start+count+1 {
				// We can only have _just_ this count if there are no mandatory springs.
				combos = append(combos, []int{count})
			}

			nextCombos := getCombos(start+count+1, maxLen-1) // Add a +1 to account for the '.'
			for _, c := range nextCombos {
				combo := append([]int{count}, c...)
				if len(combo) <= maxLen && withinConstraints(combo) {
					combos = append(combos, combo)
					//fmt.Printf("%d %v\n", count, c)
				}
			}
		}

		// Also determine what combo are available if we skip this index
		if group[start] != '#' {
			for _, combo := range getCombos(start+1, maxLen) {
				if len(combo) <= maxLen && withinConstraints(combo) {
					combos = append(combos, combo)
				}
			}
		}

		cachePerIdx[start] = combos
		return combos
	}

	// "Nothing" is valid if there are no mandatory springs.
	combos := getCombosBackwards(len(group) - 1)
	if !strings.Contains(group, "#") {
		combos = append(combos, []int{})
	}
	//fmt.Printf("**** Combos: %d ****\n", len(combos))
	return combos
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
