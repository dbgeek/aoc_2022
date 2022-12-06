package day6

import (
	"math/big"
)

func foundMarker(s string) bool {
	var (
		bits big.Int
		val  int
	)

	for _, chr := range s {
		val = int(chr)
		if bits.Bit(val) != 0 {
			return false
		}
		bits.SetBit(&bits, int(chr), 1)
	}
	return true

}

func Day6_1(input string) int {
	const (
		message_length int = 4
	)

	for i := message_length; i < len(input); i++ {
		if foundMarker(input[i-message_length : i]) {
			return i
		}
	}

	return -1
}

func Day6_2(input string) int {
	const (
		message_length int = 14
	)

	for i := message_length; i < len(input); i++ {
		if foundMarker(input[i-message_length : i]) {
			return i
		}
	}

	return -1
}
