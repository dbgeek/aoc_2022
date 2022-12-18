package day8

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day8"
)

func TestDay8_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example",
			day8.Example,
			21,
		},
		{
			"my input",
			day8.Input,
			1688,
		},
	}

	for _, tc := range tt {
		got := Day8_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay6_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example1",
			day8.Example,
			8,
		},
		{
			"my input",
			day8.Input,
			410400,
		},
	}

	for _, tc := range tt {
		got := Day8_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
