package day3

import (
	"strings"
)

func Day3_1(input string) float64 {

	parsed := strings.Fields(input)

	items := []rune{}

	for _, row := range parsed {
		set := make(map[rune]bool)
		left := row[0 : len(row)/2]
		right := row[len(row)/2:]
		for _, l := range left {
			if set[l] {
				continue
			}
			for _, r := range right {
				if l == r {
					set[l] = true
					items = append(items, l)
					break
				}
			}
		}
	}

	var sum float64
	for _, v := range items {
		if int(v) >= 97 && int(v) <= 122 {
			sum += float64(1 + (int(v) - 97))
		} else if int(v) >= 65 && int(v) <= 90 {
			sum += float64(27 + (int(v) - 65))
		}
	}

	return sum
}

func Day3_2(input string) float64 {

	parsed := strings.Fields(input)

	items := []rune{}

	for i := 0; i < len(parsed); i += 3 {
		set := make(map[rune]bool)
		row1 := parsed[i+2]
		row2 := parsed[i+1]
		row3 := parsed[i]

		for _, x := range row1 {
			if set[x] {
				continue
			}
			for _, y := range row2 {
				if x == y && !set[x] {
					for _, z := range row3 {
						if y == z {
							items = append(items, z)
							break
						}
					}
					set[y] = true
				}
			}
		}
	}
	var sum float64
	for _, v := range items {
		if int(v) >= 97 && int(v) <= 122 {
			sum += float64(1 + (int(v) - 97))
		} else if int(v) >= 65 && int(v) <= 90 {
			sum += float64(27 + (int(v) - 65))
		}
	}

	return sum
}
