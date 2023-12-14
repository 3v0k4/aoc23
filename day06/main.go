package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	timesString, _ := strings.CutPrefix(lines[0], "Time:")
	var times []int
	for _, timeString := range strings.Fields(timesString) {
		time, _ := strconv.Atoi(timeString)
		times = append(times, time)
	}

	distsString, _ := strings.CutPrefix(lines[1], "Distance:")
	var dists []int
	for _, distString := range strings.Fields(distsString) {
		dist, _ := strconv.Atoi(distString)
		dists = append(dists, dist)
	}

	return solve(times, dists)
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	timesString, _ := strings.CutPrefix(lines[0], "Time:")
	var timeSingle string
	for _, timeString := range strings.Fields(timesString) {
		timeSingle = timeSingle + timeString
	}
	time, _ := strconv.Atoi(timeSingle)
	times := []int{time}

	distsString, _ := strings.CutPrefix(lines[1], "Distance:")
	var distSingle string
	for _, distString := range strings.Fields(distsString) {
		distSingle = distSingle + distString
	}
	dist, _ := strconv.Atoi(distSingle)
	dists := []int{dist}

	return solve(times, dists)
}

func solve(times, dists []int) int {
	ways := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		dist := dists[i]
		// dist = (time - press) * press
		// => dist = (t - x) * x
		// => f(x) = t*x - x^2
		// It's a parabola with the vertex at the top, which
		// means it's at the peak when its derivative == 0.
		// f'(x) = t - 2x
		// f'(x) = 0 => t - 2x = 0 => x = t/2
		max := time / 2

		currentWays := 0
		if (time-max)*max > dist {
			currentWays++
		}

		j := 1
		for {
			if (time-(max+j))*(max+j) > dist {
				currentWays++
			}

			if (time-(max-j))*(max-j) > dist {
				currentWays++
			}

			if (time-(max+j))*(max+j) <= dist && (time-(max-j))*(max-j) <= dist {
				break
			}

			j++
		}

		ways *= currentWays
	}

	return ways
}
