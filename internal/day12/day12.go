package day12

import (
	"fmt"
	"strings"
)

type (
	Position struct {
		X     int
		Y     int
		Steps int
	}
)

func getNeigtborsBack(p Position, matrix [][]rune) []Position {
	neigtbors := []Position{}

	nrOfRows := int(len(matrix))
	numOfCols := int(len(matrix[p.Y]))

	if p.Y-1 >= 0 {
		if matrix[p.Y-1][p.X] >= matrix[p.Y][p.X]-1 {
			neigtbors = append(neigtbors, Position{Y: p.Y - 1, X: p.X, Steps: p.Steps + 1})
		}
	}
	if p.Y+1 < nrOfRows {
		if matrix[p.Y+1][p.X] >= matrix[p.Y][p.X]-1 {
			neigtbors = append(neigtbors, Position{Y: p.Y + 1, X: p.X, Steps: p.Steps + 1})
		}
	}
	if p.X-1 >= 0 {
		if matrix[p.Y][p.X-1] >= matrix[p.Y][p.X]-1 {
			neigtbors = append(neigtbors, Position{Y: p.Y, X: p.X - 1, Steps: p.Steps + 1})
		}
	}
	if p.X+1 < numOfCols {
		if matrix[p.Y][p.X+1] >= matrix[p.Y][p.X]-1 {
			neigtbors = append(neigtbors, Position{Y: p.Y, X: p.X + 1, Steps: p.Steps + 1})
		}
	}
	return neigtbors
}

func getNeigtbors(p Position, matrix [][]rune) []Position {
	neigtbors := []Position{}

	nrOfRows := int(len(matrix))
	numOfCols := int(len(matrix[p.Y]))

	if p.Y-1 >= 0 {
		if matrix[p.Y-1][p.X] <= matrix[p.Y][p.X]+1 {
			neigtbors = append(neigtbors, Position{Y: p.Y - 1, X: p.X, Steps: p.Steps + 1})
		}
	}
	if p.Y+1 < nrOfRows {
		if matrix[p.Y+1][p.X] <= matrix[p.Y][p.X]+1 {
			neigtbors = append(neigtbors, Position{Y: p.Y + 1, X: p.X, Steps: p.Steps + 1})
		}
	}
	if p.X-1 >= 0 {
		if matrix[p.Y][p.X-1] <= matrix[p.Y][p.X]+1 {
			neigtbors = append(neigtbors, Position{Y: p.Y, X: p.X - 1, Steps: p.Steps + 1})
		}
	}
	if p.X+1 < numOfCols {
		if matrix[p.Y][p.X+1] <= matrix[p.Y][p.X]+1 {
			neigtbors = append(neigtbors, Position{Y: p.Y, X: p.X + 1, Steps: p.Steps + 1})
		}
	}
	return neigtbors
}

func Day12_1(input string) int {
	var (
		matrix   [][]rune
		startPos Position
		endPos   Position
	)

	for i, v := range strings.Fields(input) {
		row := []rune{}
		for j, char := range v {
			row = append(row, char)
			chr := string(char)
			switch chr {
			case "S":
				startPos = Position{X: j, Y: i}
			case "E":
				endPos = Position{X: j, Y: i}
			}
		}
		matrix = append(matrix, row)
	}

	matrix[startPos.Y][startPos.X] = 'a'
	matrix[endPos.Y][endPos.X] = 'z'

	visit := make(map[Position]bool)
	positions := []Position{startPos}
	var (
		currPos      Position
		remainingPos []Position
	)
	for {
		currPos, remainingPos = positions[0], positions[1:]
		positions = remainingPos
		if _, ok := visit[currPos]; ok {
			continue
		}
		visit[currPos] = true

		if currPos.X == endPos.X && currPos.Y == endPos.Y {
			break
		}

		ns := getNeigtbors(currPos, matrix)
		for _, n := range ns {
			positions = append(positions, n)
		}
	}

	return currPos.Steps
}

func Day12_2(input string) int {
	var (
		matrix   [][]rune
		startPos Position
		endPos   Position
	)

	for i, v := range strings.Fields(input) {
		row := []rune{}
		for j, char := range v {
			row = append(row, char)
			chr := string(char)
			switch chr {
			case "S":
				startPos = Position{X: j, Y: i}
			case "E":
				endPos = Position{X: j, Y: i}
			}
		}
		matrix = append(matrix, row)
	}

	matrix[startPos.Y][startPos.X] = 'a'
	matrix[endPos.Y][endPos.X] = 'z'

	positions := []Position{
		{
			X: endPos.X, Y: endPos.Y},
	}
	var (
		currPos      Position
		remainingPos []Position
	)
	visit := make(map[Position]bool)
	for {
		currPos, remainingPos = positions[0], positions[1:]
		positions = remainingPos
		if _, ok := visit[currPos]; ok {
			continue
		}
		visit[currPos] = true

		if matrix[currPos.Y][currPos.X] == 'a' {
			fmt.Println("At end with steps: ", currPos.Steps)
			break
		}

		ns := getNeigtborsBack(currPos, matrix)
		for _, n := range ns {
			positions = append(positions, n)
		}
	}

	return currPos.Steps
}
