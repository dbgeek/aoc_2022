package day9

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day9"
)

func TestDay9_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example",
			day9.Example,
			13,
		},
		{
			"my input",
			day9.Input,
			5513,
		},
	}

	for _, tc := range tt {
		got := Day9_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay9_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day9.Example,
			1,
		},
		{
			"example part2",
			day9.ExamplePart2,
			36,
		},
		{
			"my input",
			day9.Input,
			2427,
		},
	}

	for _, tc := range tt {
		got := Day9_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
