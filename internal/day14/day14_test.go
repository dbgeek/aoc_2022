package day14

import (
	"testing"

	"github.com/dbgeek/aoc_2022/internal/data/day14"
)

func TestDay14_part1(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day14.Example,
			24,
		},
		{
			"example 2",
			day14.Input,
			737,
		},
	}

	for _, tc := range tt {
		got := Day14_1(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}

func TestDay14_part2(t *testing.T) {

	tt := []struct {
		name  string
		input string
		out   int
	}{
		{
			"example 1",
			day14.Example,
			93,
		},
		{
			"example 2",
			day14.Input,
			28145,
		},
	}

	for _, tc := range tt {
		got := Day14_2(tc.input)
		t.Run(tc.name, func(t *testing.T) {
			if got != tc.out {
				t.Fatalf("got: %v want: %v", got, tc.out)
			}
		})
	}
}
