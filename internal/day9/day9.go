package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	vector struct {
		x int
		y int
	}

	rope struct {
		head       vector
		tail       []*vector
		tail_visit map[string]bool
	}
)

var (
	l = vector{
		x: -1,
		y: 0,
	}
	r = vector{
		x: 1,
		y: 0,
	}
	u = vector{
		x: 0,
		y: 1,
	}
	d = vector{
		x: 0,
		y: -1,
	}
)

func (r *rope) move(m vector, s int) {

	for i := 0; i < s; i++ {
		// Moving the head
		r.head.x += m.x
		r.head.y += m.y

		prevNode := r.head
		for _, v := range r.tail {
			// Calculating distance from head
			dx := prevNode.x - v.x
			dy := prevNode.y - v.y

			// Same Row or colum
			if (prevNode.x == v.x) || (prevNode.y == v.y) {
				if (dx > 1 || dy > 1) || (dx < -1 || dy < -1) {
					if dx < 0 {
						v.x += dx + 1
					} else if dx > 0 {
						v.x += dx - 1
					}
					if dy < 0 {
						v.y += dy + 1
					} else if dy > 0 {
						v.y += dy - 1
					}
				}
			} else if (dx > 1 || dy > 1) || (dx < -1 || dy < -1) {
				if dx < 0 && dy < 0 {
					v.x += -1
					v.y += -1
				} else if dx < 0 && dy > 0 {
					v.x += -1
					v.y += 1
				} else if dx > 0 && dy < 0 {
					v.x += 1
					v.y += -1
				} else if dx > 0 && dy > 0 {
					v.x += 1
					v.y += 1
				}
			}
			prevNode = *v
		}
		l := len(r.tail)
		tailKey := fmt.Sprintf("%d:%d", r.tail[l-1].x, r.tail[l-1].y)
		if !r.tail_visit[tailKey] {
			r.tail_visit[tailKey] = true
		}
	}
}

func draw(r rope) {
	for i, v := range r.tail {
		fmt.Println(i+1, *v)
	}

	matrix := [][]string{}
	
  for row := 0; row <= 35; row++ {
		r := []string{}

    for col := 0; col < 35; col++ {
			r = append(r, ". ")
		}

		matrix = append(matrix, r)

	}

	for i, v := range r.tail {
		matrix[v.y][v.x] = fmt.Sprintf("%d ", i+1)
	}

	matrix[r.head.y][r.head.x] = "H "

	for i := len(matrix) - 1; i >= 0; i-- {
		fmt.Println(strings.Join(matrix[i], ""))
	}
}

func Day9_1(input string) int {
	fields := strings.Fields(input)
	rope := rope{
		head: vector{
			x: 0,
			y: 0,
		},
		tail: []*vector{
			{
				x: 0,
				y: 0,
			},
		},
		tail_visit: make(map[string]bool),
	}
	for i := 0; i < len(fields); i += 2 {
		steps, _ := strconv.Atoi(fields[i+1])
		switch fields[i] {
		case "L":
			rope.move(l, steps)
		case "R":
			rope.move(r, steps)
		case "D":
			rope.move(d, steps)
		case "U":
			rope.move(u, steps)
		}
	}
	return len(rope.tail_visit)
}

func Day9_2(input string) int {
	fields := strings.Fields(input)
	rope := rope{
		head: vector{
			x: 12,
			y: 10,
		},
		tail: []*vector{
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
			{
				x: 12,
				y: 10,
			},
		},
		tail_visit: make(map[string]bool),
	}
	for i := 0; i < len(fields); i += 2 {
		steps, _ := strconv.Atoi(fields[i+1])
		switch fields[i] {
		case "L":
			rope.move(l, steps)
		case "R":
			rope.move(r, steps)
		case "D":
			rope.move(d, steps)
		case "U":
			rope.move(u, steps)
		}
	}
	return len(rope.tail_visit)
}
