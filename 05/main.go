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

	lowest, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Lowest location: %d\n", lowest)
}

type ConversionRule struct {
	sourceIdxStart      int
	sourceIdxEnd        int
	destinationIdxStart int
}

type Converter struct {
	source          string
	destination     string
	rules           []ConversionRule
	conversionCache map[int]int
}

func (c *Converter) convert(sourceID int) int {
	if result, ok := c.conversionCache[sourceID]; ok {
		return result
	}
	for _, rule := range c.rules {
		if sourceID < rule.sourceIdxStart || sourceID > rule.sourceIdxEnd {
			continue
		}
		offset := rule.destinationIdxStart - rule.sourceIdxStart
		return sourceID + offset
	}
	return sourceID
}

func parseConversionMap(lines []string) (*Converter, error) {
	var rules []ConversionRule
	for _, line := range lines[1:] {
		values, err := asInts(strings.Split(line, " "))
		if err != nil {
			return nil, err
		}
		rules = append(rules, ConversionRule{
			sourceIdxStart:      values[1],
			sourceIdxEnd:        values[1] + values[2] - 1,
			destinationIdxStart: values[0], // What sicko puts desination first...
		})
	}
	types := strings.Split(strings.Split(lines[0], " ")[0], "-to-")
	return &Converter{source: types[0], destination: types[1], rules: rules}, nil
}

func asInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	var err error
	for i, str := range strs {
		if str == "" {
			continue
		}
		if ints[i], err = strconv.Atoi(str); err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func solve(lines []string, useSeedRanges bool) (int, error) {
	// Create all of our converters.
	converters := make(map[string]*Converter)
	start := 2
	for i := 2; i <= len(lines); i++ {
		if i == len(lines) || lines[i] == "" {
			converter, err := parseConversionMap(lines[start:i])
			if err != nil {
				return 0, err
			}
			converters[converter.source] = converter
			start = i + 1
		}
	}

	seeds, err := asInts(strings.Split(strings.Split(lines[0], ": ")[1], " "))
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(seeds); i += 2 {
		seeds[i+1] = seeds[i] + seeds[i+1] - 1
	}

	source := "seed"
	for {
		if source == "location" {
			break
		}

		converter, ok := converters[source]
		if !ok {
			return 0, fmt.Errorf("unknown source %s", source)
		}

		seeds = expandRanges(seeds, converter)
		for i, seed := range seeds {
			seeds[i] = converter.convert(seed)
		}

		source = converter.destination
	}

	lowest := seeds[0]
	for _, val := range seeds[1:] {
		if val < lowest {
			lowest = val
		}
	}

	return lowest, nil
}

func expandRanges(ranges []int, converter *Converter) []int {
	seen := make(map[string]struct{})
	key := func(start, end int) string {
		return fmt.Sprintf("%d/%d", start, end)
	}

	for i := 0; i < len(ranges); i += 2 {
		for _, rule := range converter.rules {
			start := ranges[i]
			end := ranges[i+1]
			// Ignore ranges outside of the rules bounds.
			if start < rule.sourceIdxStart && end < rule.sourceIdxStart {
				continue
			}
			if start > rule.sourceIdxEnd && end > rule.sourceIdxEnd {
				continue
			}

			if start < rule.sourceIdxStart {
				k := key(start, rule.sourceIdxStart-1)
				if _, ok := seen[k]; !ok {
					seen[k] = struct{}{}
					ranges = append(ranges, start, rule.sourceIdxStart-1)
				}
				ranges[i] = rule.sourceIdxStart
				seen[key(ranges[i], ranges[i+1])] = struct{}{}
			}

			if end > rule.sourceIdxEnd {
				k := key(rule.sourceIdxEnd+1, end)
				if _, ok := seen[k]; !ok {
					seen[k] = struct{}{}
					ranges = append(ranges, rule.sourceIdxEnd+1, end)
				}
				ranges[i+1] = rule.sourceIdxEnd
				seen[key(ranges[i], ranges[i+1])] = struct{}{}
			}
		}
	}

	return ranges
}
