package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	points := 0

	scoreGame := func(_, count int) {
		switch count {
		case 0:
			points += 0
		case 1:
			points += 1
		case 2:
			points += 2
		case 3:
			points += 4
		case 4:
			points += 8
		case 5:
			points += 16
		case 6:
			points += 32
		case 7:
			points += 64
		case 8:
			points += 128
		case 9:
			points += 256
		case 10:
			points += 512
		}
	}

	run(input, scoreGame)

	return points
}

func run(input string, scoreGame func(i, count int)) {
	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		scratch := strings.SplitN(line, ": ", 2)[1]
		winning_got := strings.SplitN(scratch, " | ", 2)

		winning := make(map[string]bool)
		reader := strings.NewReader(winning_got[0])
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			winning[scanner.Text()] = true
		}

		var count int
		reader = strings.NewReader(winning_got[1])
		scanner = bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			if winning[scanner.Text()] {
				count++
			}
		}

		scoreGame(i+1, count)
	}
}

func Part2(input string) int {
	copiesById := make(map[int]int)

	scoreGame := func(i, count int) {
		_, ok := copiesById[i]
		if !ok {
			copiesById[i] = 1
		}
		for j := 1; j <= count; j++ {
			_, ok := copiesById[i+j]
			if !ok {
				copiesById[i+j] = 1
			}
			copiesById[i+j] += (1 * copiesById[i])
		}
	}

	run(input, scoreGame)

	points := 0
	for _, copies := range copiesById {
		points += copies
	}
	return points
}
