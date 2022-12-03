package day2

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day2"
)

func TestDay2_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{
		{
			"example",
			day2.Example,
			15,
		},
		{
			"myinput",
			day2.Input,
			17189,
		},
		
	}

	for _, tc := range tt {
		got := Day2_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay2_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{

		{
			"example",
			day2.Example,
			12,
		},
		{
			"myinput",
			day2.Input,
			13490,
		},
	}

	for _, tc := range tt {
		got := Day2_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
