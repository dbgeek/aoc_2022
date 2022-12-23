package day15

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type (
	vector struct {
		x int
		y int
	}
)

func (v *vector) distance(d *vector) int {
	return int(math.Abs(float64(v.x-d.x)) + math.Abs(float64(v.y-d.y)))

}

func manX(i [][2]int) int {
	maxX := math.MinInt

	for _, v := range i {
		if v[1] > maxX {
			maxX = v[1]
		}
	}
	return maxX
}

func minX(i [][2]int) int {
	minX := math.MaxInt

	for _, v := range i {
		if v[0] < minX {
			minX = v[0]
		}
	}
	return minX
}

func parse(s string) ([]*vector, []*vector) {
	sensors := []*vector{}
	beacons := []*vector{}

	r, _ := regexp.Compile("(x=-?[0-9]+)|(y=-?[0-9]+)")

	for _, row := range strings.Split(s, "\n") {
		if row == "" {
			continue
		}

		p := r.FindAllString(row, -1)
		sX, _ := strconv.Atoi(strings.Split(p[0], "=")[1])
		sY, _ := strconv.Atoi(strings.Split(p[1], "=")[1])
		sensor := &vector{x: sX, y: sY}

		bX, _ := strconv.Atoi(strings.Split(p[2], "=")[1])
		bY, _ := strconv.Atoi(strings.Split(p[3], "=")[1])
		beacon := &vector{x: bX, y: bY}

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	return sensors, beacons

}

func Day15_1(input string, rowY int) int {
	s, b := parse(input)

	sensorIntervalX := [][2]int{}
	// Calculate min / max x for each sensor
	for i, v := range s {
		// Calculate sensor x distance a rowY
		sdx := v.distance(b[i]) - int(math.Abs(float64(v.y-rowY)))
		// Filter out sensors that now covers rowY positions
		if sdx <= 0 {
			continue
		}
		// Calculate min and max x positions at rowY
		sensorMinXpos, sensorMaxXpos := v.x-sdx, v.x+sdx
		sensorIntervalX = append(sensorIntervalX, [2]int{sensorMinXpos, sensorMaxXpos})
	}

	// Get becon x positions at rowY
	beaconX := make(map[int]bool)
	for _, v := range b {
		if v.y == rowY {
			beaconX[v.x] = true
		}
	}

	minX := minX(sensorIntervalX)
	maxX := manX(sensorIntervalX)

	positions := 0
	for x := minX; x <= maxX; x++ {
		// x that has a beacon should not counted
		if _, ok := beaconX[x]; ok {
			continue
		}

		// Check if x positions exists in sensor intervals x positions
		for _, si := range sensorIntervalX {
			if si[0] <= x && si[1] >= x {
				positions += 1
				break
			}
		}
	}
	return positions
}

func Day15_2(input string) int {
	return -1
}
