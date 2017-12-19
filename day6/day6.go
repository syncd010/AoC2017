package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input string) error {
	numbers := strings.Fields(input)
	for _, n := range numbers {
		_, err := strconv.Atoi(n)
		if err != nil {
			return errors.New("Input must be only numbers")
		}
	}
	return nil
}

// Converts to []int
func convert(input string) []int {
	numbersStr := strings.Fields(input)
	numbers := make([]int, len(numbersStr))

	for i, nStr := range numbersStr {
		numbers[i], _ = strconv.Atoi(nStr)
	}
	return numbers
}

func sliceMax(numbers []int) (int, int) {
	idx := 0
	for i, n := range numbers {
		if n > numbers[idx] {
			idx = i
		}
	}
	return idx, numbers[idx]
}

func sliceToString(numbers []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func solvePart1(input []int) int {
	// Make a copy, because we're going to modify it
	numbers := make([]int, len(input))
	copy(numbers, input)

	steps := 0
	m := make(map[string]bool)
	seen := false
	for !seen {
		idx, max := sliceMax(numbers)
		numbers[idx] = 0
		for i := 0; i < max; i++ {
			numbers[(idx+i+1)%len(numbers)]++
		}
		key := sliceToString(numbers)
		seen = m[key]
		m[key] = true
		steps++
	}

	return steps
}

func solvePart2(input []int) int {
	// Make a copy, because we're going to modify it
	numbers := make([]int, len(input))
	copy(numbers, input)

	m := make(map[string]int)
	steps := 0
	for {
		idx, max := sliceMax(numbers)
		numbers[idx] = 0
		for i := 0; i < max; i++ {
			numbers[(idx+i+1)%len(numbers)]++
		}
		key := sliceToString(numbers)
		prevStep, seen := m[key]
		if seen {
			return steps - prevStep
		}
		m[key] = steps
		steps++
	}
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input[0]), "Please provide a valid input")

	numbers := convert(input[0])

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(numbers))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(numbers))
}
