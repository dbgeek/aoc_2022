package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Day1_1(input string) int64 {

	parsed := strings.Split(input, "\n")

	var max_calories int64
	var current_calories int64

	for _, v := range parsed {
		if v == "" {
			if current_calories > max_calories {
				max_calories = current_calories
			}
			current_calories = 0
			continue
		}

		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		current_calories += i
	}

	if current_calories > max_calories {
		max_calories = current_calories
		current_calories = 0
	}

	return max_calories
}

func Day1_2(input string) float64 {
	parsed := strings.Split(input, "\n")

	var elfs []float64
	var calories float64

	for _, v := range parsed {
		if v == "" {
			elfs = append(elfs, calories)
			calories = 0
		}

		if v != "" {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			calories += float64(i)
		}
	}

	sort.Float64s(elfs)
	elfs_len := len(elfs)
	return elfs[elfs_len-1] + elfs[elfs_len-2] + elfs[elfs_len-3]
}
