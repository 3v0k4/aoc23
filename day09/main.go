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

func Part1(input string) int {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		tree := make(map[int][]int, len(lines))
		fields := strings.Fields(line)
		tree[0] = make([]int, len(fields))

		for i, field := range fields {
			var n int
			fmt.Sscanf(field, "%d", &n)
			tree[0][i] = n
		}

		for level := 1; level < len(tree[0]); level++ {
			tree[level] = make([]int, len(tree[level-1])-1)
			allZeroes := true
			for i := 0; i < len(tree[level-1])-1; i++ {
				a := tree[level-1][i]
				b := tree[level-1][i+1]
				tree[level][i] = b - a
				if b-a != 0 {
					allZeroes = false
				}
			}
			if allZeroes {
				break
			}
		}

		extrapolated := 0
		for level := len(tree) - 2; level >= 0; level = level - 1 {
			l := len(tree[level])
			last := tree[level][l-1]
			extrapolated += last
		}
		result += extrapolated

	}
	return result
}

func Part2(input string) int {
	var result int

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		tree := make(map[int][]int, len(lines))
		fields := strings.Fields(line)
		tree[0] = make([]int, len(fields))

		for i, field := range fields {
			var n int
			fmt.Sscanf(field, "%d", &n)
			tree[0][i] = n
		}

		for level := 1; level < len(tree[0]); level++ {
			tree[level] = make([]int, len(tree[level-1])-1)
			allZeroes := true
			for i := 0; i < len(tree[level-1])-1; i++ {
				a := tree[level-1][i]
				b := tree[level-1][i+1]
				tree[level][i] = b - a
				if b-a != 0 {
					allZeroes = false
				}
			}
			if allZeroes {
				break
			}
		}

		extrapolated := 0
		for level := len(tree) - 2; level >= 0; level = level - 1 {
			first := tree[level][0]
			extrapolated = first - extrapolated
		}
		result += extrapolated

	}
	return result
}
