package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	for _, line := range input {
		_, err := strconv.Atoi(line)
		if err != nil {
			return errors.New("Input must be only numbers")
		}
	}
	return nil
}

// Converts to []int
func convert(input []string) []int {
	numbers := make([]int, len(input))

	for i, line := range input {
		numbers[i], _ = strconv.Atoi(line)
	}
	return numbers
}

func solve(input []int, postIncrease func(int) int) int {
	numbers := make([]int, len(input))
	copy(numbers, input)

	step, jump := 0, 0
	for idx := 0; idx < len(numbers); idx += jump {
		jump = numbers[idx]
		numbers[idx] += postIncrease(jump)
		step++
	}
	return step
}

func solvePart1(input []int) int {
	return solve(input, func(_ int) int { return 1 })
}

func solvePart2(input []int) int {
	return solve(input, func(n int) int {
		if n >= 3 {
			return -1
		}
		return 1
	})
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	numbers := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(numbers))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(numbers))
}
