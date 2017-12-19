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
func validate(input []string) error {
	for _, line := range input {
		numbers := strings.Fields(line)
		for _, n := range numbers {
			_, err := strconv.Atoi(n)
			if err != nil {
				return errors.New("Input must be only numbers")
			}
		}
	}
	return nil
}

// Converts to [][]int
func convert(input []string) [][]int {
	matrix := make([][]int, len(input))

	for i, line := range input {
		numbers := strings.Fields(line)
		matrix[i] = make([]int, len(numbers))
		for j, n := range numbers {
			matrix[i][j], _ = strconv.Atoi(n)
		}
	}
	return matrix
}

func rowFuncPart1(row []int) int {
	if len(row) == 0 {
		return 0
	}
	min, max := helpers.MaxInt, helpers.MinInt
	for _, n := range row {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}
	return max - min
}

func rowFuncPart2(row []int) int {
	res := 0
	for i, n := range row {
		for _, m := range row[i+1:] {
			switch {
			case n >= m && (n%m) == 0:
				res = n / m
				break
			case n < m && (m%n) == 0:
				res = m / n
				break
			}
		}
		if res != 0 {
			break
		}
	}
	return res
}

func solve(input [][]int, rowFunc func([]int) int) int {
	sum := 0
	for _, row := range input {
		sum += rowFunc(row)
	}
	return sum
}

func solvePart1(input [][]int) int {
	return solve(input, rowFuncPart1)
}

func solvePart2(input [][]int) int {
	return solve(input, rowFuncPart2)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	matrix := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(matrix))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(matrix))
}
