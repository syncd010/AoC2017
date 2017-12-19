package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input string) error {
	if len(input) < 1 {
		return errors.New("Input must not be empty")
	}

	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Input must be a number")
	}

	return nil
}

// Converts to [][]int
func convert(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

// Returns a path from the center of the board to the requested input
func getPath(input int) []byte {
	dirs := []byte{'R', 'U', 'L', 'D'}
	currDir := 0
	targetLength := 1
	currLength := 0

	path := make([]byte, input-1)
	// Make something like RULLDDRRRUUULLLLDDDD
	for i := 1; i < input; i++ {
		path[i-1] = dirs[currDir]

		// Change dir if we've reached the desired length
		currLength++
		if currLength == targetLength {
			currLength = 0
			currDir = (currDir + 1) % len(dirs)
			// If changing dir to horizontal, increase desire lenght
			if dirs[currDir] == 'L' || dirs[currDir] == 'R' {
				targetLength++
			}
		}
	}
	return path
}

// Solve the problem
func solvePart1(input int) int {
	path := getPath(input)

	// Calculate Manhatan distance
	distH, distV := 0, 0
	for _, b := range path {
		switch b {
		case 'R':
			distH++
		case 'L':
			distH--
		case 'U':
			distV++
		case 'D':
			distV--
		}
	}

	return helpers.Abs(distH) + helpers.Abs(distV)
}

// Solve the problem
func solvePart2(input int) int {
	// Lets define a board big enough to contain the input number.
	// Highly inneficient but i'm in a hurry
	sz := int(math.Sqrt(float64(input))) + 4
	board := make([]int, sz*sz)

	x, y := sz/2, sz/2
	path := getPath(input)

	board[x*sz+y] = 1
	for _, b := range path {
		switch b {
		case 'R':
			y++
		case 'L':
			y--
		case 'U':
			x--
		case 'D':
			x++
		}
		board[x*sz+y] = board[(x*sz)+y-1] +
			board[(x*sz)+y+1] +
			board[((x-1)*sz)+y] +
			board[((x-1)*sz)+y-1] +
			board[((x-1)*sz)+y+1] +
			board[((x+1)*sz)+y] +
			board[((x+1)*sz)+y-1] +
			board[((x+1)*sz)+y+1]
		// printBoard(sz, board)
		if board[x*sz+y] >= input {
			return board[x*sz+y]
		}
	}

	return 0
}

func printBoard(sz int, board []int) {
	for x := 0; x < sz; x++ {
		for y := 0; y < sz; y++ {
			fmt.Printf(" %3v ", board[x*sz+y])
		}
		fmt.Printf("\n")
	}
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input[0]), "Please provide a valid input")

	num := convert(input[0])

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(num))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(num))
}
