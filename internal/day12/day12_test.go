package day12

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day12"
)

func TestDay9_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example",
			day12.Example,
			31,
		},
		{
			"myInput",
			day12.Input,
			361,
		},
	}

	for i, tc := range tt {
		if i == 0 {
			got := Day12_1(tc.input)
			t.Run(tc.name, func(t *testing.T) {
				if got != tc.out {
					t.Fatalf("got: %v want: %v", got, tc.out)
				}
			})
		}
	}
}

func TestDay12_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day12.Example,
      29,
		},
		{
			"my input",
			day12.Input,
      354,
		},
	}

	for _, tc := range tt {
		got := Day12_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
