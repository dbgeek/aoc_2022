package day2

import (
	"fmt"
	"strings"
)

/*
  Rock      = 0
  Paper     = 1
  Scissors  = 2
*/

const (
	LosePoints = iota * 3
	DrawPoints
	WinPoints

	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

func Day2_1(input string) float64 {
	fmt.Println(LosePoints, DrawPoints, WinPoints)

	rows := strings.Split(input, "\n")

	var score float64

	for _, v := range rows {
		round := strings.Fields(v)

		if len(round) != 2 {
			continue
		}

		opponent := int(round[0][0] - 'A')
		me := int(round[1][0] - 'X')

		score += float64(me) + 1

		switch {
		case me == (opponent+1)%3: //win
			score += WinPoints
		case me == opponent: //draw
			score += DrawPoints
		case me == (opponent+2)%3: //lose
			continue
		}
	}
	return score
}

func Day2_2(input string) float64 {
	rows := strings.Split(input, "\n")

	var score float64

	for _, v := range rows {
		round := strings.Fields(v)
		if len(round) != 2 {
			continue
		}

		opponent := int(round[0][0] - 'A')

		switch round[1] {
		case Lose:
			score += float64((opponent+2)%3 + 1)
		case Draw:
			score += float64(opponent + 1 + 3)
		case Win:
			score += float64((opponent+1)%3 + 1 + 6)
		}

	}
	return score
}
