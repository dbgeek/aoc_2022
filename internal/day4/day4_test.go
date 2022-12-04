package day4

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day4"
)

func TestDay4_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{
		{
			"example",
			day4.Example,
			2,
		},
		
	}

	for _, tc := range tt {
		got := Day4_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay4_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{

		{
			"example",
			day4.Example,
			4,
		},
	}

	for _, tc := range tt {
		got := Day4_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
