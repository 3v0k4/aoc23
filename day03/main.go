package main

import (
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

type coord struct {
	x, y int
}

func (c coord) neighbors9() []coord {
	var neighbors []coord
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			neighbors = append(neighbors, coord{x, y})
		}
	}
	return neighbors
}

func Part1(input string) int {
	matchSymbol := func(char rune) bool {
		return char != '.' && !(char >= '0' && char <= '9')
	}

	var sum int
	onValidPart := func(symbolCoord coord, part int) {
		sum += part
	}

	run(matchSymbol, onValidPart)

	return sum
}

func run(matchSymbol func(char rune) bool, onValidPart func(symbolCoord coord, part int)) {
	symbols := make(map[coord]bool)
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if matchSymbol(char) {
				symbols[coord{x, y}] = true
			}
		}
	}

	for y, line := range strings.Split(input, "\n") {
		var start int
		var end int
		var part int

		for x, char := range line {
			if char >= '0' && char <= '9' {
				if part == 0 {
					part = int(char - '0')
					start = x
				} else {
					part = part*10 + int(char-'0')
				}
			} else if part > 0 {
				end = x - 1
				if symbolCoord, ok := isValid(symbols, start, end, y); ok {
					onValidPart(symbolCoord, part)
				}
				part = 0
			}
		}

		if part > 0 {
			end = len(line) - 1
			if symbolCoord, ok := isValid(symbols, start, end, y); ok {
				onValidPart(symbolCoord, part)
			}
			part = 0
		}
	}
}

func isValid(symbols map[coord]bool, start, end, y int) (coord, bool) {
	for x := start; x <= end; x++ {
		for _, c := range (coord{x, y}).neighbors9() {
			if symbols[coord{c.x, c.y}] {
				return coord{c.x, c.y}, true
			}
		}
	}
	return coord{-1, -1}, false
}

func Part2(input string) int {
	matchSymbol := func(char rune) bool {
		return char == '*'
	}

	partsByGear := make(map[coord][]int)
	onValidPart := func(symbolCoord coord, part int) {
		partsByGear[symbolCoord] = append(partsByGear[symbolCoord], part)
	}

	run(matchSymbol, onValidPart)

	var ratio int
	for _, parts := range partsByGear {
		if len(parts) == 2 {
			ratio += parts[0] * parts[1]
		}
	}
	return ratio
}
