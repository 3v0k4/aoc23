package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/google/btree"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	tree := btree.New(2)
	tally := make(map[Hand]int)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var hand string
		var bid int
		fmt.Sscanf(line, "%s %d", &hand, &bid)
		parsed := Hand(hand)
		tree.ReplaceOrInsert(parsed)
		tally[parsed] = bid
	}

	var sum int
	rank := 1
	iterator := btree.ItemIterator(func(item btree.Item) bool {
		sum += rank * tally[item.(Hand)]
		rank++
		return true
	})
	tree.Ascend(iterator)

	return sum
}

func Part2(input string) int {
	tree := btree.New(2)
	tally := make(map[Hand]int)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var hand string
		var bid int
		fmt.Sscanf(line, "%s %d", &hand, &bid)
		hand = strings.ReplaceAll(hand, "J", "*")
		parsed := Hand(hand)
		tree.ReplaceOrInsert(parsed)
		tally[parsed] = bid
	}

	var sum int
	rank := 1
	iterator := btree.ItemIterator(func(item btree.Item) bool {
		sum += rank * tally[item.(Hand)]
		rank++
		return true
	})
	tree.Ascend(iterator)

	return sum
}

type Hand string

func (hand Hand) Less(than btree.Item) bool {
	return hand.Strength() < than.(Hand).Strength()
}

func (hand Hand) Strength() string {
	return hand.TypeStrength() + hand.CardsStrength()
}

func (hand Hand) TypeStrength() string {
	jokers := 0
	tally := make(map[rune]int)
	for _, card := range hand {
		if card == '*' {
			jokers++
			continue
		}
		_, ok := tally[card]
		if ok {
			tally[card] += 1
		} else {
			tally[card] = 1
		}
	}

	var same []int
	for _, v := range tally {
		same = append(same, v)
	}
	sort.Sort(sort.IntSlice(same))

	if jokers == 5 {
		same = []int{5}
	} else if jokers > 0 {
		same[len(same)-1] += jokers
	}

	if eq(same, []int{5}) {
		return "7"
	} else if eq(same, []int{1, 4}) {
		return "6"
	} else if eq(same, []int{2, 3}) {
		return "5"
	} else if eq(same, []int{1, 1, 3}) {
		return "4"
	} else if eq(same, []int{1, 2, 2}) {
		return "3"
	} else if eq(same, []int{1, 1, 1, 2}) {
		return "2"
	} else if eq(same, []int{1, 1, 1, 1, 1}) {
		return "1"
	}
	panic("wth")
}

func eq(as, bs []int) bool {
	if len(as) != len(bs) {
		return false
	}
	for i, a := range as {
		if a != bs[i] {
			return false
		}
	}
	return true
}

func (hand Hand) CardsStrength() string {
	var sum string
	for _, r := range hand {
		if r == 'A' {
			sum += "14"
		} else if r == 'K' {
			sum += "13"
		} else if r == 'Q' {
			sum += "12"
		} else if r == 'J' {
			sum += "11"
		} else if r == '*' {
			sum += "01"
		} else if r == 'T' {
			sum += "10"
		} else if r >= '2' && r <= '9' {
			sum += ("0" + string(r))
		} else {
			panic("wth")
		}
	}
	return sum
}
