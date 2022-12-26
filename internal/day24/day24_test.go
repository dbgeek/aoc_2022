package day24

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day24"
)

func TestDay24_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day24.Example,
			18,
		},
	}

	for _, tc := range tt {
		got := Day24_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay24_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day24.Example,
			54,
		},
	}

	for _, tc := range tt {
		got := Day24_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
