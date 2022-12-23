package day10

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Idle = iota
	Running
)

type (
	Instruction struct {
		Name   string
		Cycles uint
		Value  int
	}

	Cpu struct {
		Instructions *Instruction
		State        int
		X            int
		Cycle        int
	}
)

func CalcSignalStrength(c Cpu) int {
	return c.Cycle * c.X
}

func (c *Cpu) AddInstruction(i *Instruction) {
	c.Instructions = i
}

func (c *Cpu) RunCycle() {

	switch c.State {
	case Idle:
		c.State = Running
		c.Instructions.Cycles -= 1
	case Running:
		c.Instructions.Cycles -= 1
	}

	if c.Instructions.Cycles == 0 {
		if c.Instructions.Name != "noop" {
			c.X += c.Instructions.Value
		}
		c.State = Idle
		c.Instructions = nil
	}
}

func (c Cpu) String() string {
	return fmt.Sprintf("CPU State: %d Cycle: %d register value: %d", c.State, c.Cycle, c.X)
}

func newInstruction(s string) Instruction {
	f := strings.Fields(s)
	switch f[0] {
	case "noop":
		return Instruction{
			Name:   f[0],
			Cycles: 1,
		}
	case "addx":
		value, _ := strconv.Atoi(f[1])
		return Instruction{
			Name:   f[0],
			Cycles: 2,
			Value:  value,
		}
	}
	return Instruction{}
}

func Day10_part1(input string) int {

	sum := 0

	fields := strings.Split(input, "\n")

	cpu := Cpu{
		X: 1,
	}

	for {
		if len(fields[0]) == 0 && cpu.State == Idle {
			fields = fields[1:]
			break
		}
		if len(fields) == 0 && cpu.State == Idle {
			break
		}
		cpu.Cycle += 1

		if cpu.State == Idle {
			i := newInstruction(fields[0])
			cpu.AddInstruction(&i)
			fields = fields[1:]
		}
		if cpu.Cycle%40 == 20 {
			signalStrength := CalcSignalStrength(cpu)
			sum += signalStrength
		}
		cpu.RunCycle()
	}
	return sum
}

func Day10_part2(input string) int {

	sum := 0
	crtPos := 1

	fields := strings.Split(input, "\n")

	cpu := Cpu{
		X: 1,
	}

	t := cpu.X
	a := 0
	for {
		if len(fields[0]) == 0 && cpu.State == Idle {
			fields = fields[1:]
			break
		}
		if len(fields) == 0 && cpu.State == Idle {
			break
		}
		cpu.Cycle += 1
		if cpu.Cycle >= crtPos && cpu.Cycle <= crtPos+2 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if cpu.Cycle%40 == 0 {
			a += 40
			fmt.Print("\n")
		}

		if cpu.State == Idle {
			i := newInstruction(fields[0])
			cpu.AddInstruction(&i)
			fields = fields[1:]
		}
		cpu.RunCycle()
		if t != cpu.X {
			t = cpu.X
			crtPos = t + a
		}
	}

	return sum
}
