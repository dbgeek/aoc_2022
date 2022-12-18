package day8

import (
	"strconv"
	"strings"
)

func createMatrix(s string) [][]uint8 {
	matrix := [][]uint8{}
	for y, row := range strings.Fields(s) {
		matrix = append(matrix, []uint8{})
		for _, char := range row {
			treeHigh, _ := strconv.Atoi(string(char))
			matrix[y] = append(matrix[y], uint8(treeHigh))
		}
	}
	return matrix
}

func getNeigtbors(x int8, y int8, matrix [][]uint8) [][]int8 {
	neigtbors := [][]int8{}

	nrOfRows := int8(len(matrix))
	numOfCols := int8(len(matrix[y]))

	if y-1 >= 0 {
		neigtbors = append(neigtbors, []int8{y - 1, x})
	}
	if y+1 < nrOfRows {
		neigtbors = append(neigtbors, []int8{y + 1, x})
	}
	if x-1 >= 0 {
		neigtbors = append(neigtbors, []int8{y, x - 1})
	}
	if x+1 < numOfCols {
		neigtbors = append(neigtbors, []int8{y, x + 1})
	}

	return neigtbors
}

func numberOfTrees(y int8, x int8, matrix [][]uint8) uint {

	var (
		trees uint = 1
		tmp   uint
	)
	tmp = numberTressDirection(x, y, -1, 0, matrix)
	trees *= tmp
	tmp = numberTressDirection(x, y, 1, 0, matrix)
	trees *= tmp
	tmp = numberTressDirection(x, y, 0, -1, matrix)
	trees *= tmp
	tmp = numberTressDirection(x, y, 0, 1, matrix)
	trees *= tmp
	return trees
}

func isTreeVisible(x int8, y int8, matrix [][]uint8) bool {

	if checkDirection(x, y, -1, 0, matrix) {
		return true
	}
	if checkDirection(x, y, 1, 0, matrix) {
		return true
	}
	if checkDirection(x, y, 0, -1, matrix) {
		return true
	}
	if checkDirection(x, y, 0, 1, matrix) {
		return true
	}

	return false
}
func numberTressDirection(x, y, dx, dy int8, matrix [][]uint8) uint {
	nrOfRows := int8(len(matrix))
	numOfCols := int8(len(matrix[y]))
	size := matrix[y][x]
	var treeCount uint
	for {
		x += dx
		y += dy
		if (y >= nrOfRows || y < 0) || (x < 0 || x >= numOfCols) {
			break
		}
		if matrix[y][x] >= size {
			treeCount += 1
			return treeCount
		}

		treeCount += 1

	}

	return treeCount

}

func checkDirection(x, y, dx, dy int8, matrix [][]uint8) bool {
	nrOfRows := int8(len(matrix))
	numOfCols := int8(len(matrix[y]))
	size := matrix[y][x]
	for {
		x += dx
		y += dy
		if (y >= nrOfRows || y < 0) || (x < 0 || x >= numOfCols) {
			break
		}

		if matrix[y][x] >= size {

			return false
		}

	}

	return true

}

func Day8_1(input string) int {

	matrix := createMatrix(input)

	nrOfVisibleTree := 0

	for y, row := range matrix {
		if y == 0 || y == len(matrix)-1 {
			// skip top and bottom edge
			continue
		}
		for x := range row {
			if x == 0 || x == len(row)-1 {
				// skip left and right edge
				continue
			}
			if isTreeVisible(int8(y), int8(x), matrix) {
				nrOfVisibleTree += 1
			}
		}
	}
	nrOfVisibleTree += (len(matrix) * 2) + ((len(matrix[0]) * 2) - 4)

	return nrOfVisibleTree
}

func Day8_2(input string) int {
	matrix := createMatrix(input)

	var (
		result uint
	)
	for y, row := range matrix {
		if y == 0 || y == len(matrix)-1 {
			// skip top and bottom edge
			continue
		}
		for x := range row {
			if x == 0 || x == len(row)-1 {
				// skip left and right edge
				continue
			}
			tmp := numberOfTrees(int8(y), int8(x), matrix)
			if tmp > result {
				result = tmp
			}
		}
	}
	return int(result)
}
