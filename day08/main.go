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

type LR struct {
	L, R string
}

func Part1(input string) int {
	lrByNode, instructions := parseInput(input)
	return stepsFromTo("AAA", func(node string) bool { return node == "ZZZ" }, instructions, lrByNode)
}

func parseInput(input string) (map[string]LR, []string) {
	sections := strings.SplitN(strings.TrimSpace(input), "\n\n", 2)

	lrByNode := make(map[string]LR)
	for _, line := range strings.Split(sections[1], "\n") {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == '=' || r == '(' || r == ')' || r == ','
		})
		lrByNode[fields[0]] = LR{fields[1], fields[2]}
	}

	instructions := strings.Split(sections[0], "")

	return lrByNode, instructions
}

func stepsFromTo(node string, end func(node string) bool, instructions []string, lrByNode map[string]LR) int {
	steps := 0
	for !end(node) {
		instruction := instructions[steps%len(instructions)]
		if instruction == "L" {
			node = lrByNode[node].L
		} else {
			node = lrByNode[node].R
		}
		steps++
	}
	return steps
}

func Part2(input string) int {
	lrByNode, instructions := parseInput(input)

	var starts []string
	for n := range lrByNode {
		if n[2] == 'A' {
			starts = append(starts, n)
		}
	}

	var steps []int
	for _, start := range starts {
		s := stepsFromTo(start, func(node string) bool { return node[2] == 'Z' }, instructions, lrByNode)
		steps = append(steps, s)
	}

	return LCM(steps[0], steps[1], steps[2:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
