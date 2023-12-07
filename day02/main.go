package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

type game struct {
	id     int
	maxRgb rgb
}

type rgb struct {
	r, g, b int
}

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	games := parseGames(input)

	var sum int
	for _, game := range games {
		if game.maxRgb.r <= redCubes && game.maxRgb.g <= greenCubes && game.maxRgb.b <= blueCubes {
			sum += game.id
		}
	}
	return sum
}

func Part2(input string) int {
	games := parseGames(input)

	var power int
	for _, game := range games {
		power += game.maxRgb.r * game.maxRgb.g * game.maxRgb.b
	}
	return power
}

func parseGames(input string) []game {
	var games []game

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		id_game := strings.SplitN(line, ": ", 2)
		idString := id_game[0][len("Game "):]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		maxRgb := parseSets(id_game[1])
		games = append(games, game{id: id, maxRgb: maxRgb})
	}

	return games
}

func parseSets(sets string) rgb {
	var maxRgb rgb

	for _, set := range strings.Split(sets, "; ") {
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			amount_color := strings.Split(cube, " ")

			amount, err := strconv.Atoi(amount_color[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			switch amount_color[1] {
			case "red":
				if amount > maxRgb.r {
					maxRgb.r = amount
				}
			case "green":
				if amount > maxRgb.g {
					maxRgb.g = amount
				}
			case "blue":
				if amount > maxRgb.b {
					maxRgb.b = amount
				}
			}
		}
	}

	return maxRgb
}
