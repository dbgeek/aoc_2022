package day4

import (
	"math"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == '-' || r == ','
}

func Day4_1(input string) float64 {

	parsed := strings.Fields(input)

	counter := 0

	for _, v := range parsed {
		elfs := strings.FieldsFunc(v, split)

		left_low_int, _ := strconv.Atoi(elfs[0])
		right_low_int, _ := strconv.Atoi(elfs[2])

		left_high_int, _ := strconv.Atoi(elfs[1])
		right_high_int, _ := strconv.Atoi(elfs[3])

		if right_low_int >= left_low_int && right_high_int <= left_high_int || left_low_int >= right_low_int && left_high_int <= right_high_int {
			counter += 1
		}
	}

	return float64(counter)
}

func Day4_2(input string) float64 {

	parsed := strings.Fields(input)

	counter := 0

	for _, v := range parsed {
		elfs := strings.FieldsFunc(v, split)

		left_low_int, _ := strconv.Atoi(elfs[0])
		right_low_int, _ := strconv.Atoi(elfs[2])

		left_high_int, _ := strconv.Atoi(elfs[1])
		right_high_int, _ := strconv.Atoi(elfs[3])

		if math.Max(float64(left_low_int), float64(right_low_int))-math.Min(float64(left_high_int), float64(right_high_int)) <= 0 {
			counter += 1
		}
	}
	return float64(counter)

}
