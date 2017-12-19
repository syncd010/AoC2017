package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input string) error {
	if len(input) < 1 {
		return errors.New("Input must not be empty")
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return errors.New("Input must be only numbers")
		}
	}
	return nil
}

// Convert to int
func convert(input string) []int {
	numbers := make([]int, len(input))
	for i, c := range input {
		numbers[i] = int(c - '0')
	}
	return numbers
}

func sumInDistance(numbers []int, dist int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numbers[(i+dist)%len(numbers)] {
			sum += numbers[i]
		}
	}
	return sum
}

func solvePart1(numbers []int) int {
	return sumInDistance(numbers, 1)
}

func solvePart2(numbers []int) int {
	return sumInDistance(numbers, len(numbers)/2)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input[0]), "Please provide a valid input")

	numbers := convert(input[0])

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(numbers))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(numbers))

}
