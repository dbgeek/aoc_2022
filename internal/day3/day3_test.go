package day3

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day3"
)

func TestDay3_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{
		{
			"example",
			day3.Example,
			157,
		},
		
	}

	for _, tc := range tt {
		got := Day3_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay3_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{

		{
			"example",
			day3.Example,
			70,
		},
	}

	for _, tc := range tt {
		got := Day3_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
