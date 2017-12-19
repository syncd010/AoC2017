package main

import (
	"fmt"
	"os"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) (err error) {
	return nil
}

// Converts to []int
func convert(input []string) []string {
	return input
}

func solvePart1(input string) int {
	inGarbage, ignoreNext := false, false
	depth, score := 0, 0
	for _, c := range input {
		if inGarbage {
			switch {
			case ignoreNext:
				ignoreNext = false
			case c == '!':
				ignoreNext = true
			case c == '>':
				inGarbage = false
			}
			continue
		}

		switch c {
		case '<':
			inGarbage = true
		case '{':
			depth++
			score += depth
		case '}':
			depth--
		}
	}
	return score
}

func solvePart2(input string) int {
	inGarbage, ignoreNext := false, false
	n := 0
	for _, c := range input {
		if inGarbage {
			switch {
			case ignoreNext:
				ignoreNext = false
				continue
			case c == '!':
				ignoreNext = true
				continue
			case c == '>':
				inGarbage = false
				continue
			}
			n++
			continue
		}

		switch c {
		case '<':
			inGarbage = true
		}
	}
	return n
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(input[0]))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(input[0]))
}
