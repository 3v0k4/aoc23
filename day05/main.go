package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	sections := strings.Split(input, "\n\n")
	seedLine, _ := strings.CutPrefix(sections[0], "seeds: ")
	seedStrings := strings.Fields(seedLine)
	var seeds []int
	for _, seedString := range seedStrings {
		var seed int
		fmt.Sscanf(seedString, "%d", &seed)
		seeds = append(seeds, seed)
	}

	var converted, todos []int
	converted = seeds
	for _, sectionString := range sections[1:] {
		conversions := strings.Split(sectionString, "\n")[1:]
		todos = append(todos, converted...)
		converted = nil
		for _, conversion := range conversions {
			var dest, source, length int
			fmt.Sscanf(conversion, "%d %d %d", &dest, &source, &length)
			var newConverted []int
			newConverted, todos = convert(todos, source, source+length-1, dest-source)
			converted = append(converted, newConverted...)
		}
	}

	min := math.MaxInt64
	all := append(converted, todos...)
	for _, location := range all {
		if location < min {
			min = location
		}
	}
	return min
}

func convert(elements []int, start int, end int, delta int) (converted []int, todos []int) {
	for _, element := range elements {
		if start <= element && element <= end {
			converted = append(converted, element+delta)
		} else {
			todos = append(todos, element)
		}
	}

	return
}

type Range struct {
	start, end int
}

func Part2(input string) int {
	sections := strings.Split(input, "\n\n")
	seedLine, _ := strings.CutPrefix(sections[0], "seeds: ")
	seedStrings := strings.Fields(seedLine)
	var seeds []Range
	for i := 0; i < len(seedStrings); i += 2 {
		seedString := seedStrings[i] + " " + seedStrings[i+1]
		var start, end int
		fmt.Sscanf(seedString, "%d %d", &start, &end)
		seeds = append(seeds, Range{start, start + end - 1})
	}

	var converted, todos []Range
	converted = seeds
	for _, sectionString := range sections[1:] {
		conversions := strings.Split(sectionString, "\n")[1:]
		todos = append(todos, converted...)
		converted = nil
		for _, conversion := range conversions {
			var dest, source, length int
			fmt.Sscanf(conversion, "%d %d %d", &dest, &source, &length)
			var newConverted []Range
			newConverted, todos = convertRanges(todos, source, source+length-1, dest-source)
			converted = append(converted, newConverted...)
		}
	}

	min := math.MaxInt64
	all := append(converted, todos...)

	for _, location := range all {
		if location.start < min {
			min = location.start
		}
	}

	return min
}

func convertRanges(elements []Range, start int, end int, delta int) (converted []Range, todos []Range) {
	for _, element := range elements {
		// E    -----
		// C  ----------
		if start <= element.start && element.end <= end {
			converted = append(converted, Range{element.start + delta, element.end + delta})
			continue
		}

		// E            -----
		// C  ----------
		if end < element.start {
			todos = append(todos, element)
			continue
		}

		// E  -----
		// C       ----------
		if element.end < start {
			todos = append(todos, element)
			continue
		}

		// E  -----
		// C     ----------
		if element.start < start && start <= element.end {
			todos = append(todos, Range{element.start, start - 1})
			converted = append(converted, Range{start + delta, element.end + delta})
			continue
		}

		// E          -----
		// C  ----------
		if element.start <= end && end < element.end {
			todos = append(todos, Range{end + 1, element.end})
			converted = append(converted, Range{element.start + delta, end + delta})
			continue
		}
	}

	return
}
