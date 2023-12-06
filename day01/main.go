package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(Part1(string(file)))
	fmt.Println(Part2(string(file)))
}

func Part1(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last int

		for i := 0; i < len(line); i++ {
			parsed, ok := parseDigit1(line[i])
			if ok {
				first = parsed
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			parsed, ok := parseDigit1(line[i])
			if ok {
				last = parsed
				break
			}
		}

		sum += first*10 + last
	}

	return sum
}

func parseDigit1(b byte) (int, bool) {
	if b >= '0' && b <= '9' {
		return int(b - '0'), true
	} else {
		return 0, false
	}
}

func Part2(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		var first, last int

		for i := 0; i < len(line); i++ {
			parsed, ok := parseDigit2(line[i:])
			if ok {
				first = parsed
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			parsed, ok := parseDigit2(line[i:])
			if ok {
				last = parsed
				break
			}
		}

		sum += first*10 + last
	}

	return sum
}

func parseDigit2(str string) (int, bool) {
	if len(str) == 0 {
		return 0, false
	}

	parsed, ok := parseDigit1(str[0])
	if ok {
		return parsed, true
	}

	if len(str) < 3 {
		return 0, false
	}

	switch str[0] {
	case 'e':
		if len(str) >= 5 && str[1:5] == "ight" {
			return int('8' - '0'), true
		}
	case 'f':
		if len(str) >= 4 && str[1:4] == "ive" {
			return int('5' - '0'), true
		} else if len(str) >= 4 && str[1:4] == "our" {
			return int('4' - '0'), true
		}
	case 'n':
		if len(str) >= 4 && str[1:4] == "ine" {
			return int('9' - '0'), true
		}
	case 'o':
		if len(str) >= 3 && str[1:3] == "ne" {
			return int('1' - '0'), true
		}
	case 's':
		if len(str) >= 3 && str[1:3] == "ix" {
			return int('6' - '0'), true
		} else if len(str) >= 5 && str[1:5] == "even" {
			return int('7' - '0'), true
		}
	case 't':
		if len(str) >= 5 && str[1:5] == "hree" {
			return int('3' - '0'), true
		} else if len(str) >= 3 && str[1:3] == "wo" {
			return int('2' - '0'), true
		}
	}

	return 0, false
}
