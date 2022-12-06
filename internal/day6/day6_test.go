package day6

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day6"
)

func TestFoundMarker(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   bool
	}{
		{
			"example1",
			"abcd",
			true,
		},
		{
			"example 2",
			"aBcC",
			true,
		},
		{
			"example 3",
			"aAAb",
			false,
		},
		{
			"example 4",
			"AaaB",
			false,
		},
	}

	for _, tc := range tt {
		got := foundMarker(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
func TestDay6_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example1",
			day6.Example1,
			7,
		},
		{
			"example 2",
			day6.Example2,
			5,
		},
		{
			"example 3",
			day6.Example3,
			6,
		},
		{
			"example 4",
			day6.Example4,
			10,
		},
		{
			"example 5",
			day6.Example5,
			11,
		},
	}

	for _, tc := range tt {
		got := Day6_1(tc.input)
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
			day6.Example1,
			19,
		},
		{
			"example 2",
			day6.Example2,
			23,
		},
		{
			"example 3",
			day6.Example3,
			23,
		},
		{
			"example 4",
			day6.Example4,
			29,
		},
		{
			"example 5",
			day6.Example5,
			26,
		},
	}

	for _, tc := range tt {
		got := Day6_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
