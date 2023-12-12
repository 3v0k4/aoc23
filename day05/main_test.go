package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	input := string(file)
	actual := Part1(input)
	expect := 579439039

	if actual != expect {
		t.Errorf("Actual: %d, Expect: %d", actual, expect)
	}
}

func BenchmarkPart1(b *testing.B) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		b.Errorf(err.Error())
	}

	input := string(file)

	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	input := string(file)
	actual := Part2(input)
	expect := 7873084

	if actual != expect {
		t.Errorf("Actual: %d, Expect: %d", actual, expect)
	}
}

func BenchmarkPart2(b *testing.B) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		b.Errorf(err.Error())
	}

	input := string(file)

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
