package day1

import (
	"github.com/dbgeek/aoc_2022/internal/data/day1"
	"testing"
)

func TestDay1_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int64
	}{
		{
			"example",
			day1.Example,
			24000,
		},
		{
			"myinput",
			day1.Part,
			75622,
		},
	}

	for _, tc := range tt {
		got := Day1_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}


func TestDay1_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   float64
	}{

		{
			"example",
			day1.Example,
			45000,
		},
		{
			"myinput",
			day1.Part,
			213159,
		},
	}

	for _, tc := range tt {
		got := Day1_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
