package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	if len(input[0]) < 1 {
		return errors.New("Input must not be empty")
	}
	// Assume everything else is ok... lazy
	return nil
}

func convert(input []string) []string {
	return strings.Split(input[0], ",")
}

// Position2D represents a position on the board
type Position2D struct {
	x, y float64
}

func (p *Position2D) move(dir string) {
	switch dir {
	case "n":
		p.y += 1.0
	case "s":
		p.y -= 1.0
	case "ne":
		p.y += .5
		p.x += .5
	case "nw":
		p.y += .5
		p.x -= .5
	case "se":
		p.y -= .5
		p.x += .5
	case "sw":
		p.y -= .5
		p.x -= .5
	}
}

// Returns the distance from origin of the current position
func (p *Position2D) distanceFromOrigin() int {
	return int(math.Abs(p.x) + math.Abs(p.y))
}

func solvePart1(directions []string) int {
	pos := Position2D{0.0, 0.0}
	for _, dir := range directions {
		pos.move(dir)
	}
	return pos.distanceFromOrigin()
}

func solvePart2(directions []string) int {
	pos := Position2D{0.0, 0.0}
	maxDist := 0
	for _, dir := range directions {
		pos.move(dir)
		maxDist = helpers.Max(pos.distanceFromOrigin(), maxDist)
	}
	return maxDist
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	directions := convert(input)
	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(directions))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(directions))

}
