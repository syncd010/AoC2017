package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

type Port struct {
	pins []int
	used bool
}

func (port Port) hasPin(pin int) bool {
	for _, p := range port.pins {
		if p == pin {
			return true
		}
	}
	return false
}

func (port Port) strength() int {
	st := 0
	for _, p := range port.pins {
		st += p
	}
	return st
}

func (port Port) nextPin(pin int) int {
	for i, p := range port.pins {
		if p == pin {
			return port.pins[(i+1)%len(port.pins)]
		}
	}
	return 0
}

// Converts to an appropriate format
func convert(input []string) []*Port {
	ports := make([]*Port, len(input))
	for i, line := range input {
		nums := strings.Split(line, "/")
		ports[i] = &Port{
			pins: []int{helpers.Atoi(nums[0]), helpers.Atoi(nums[1])},
			used: false}
	}
	return ports
}

func getMaxStrength(startPin int, ports []*Port) int {
	maxStrength := 0
	for _, end := range ports {
		if !end.used && end.hasPin(startPin) {
			end.used = true
			strength := end.strength() + getMaxStrength(end.nextPin(startPin), ports)
			if strength > maxStrength {
				maxStrength = strength
			}
			end.used = false
		}
	}
	return maxStrength
}

func getMaxLengthStrength(startPin int, ports []*Port) (maxLength int, maxStrength int) {
	for _, end := range ports {
		if !end.used && end.hasPin(startPin) {
			end.used = true

			length, strength := getMaxLengthStrength(end.nextPin(startPin), ports)

			if (length + 1) > maxLength {
				maxLength = length + 1
				maxStrength = strength + end.strength()
			} else if (length+1) == maxLength &&
				strength+end.strength() > maxStrength {
				maxStrength = strength + end.strength()
			}
			end.used = false
		}
	}
	return maxLength, maxStrength
}

func solvePart1(input []*Port) int {
	return getMaxStrength(0, input)
}

func solvePart2(input []*Port) int {
	_, s := getMaxLengthStrength(0, input)
	return s
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
