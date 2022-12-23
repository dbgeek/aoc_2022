package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type (
	vector struct {
		x int
		y int
	}

	filled map[vector]string
)

const (
	stone  = "#"
	sand   = "o"
	startX = 500
	startY = 0
)

func getBottomPos(f filled) vector {
	var (
		b vector
	)

	for k := range f {
		if k.y > b.y {
			b = k
		}
	}
	return b
}

func (v vector) String() string {
	return fmt.Sprintf("x: %d, y: %d", v.x, v.y)
}

func (v *vector) move(f filled, b vector) {

	for {

		if v.y+1 > b.y {
			break
		}
		// move down
		if _, ok := f[vector{x: v.x, y: v.y + 1}]; !ok {
			v.y += 1
			continue
		} else if _, ok := f[vector{x: v.x - 1, y: v.y + 1}]; !ok {
			// move left
			v.x -= 1
			v.y += 1
			continue
		} else if _, ok := f[vector{x: v.x + 1, y: v.y + 1}]; !ok {
			// move right
			v.x += 1
			v.y += 1
			continue
		}
		break
	}

}

func (v *vector) distance(s vector) vector {
	r := vector{
		x: s.x - v.x,
		y: s.y - v.y,
	}

	if r.x > 0 {
		r.x -= 1
	} else if r.x < 0 {
		r.x += 1
	} else if r.y > 0 {
		r.y -= 1
	} else if r.y < 0 {
		r.y += 1
	}

	return r
}

func visual(f filled) {

	minX, maxX := math.MaxInt, math.MinInt
	bottom := getBottomPos(f)

	for k := range f {

		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}

	}

	width := maxX - minX
	matrix := []string{}
	for a := 0; a <= bottom.y+3; a++ {
		row := []string{}
		for i := 0; i <= width; i++ {
			p := vector{x: i + minX, y: a}
			if v, ok := f[p]; ok {
				row = append(row, v)
			} else {
				row = append(row, ".")
			}
		}
		matrix = append(matrix, strings.Join(row, ""))
	}

	fmt.Println(strings.Join(matrix, "\n"))

}

func fillSand(s string) filled {

	f := make(filled)
	for _, row := range strings.Split(s, "\n") {
		if row == "" {
			continue
		}
		var (
			prevNode vector
		)
		for _, v := range strings.Split(row, " -> ") {
			w := strings.Split(strings.TrimSpace(v), ",")
			x, _ := strconv.Atoi(string(w[0]))
			y, _ := strconv.Atoi(string(w[1]))
			node := vector{x, y}
			f[node] = stone
			if prevNode != (vector{}) {
				distance := prevNode.distance(node)
				currentNode := prevNode

				for {
					if distance.x == 0 && distance.y == 0 {
						break
					}
					if distance.x > 0 {
						distance.x -= 1
						currentNode.x += 1
					} else if distance.x < 0 {
						distance.x += 1
						currentNode.x -= 1
					} else if distance.y > 0 {
						distance.y -= 1
						currentNode.y += 1
					} else if distance.y < 0 {
						distance.y += 1
						currentNode.y -= 1
					}
					f[currentNode] = stone
				}

			}
			prevNode = node
		}
	}
	return f
}

func Day14_1(input string) int {
	filled := fillSand(input)
	bottom := getBottomPos(filled)

	sandIdx := 0
	for {
		s := vector{x: startX, y: startY}
		s.move(filled, bottom)

		if s.y >= bottom.y {
			break
		} else {

			filled[s] = sand
		}

		sandIdx++
	}

	//visual(filled)
	return sandIdx
}

func Day14_2(input string) int {
	filled := fillSand(input)
	bottom := getBottomPos(filled)

	bottom.y += 1

	sandIdx := 0
	for {
		s := vector{x: startX, y: startY}
		s.move(filled, bottom)
		filled[s] = sand

		sandIdx++
		if s.y == startY && s.x == startX {
			break
		}

	}
	// visual(filled)
	return sandIdx
}
