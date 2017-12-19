package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

// Converts to an appropria format
func convert(input []string) []string {
	return input
}

const UP, DOWN, LEFT, RIGHT = 'U', 'D', 'L', 'R'

type Position struct {
	x, y int
	dir  byte
}

func (pos Position) move(dir byte) (newPos Position) {
	newPos = pos
	switch dir {
	case DOWN:
		newPos.y++
	case UP:
		newPos.y--
	case RIGHT:
		newPos.x++
	case LEFT:
		newPos.x--
	}
	return newPos
}

func findStartingPosition(input []string) (pos Position) {
	pos.dir = DOWN
	pos.y = 0
	for i, c := range input[0] {
		if c == '|' {
			pos.x = i
			return pos
		}
	}
	return pos
}

func inBounds(pos Position, lenX, lenY int) bool {
	if (pos.x >= 0) && (pos.x < lenX) &&
		(pos.y >= 0) && (pos.y < lenY) {
		return true
	}
	return false
}

func getCell(input []string, pos Position) (in bool, cell byte) {
	lenX, lenY := len(input[0]), len(input)
	if !inBounds(pos, lenX, lenY) {
		return false, 0
	}
	return true, input[pos.y][pos.x]
}

func changeDirection(input []string, currPos Position) (pos Position) {
	directions := map[byte][]byte{
		UP:    []byte{LEFT, RIGHT},
		DOWN:  []byte{LEFT, RIGHT},
		LEFT:  []byte{UP, DOWN},
		RIGHT: []byte{UP, DOWN},
	}

	for _, dir := range directions[currPos.dir] {
		newPos := currPos.move(dir)
		if in, c := getCell(input, newPos); in && (c != ' ') {
			currPos.dir = dir
			return currPos
		}
	}

	return currPos
}

func solve(input []string) (string, int) {
	currPos := findStartingPosition(input)
	var letters []byte
	steps := 0

	for in, c := getCell(input, currPos); in && (c != ' '); in, c = getCell(input, currPos) {
		switch {
		case c == '+':
			currPos = changeDirection(input, currPos)
		case unicode.IsLetter(rune(c)):
			letters = append(letters, c)
		}

		currPos = currPos.move(currPos.dir)
		steps++
	}

	return string(letters), steps
}

func solvePart1(input []string) string {
	letters, _ := solve(input)
	return letters
}

func solvePart2(input []string) int {
	_, steps := solve(input)
	return steps
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
