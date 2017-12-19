package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

// Converts to an appropriate format
func convert(input []string) []int {
	numbers := make([]int, len(input))

	var err error
	for i, line := range input {
		numbers[i], err = strconv.Atoi(line)
		helpers.Check(err, "Invalid number format")
	}
	return numbers
}

func generate(seed int, factor int, multiple int, iter int) []int {
	values := make([]int, iter)
	val := seed
	i := 0
	for i < iter {
		val = (val * factor) % 2147483647
		if val%multiple == 0 {
			values[i] = val
			i++
		}
	}
	return values
}

func countEqual(in1 []int, in2 []int) int {
	equal := 0
	for i := 0; i < helpers.Min(len(in1), len(in2)); i++ {
		if (in1[i] & 0xffff) == (in2[i] & 0xffff) {
			equal++
		}
	}
	return equal
}

func solvePart1(input []int) int {
	iter := 40000000
	generator1 := generate(input[0], 16807, 1, iter)
	generator2 := generate(input[1], 48271, 1, iter)

	return countEqual(generator1, generator2)
}

func solvePart2(input []int) int {
	iter := 5000000
	generator1 := generate(input[0], 16807, 4, iter)
	generator2 := generate(input[1], 48271, 8, iter)

	return countEqual(generator1, generator2)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
