package day5

import (
	"strconv"
	"strings"
)

type (
	crates struct {
		stack []string
	}
)

func (stack *crates) peak() string {
	return stack.stack[len(stack.stack)-1]
}

func (stack *crates) pop() string {
	out := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return out
}

func (stack *crates) push(s string) {
	stack.stack = append(stack.stack, s)
}

func (stack *crates) insert(s string, p int) {
	stack.stack = append(stack.stack[:p], append([]string{s}, stack.stack[p:]...)...)
}

func newCreate() *crates {
	return &crates{}
}


func Day5_1(input string) string {
	parsed := strings.Split(input, "\n\n")
	stacks := strings.Split(parsed[0], "\n")
	moves := strings.Split(parsed[1], "\n")

	crate_stacks := []crates{}

	for idx, line := range stacks {
		if idx == len(stacks)-1 {
			continue
		}
		stack := 0
		for i, v := range line {
			if i%4 == 1 {
				stack++
				if len(crate_stacks) < stack {
					crate_stacks = append(crate_stacks, *newCreate())
				}
				if string(v) != " " {
					crate_stacks[stack-1].insert(string(v), 0)
				}
			}
		}
	}
	for _, v := range moves {
		move := strings.Fields(v)
		if len(move) == 0 {
			continue
		}

		nr_moves, _ := strconv.Atoi(move[1])
		for i := 0; i < nr_moves; i++ {
			src, _ := strconv.Atoi(move[3])
			dst, _ := strconv.Atoi(move[5])
			crate := crate_stacks[src-1].pop()
			crate_stacks[dst-1].push(crate)
		}
	}
	r := []string{}

	for _, v := range crate_stacks {
		r = append(r, v.peak())
	}
	return strings.Join(r, "")
}

func Day5_2(input string) string {
	parsed := strings.Split(input, "\n\n")
	stacks := strings.Split(parsed[0], "\n")
	moves := strings.Split(parsed[1], "\n")

	crate_stacks := []crates{}

	for idx, line := range stacks {
		if idx == len(stacks)-1 {
			continue
		}
		stack := 0
		for i, v := range line {
			if i%4 == 1 {
				stack++
				if len(crate_stacks) < stack {
					crate_stacks = append(crate_stacks, *newCreate())
				}
				if string(v) != " " {
					crate_stacks[stack-1].insert(string(v), 0)
				}
			}
		}
	}
	for _, v := range moves {
		move := strings.Fields(v)
		if len(move) == 0 {
			continue
		}

		nr_moves, _ := strconv.Atoi(move[1])
		crates := []string{}
		for i := 0; i < nr_moves; i++ {
			src, _ := strconv.Atoi(move[3])
			crate := crate_stacks[src-1].pop()
			crates = append(crates, crate)
		}
		dst, _ := strconv.Atoi(move[5])
		for i := len(crates) - 1; i >= 0; i-- {
			crate_stacks[dst-1].push(crates[i])
		}
	}
	r := []string{}

	for _, v := range crate_stacks {
		r = append(r, v.peak())
	}
	return strings.Join(r, "")
}
