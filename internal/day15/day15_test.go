package day15

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day15"
)

func TestDay15_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		row   int
		out   int
	}{
		{
			"example",
			day15.Example,
			10,
			26,
		},
		{
			"my input",
			day15.Input,
			2000000,
			6275922,
		},
	}

	for _, tc := range tt {
		got := Day15_1(tc.input, tc.row)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay15_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{}

	for _, tc := range tt {
		got := Day15_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
