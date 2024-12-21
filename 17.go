package main

import (
	"fmt"
	"math"
	"regexp"
)

type Computer struct {
	A, B, C int
	program []int
	output  []int
	ip      int
}

func (c *Computer) run() {
	for ; c.ip < len(c.program); c.ip += 2 {
		opcode := c.program[c.ip]
		operand := c.program[c.ip+1]
		switch opcode {
		case 0:
			c.A = c.div(operand)
		case 1:
			c.B = c.xor(operand)
		case 2:
			c.B = c.mod8(operand)
		case 3:
			c.jump(operand)
		case 4:
			c.B = c.xor(c.C)
		case 5:
			c.output = append(c.output, c.mod8(operand))
		case 6:
			c.B = c.div(operand)
		case 7:
			c.C = c.div(operand)
		}
	}
}

func (c Computer) div(operand int) int {
	return c.A / int(math.Pow(2, float64(c.combo(operand))))
}

func (c Computer) xor(operand int) int {
	return c.B ^ operand
}

func (c Computer) mod8(operand int) int {
	return c.combo(operand) % 8
}

func (c *Computer) jump(operand int) {
	if c.A != 0 {
		c.ip = operand - 2
	}
}

func (c Computer) combo(operand int) int {
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	}
	return operand
}

func Day17() {
	lines := getInputString(17)
	// lines := getInputString(0)
	computer := getComputer(lines)
	dfs(computer, int(math.Pow(8, float64(len(computer.program))-1)), len(computer.program)-1)
	computer.run()
	for _, out := range computer.output {
		fmt.Print(out, ",")
	}
	fmt.Println()
}

func dfs(c Computer, a, exp int) bool {
	if exp < 0 {
		fmt.Println("Found A:", a)
		return true
	}

	for i := 0; i < 8; i++ {
		tmpc := c
		tmpa := a + int(math.Pow(8, float64(exp)))*i
		tmpc.A = tmpa
		tmpc.run()
		if tmpc.output[exp] == c.program[exp] {
			if dfs(c, tmpa, exp-1) {
				return true
			}
		}
	}
	return false
}

func getComputer(lines string) Computer {
	num := regexp.MustCompile(`\d+`)
	args := num.FindAllString(lines, -1)
	intargs := make([]int, len(args))
	for i, arg := range args {
		intargs[i] = strToInt(arg)
	}
	return Computer{intargs[0], intargs[1], intargs[2], intargs[3:], []int{}, 0}
}
