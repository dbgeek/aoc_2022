package day10

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day10"
)

func TestDay10_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example",
			day10.Example,
			13140,
		},
	}

	for i, tc := range tt {
		if i == 0 {
			got := Day10_part1(tc.input)
			t.Run(tc.name, func(t *testing.T) {
				if got != tc.out {
					t.Fatalf("got: %v want: %v", got, tc.out)
				}
			})
		}
	}
}

func TestDay10_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{}

	for _, tc := range tt {
		got := Day10_part2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
