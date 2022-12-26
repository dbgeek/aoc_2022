package day24

import (
	"fmt"
	"strings"
)

type (
	position struct {
		x     int
		y     int
		m     int
		start bool
		end   bool
	}

	direction int

	blizzard struct {
		currentPosition position
		direction       direction
	}
)

const (
	right direction = iota
	down
	left
	up
)

var (
	directionMap = map[string]direction{
		"<": left,
		"v": down,
		"^": up,
		">": right,
	}
)

func (p position) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y) //, p.m)
}

func (b *blizzard) move(w, h int) {
	d := []position{
		{
			x: 1, y: 0,
		},
		{
			x: 0, y: 1,
		},
		{
			x: -1, y: 0,
		},
		{
			x: 0, y: -1,
		},
	}
	newX, newY := b.currentPosition.x+d[b.direction].x, b.currentPosition.y+d[b.direction].y
	if newX == 0 {
		newX = w - 1
	} else if w == newX {
		newX = 1
	}

	if newY == 0 {
		newY = h - 1
	} else if h == newY {
		newY = 1
	}

	b.currentPosition.x = newX
	b.currentPosition.y = newY

}

func parseDirectionString(s string) (direction, bool) {
	c, ok := directionMap[strings.ToLower(s)]
	return c, ok
}

func (d direction) String() string {
	switch d {
	case left:
		return "<"
	case right:
		return ">"
	case up:
		return "^"
	case down:
		return "v"
	}
	return "unknown"
}

func (d direction) newPos(p position) position {
	switch d {
	case left:
		return position{x: p.x - 1, y: p.y}
	case right:
		return position{x: p.x + 1, y: p.y}
	case up:
		return position{x: p.x, y: p.y - 1}
	case down:
		return position{x: p.x, y: p.y + 1}
	}

	return position{}

}

// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via gcd
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func parseInput(s string) ([]blizzard, [2]int) {
	blizzards := []blizzard{}
	widthHight := [2]int{}
	for rIdx, row := range strings.Split(s, "\n") {
		if row == "" {
			continue
		}
		for cIdx, col := range strings.Split(row, "") {
			if v, ok := parseDirectionString(col); ok {
				blizzards = append(blizzards, blizzard{
					currentPosition: position{x: cIdx, y: rIdx}, direction: v,
				})
			}
			widthHight[0] = cIdx
		}
		widthHight[1] = rIdx
	}
	return blizzards, widthHight
}

func printGrid(width, height int, cp position, bs []blizzard, startEndX [2]int) {
	matrix := [][]string{}

	for i := 0; i <= height; i++ {
		s := []string{}
		for j := 0; j <= width; j++ {
			if startEndX[0] == j && i == 0 {
				s = append(s, ".")
			} else if startEndX[1] == j && i == height {
				s = append(s, ".")
			} else if i == 0 || i == height {
				s = append(s, "#")
			} else if j == 0 || j == width {
				s = append(s, "#")
			} else {
				s = append(s, ".")
			}
		}
		matrix = append(matrix, s)
	}

	for _, b := range bs {
		matrix[b.currentPosition.y][b.currentPosition.x] = b.direction.String()
	}

	matrix[cp.y][cp.x] = "E"

	for _, v := range matrix {
		fmt.Println(strings.Join(v, ""))
	}
}

func getNeigtbors(p position, b []blizzard, se []position, width, height int) []position {
	//neigtbors := []position{{x: p.x, y: p.y, m: p.m + 1}}
	neigtbors := []position{}
	blizzardMap := make(map[position]bool)

	for _, v := range b {
		blizzardMap[v.currentPosition] = true
	}

	// Do not Move
	if _, ok := blizzardMap[position{x: p.x, y: p.y}]; !ok {
		neigtbors = append(neigtbors, position{y: p.y, x: p.x, m: p.m + 1, start: p.start, end: p.end})
	}

	// end position
	if p.y+1 == se[1].y && p.x == se[1].x {
		neigtbors = append(neigtbors, position{y: p.y + 1, x: p.x, m: p.m + 1, start: p.start, end: p.end})
	}
	// start position
	if p.y-1 == se[0].y && p.x == se[0].x {
		neigtbors = append(neigtbors, position{y: p.y - 1, x: p.x, m: p.m + 1, start: p.start, end: p.end})
	}

	// left, up, right, down
	for _, v := range []direction{left, up, right, down} {
		newPos := v.newPos(p)
		if _, ok := blizzardMap[newPos]; newPos.y > 0 && newPos.y < height && newPos.x > 0 && newPos.x < width && !ok {
			newPos.m = p.m + 1
			newPos.start = p.start
			newPos.end = p.end
			neigtbors = append(neigtbors, newPos)
		}
	}

	return neigtbors
}

func occupied(p position, b []blizzard) bool {
	blizzardMap := make(map[position]bool)

	for _, v := range b {
		blizzardMap[v.currentPosition] = true
	}

	if _, ok := blizzardMap[p]; ok {
		return true
	}

	return false

}

func getBlizzardStates(b []blizzard, r int, w, h int) [][]blizzard {

	blizzardsStates := [][]blizzard{}
	state := make([]blizzard, len(b))
	copy(state, b)
	blizzardsStates = append(blizzardsStates, state)

	for i := 1; i < r; i++ {
		for i := range b {
			b[i].move(w, h)
		}
		state := make([]blizzard, len(b))
		copy(state, b)
		blizzardsStates = append(blizzardsStates, state)
	}

	return blizzardsStates
}

func Day24_1(input string) int {
	blizzards, widthHeight := parseInput(input)

	rounds := lcm(widthHeight[0]-1, widthHeight[1]-1)

	blizzardsStates := getBlizzardStates(blizzards, rounds, widthHeight[0], widthHeight[1])

	endPos := position{x: widthHeight[0] - 1, y: widthHeight[1]}
	queue := []position{{x: 1, y: 0, m: 0}}
	visit := make(map[position]bool)
	r := 0

	for {
		cp, rq := queue[0], queue[1:]
		queue = rq

		if _, ok := visit[cp]; ok {
			continue
		}
		visit[cp] = true

		if cp.x == endPos.x && cp.y == endPos.y {
			r = cp.m
			break
		}

		ns := getNeigtbors(cp, blizzardsStates[(cp.m+1)%rounds], []position{{x: 1, y: 0}, endPos}, widthHeight[0], widthHeight[1])
		for _, n := range ns {
			queue = append(queue, n)
		}
	}

	return r
}

func Day24_2(input string) int {
	blizzards, widthHeight := parseInput(input)

	rounds := lcm(widthHeight[0]-1, widthHeight[1]-1)
	blizzardsStates := getBlizzardStates(blizzards, rounds, widthHeight[0], widthHeight[1])
	endPos := position{x: widthHeight[0] - 1, y: widthHeight[1]}

	queue := []position{{x: 1, y: 0, m: 0}}

	visit := make(map[position]bool)

	r := 0

	for {
		cp, rq := queue[0], queue[1:]
		queue = rq

		if _, ok := visit[cp]; ok {
			continue
		}

		visit[cp] = true

		if cp.start && !cp.end {
			continue
		}

		// Check for end position
		if cp.x == endPos.x && cp.y == endPos.y {
			if cp.start && cp.end {
				r = cp.m
				break
			}
			cp.end = true
		}

		// Check if we are start and has been at end
		if cp.x == 1 && cp.y == 0 && cp.end {
			cp.start = true
		}

		ns := getNeigtbors(cp, blizzardsStates[(cp.m+1)%rounds], []position{{x: 1, y: 0}, endPos}, widthHeight[0], widthHeight[1])
		for _, n := range ns {
			queue = append(queue, n)
		}
	}

	return r
}
