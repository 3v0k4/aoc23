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

type Coord struct {
	x, y int
}

func Part1(input string) int {
	start := Coord{-1, -1}
	tiles := make(map[Coord]rune)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		for x, char := range line {
			if char != '|' && char != '-' && char != '7' && char != 'F' && char != 'L' && char != 'J' && char != 'S' {
				continue
			}
			tiles[Coord{x, y}] = char
			if char == 'S' {
				start = Coord{x, y}
			}
		}
	}

	prev := Coord{-1, -1}
	curr := start
	steps := 0
	for curr != start || steps == 0 {
		next := nextFrom(curr, prev, tiles)
		prev = curr
		curr = next
		steps++
	}
	return steps / 2
}

func nextFrom(curr, exclude Coord, tiles map[Coord]rune) Coord {
	cChar := tiles[curr]

	// south
	coord := Coord{curr.x, curr.y + 1}
	if coord != exclude {
		char, ok := tiles[coord]
		if ok {
			if (char == 'S' || char == '|' || char == 'L' || char == 'J') && (cChar != '-' && cChar != 'J' && cChar != 'L') {
				return coord
			}
		}
	}

	// east
	coord = Coord{curr.x + 1, curr.y}
	if coord != exclude {
		char, ok := tiles[coord]
		if ok {
			if (char == 'S' || char == '-' || char == 'J' || char == '7') && (cChar != '|' && cChar != 'J' && cChar != '7') {
				return coord
			}
		}
	}

	// north
	coord = Coord{curr.x, curr.y - 1}
	if coord != exclude {
		char, ok := tiles[coord]
		if ok {
			if (char == 'S' || char == '|' || char == '7' || char == 'F') && (cChar != '-' && cChar != '7' && cChar != 'F') {
				return coord
			}
		}
	}

	// west
	coord = Coord{curr.x - 1, curr.y}
	if coord != exclude {
		char, ok := tiles[coord]
		if ok {
			if (char == 'S' || char == '-' || char == 'L' || char == 'F') && (cChar != '|' && cChar != 'L' && cChar != 'F') {
				return coord
			}
		}
	}

	panic("wth")
}

func Part2(input string) int {
	start := Coord{-1, -1}
	tiles := make(map[Coord]rune)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		for x, char := range line {
			if char != '|' && char != '-' && char != '7' && char != 'F' && char != 'L' && char != 'J' && char != 'S' {
				continue
			}
			tiles[Coord{x, y}] = char
			if char == 'S' {
				start = Coord{x, y}
			}
		}
	}

	loop := make(map[Coord]rune)
	prev := Coord{-1, -1}
	second := Coord{-1, -1}
	curr := start
	steps := 0
	topLeft := Coord{math.MaxInt, math.MaxInt}
	bottomRight := Coord{0, 0}
	for curr != start || steps == 0 {
		if curr.y < topLeft.y {
			topLeft = Coord{topLeft.x, curr.y}
		}
		if curr.x < topLeft.x {
			topLeft = Coord{curr.x, topLeft.y}
		}
		if curr.y > bottomRight.y {
			bottomRight = Coord{bottomRight.x, curr.y}
		}
		if curr.x > bottomRight.x {
			bottomRight = Coord{curr.x, bottomRight.y}
		}
		if steps == 1 {
			second = curr
		}
		loop[curr] = tiles[curr]
		next := nextFrom(curr, prev, tiles)
		prev = curr
		curr = next
		steps++
	}

	if prev.x == second.x {
		loop[start] = '|'
	} else if prev.y == second.y {
		loop[start] = '-'
	} else if prev.y == second.y-1 && prev.x == second.x-1 && start.y == prev.y {
		loop[start] = '7'
	} else if prev.y == second.y-1 && prev.x == second.x-1 && start.x == prev.x {
		loop[start] = 'L'
	} else if prev.y == second.y-1 && prev.x == second.x+1 && start.y == prev.y {
		loop[start] = 'F'
	} else if prev.y == second.y-1 && prev.x == second.x+1 && start.x == prev.x {
		loop[start] = 'J'
	} else if prev.y == second.y+1 && prev.x == second.x-1 && start.y == prev.y {
		loop[start] = 'F'
	} else if prev.y == second.y+1 && prev.x == second.x-1 && start.x == prev.x {
		loop[start] = 'J'
	} else if prev.y == second.y+1 && prev.x == second.x+1 && start.y == prev.y {
		loop[start] = 'L'
	} else if prev.y == second.y+1 && prev.x == second.x+1 && start.x == prev.x {
		loop[start] = '7'
	}

	// Ray casting algorithm
	// https://en.wikipedia.org/wiki/Point_in_polygon#Ray_casting_algorithm
	count := 0
	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			_, ok := loop[Coord{x, y}]
			if ok {
				continue
			}
			edges := 0
			p := ' '
			for xx := x; xx <= bottomRight.x; xx++ {
				t := loop[Coord{xx, y}]
				if t == '|' { // crossed an edge
					edges++
					p = ' '
					continue
				}
				if t == '-' { // on an edge
					continue
				}
				if (p == ' ') && (t == 'L' || t == 'J' || t == '7' || t == 'F') { // on a turn
					p = t
					continue
				}
				if (p != ' ') && (t == 'L' || t == 'J' || t == '7' || t == 'F') { // on a subsequent turn
					if p == 'L' && t == 'L' {
						edges++
					} else if p == 'L' && t == 'J' {
						edges++
						edges++
					} else if p == 'L' && t == '7' {
						edges++
					} else if p == 'L' && t == 'F' {
						edges++
						edges++
					} else if p == 'J' && t == 'L' {
						edges++
						edges++
					} else if p == 'J' && t == 'J' {
						edges++
					} else if p == 'J' && t == '7' {
						edges++
						edges++
					} else if p == 'J' && t == 'F' {
						edges++
					} else if p == '7' && t == 'L' {
						edges++
					} else if p == '7' && t == 'J' {
						edges++
						edges++
					} else if p == '7' && t == '7' {
						edges++
					} else if p == '7' && t == 'F' {
						edges++
						edges++
					} else if p == 'F' && t == 'L' {
						edges++
						edges++
					} else if p == 'F' && t == 'J' {
						edges++
					} else if p == 'F' && t == '7' {
						edges++
						edges++
					} else if p == 'F' && t == 'F' {
						edges++
					}
					p = ' '
					continue
				}
			}
			if edges%2 == 1 {
				count++
			}
		}
	}
	return count
}
