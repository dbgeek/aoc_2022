package day5

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day5"
)

func TestDay4_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   string
	}{
		{
			"example",
			day5.Example,
		  "CMZ",
		},
		
	}

	for _, tc := range tt {
		got := Day5_1(tc.input)
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
		out   string
	}{

		{
			"example",
			day5.Example,
			"MCD",
		},
	}

	for _, tc := range tt {
		got := Day5_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
