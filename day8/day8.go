package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Lazy, just trust the input...
	return nil
}

// Converts the input to an appropriate structure
func convert(input []string) []Instruction {
	instructions := make([]Instruction, len(input))

	for i, line := range input {
		instructions[i] = newInstruction(line)
	}

	return instructions
}

// Condition type
type Condition struct {
	reg   string
	op    string
	value int
}

// Instruction type
type Instruction struct {
	reg   string
	op    string
	value int
	cond  Condition
}

func newInstruction(line string) Instruction {
	re := regexp.MustCompile("[A-Za-z0-9<>!=-]+")
	stripped := re.FindAllString(line, -1)

	// b inc 5 if a > 1
	v, err := strconv.Atoi(stripped[6])
	helpers.Check(err, "Invalid number format")
	cond := Condition{reg: stripped[4], op: stripped[5], value: v}

	v, err = strconv.Atoi(stripped[2])
	helpers.Check(err, "Invalid number format")
	inst := Instruction{reg: stripped[0], op: stripped[1], value: v, cond: cond}
	return inst
}

func checkCond(cond Condition, regs map[string]int) bool {
	var op func(int, int) bool
	switch cond.op {
	case ">":
		op = func(a, b int) bool { return a > b }
	case "<":
		op = func(a, b int) bool { return a < b }
	case ">=":
		op = func(a, b int) bool { return a >= b }
	case "<=":
		op = func(a, b int) bool { return a <= b }
	case "==":
		op = func(a, b int) bool { return a == b }
	case "!=":
		op = func(a, b int) bool { return a != b }
	}
	return op(regs[cond.reg], cond.value)
}

func initializeRegisters(code []Instruction) map[string]int {
	regs := make(map[string]int)
	for _, inst := range code {
		if _, ok := regs[inst.reg]; !ok {
			regs[inst.reg] = 0
		}
		if _, ok := regs[inst.cond.reg]; !ok {
			regs[inst.cond.reg] = 0
		}
	}
	return regs
}

func exec(code []Instruction, regs map[string]int) (map[string]int, int) {
	max := 0
	for _, inst := range code {
		if checkCond(inst.cond, regs) {
			switch inst.op {
			case "inc":
				regs[inst.reg] += inst.value
			case "dec":
				regs[inst.reg] -= inst.value
			}
			newMax := sliceMax(regs)
			if newMax > max {
				max = newMax
			}
		}
	}
	return regs, max
}

func sliceMax(regs map[string]int) int {
	res := 0
	for _, v := range regs {
		if v > res {
			res = v
		}
	}
	return res
}

func solvePart1(code []Instruction) int {
	regs, _ := exec(code, initializeRegisters(code))
	return sliceMax(regs)
}

func solvePart2(code []Instruction) int {
	_, max := exec(code, initializeRegisters(code))
	return max
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	code := convert(input)
	//	printTree(tree, root, 0)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(code))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(code))

}
