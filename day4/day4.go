package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Nothing to do here
	return nil
}

// Converts to [][]string
func convert(input []string) [][]string {
	matrix := make([][]string, len(input))

	for i, line := range input {
		matrix[i] = strings.Fields(line)
	}
	return matrix
}

func solve(input [][]string, getKey func(string) string) int {
	sum := 0
	for _, row := range input {
		m := make(map[string]bool)
		valid := true
		for _, str := range row {
			key := getKey(str)
			_, ok := m[key]
			if ok {
				valid = false
				break
			}
			m[key] = true
		}
		if valid {
			sum++
		}
	}
	return sum
}

func solvePart1(input [][]string) int {
	return solve(input,
		func(input string) string {
			return input
		})
}

func solvePart2(input [][]string) int {
	return solve(input,
		func(input string) string {
			s := strings.Split(input, "")
			sort.Strings(s)
			return strings.Join(s, "")
		})
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	matrix := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(matrix))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(matrix))
}
