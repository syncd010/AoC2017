package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

func transpose(s string) string {
	sz := helpers.IntSqrt(len(s))

	t := make([]byte, sz*sz)
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			t[j*sz+i] = s[i*sz+j]
		}
	}
	return string(t)
}

func flipH(s string) string {
	sz := helpers.IntSqrt(len(s))

	t := make([]byte, sz*sz)
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			t[i*sz+j] = s[i*sz+sz-1-j]
		}
	}
	return string(t)
}

func flipV(s string) string {
	sz := helpers.IntSqrt(len(s))

	t := make([]byte, sz*sz)
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			t[i*sz+j] = s[(sz-1-i)*sz+j]
		}
	}
	return string(t)
}

func rotate(s string) string {
	return flipH(transpose(s))
}

// Converts to an appropria format
func convert(input []string) map[string]string {
	patterns := make(map[string]string)

	generatingOps := []func(string) string{
		flipH, rotate, flipV, rotate, flipH, rotate, flipV,
	}

	for _, line := range input {
		line = strings.Replace(line, "/", "", -1)
		rule := strings.Split(line, " => ")

		p := rule[0]
		patterns[p] = rule[1]
		for _, op := range generatingOps {
			p = op(p)
			patterns[p] = rule[1]
		}

	}
	return patterns
}

func solve(patterns map[string]string, steps int) int {
	grid := []byte(".#...####")

	var numSubgrids, szNewgrid, szSubgrid int
	var gridNew []byte
	for step := 0; step < steps; step++ {
		szGrid := helpers.IntSqrt(len(grid))

		switch {
		case szGrid%2 == 0:
			szSubgrid = 2
		case szGrid%3 == 0:
			szSubgrid = 3
		default:
			panic(errors.New("Board with wrong size"))
		}
		numSubgrids = szGrid / szSubgrid
		szNewgrid = numSubgrids * (szSubgrid + 1)
		gridNew = make([]byte, szNewgrid*szNewgrid)
		for i := 0; i < numSubgrids; i++ {
			for j := 0; j < numSubgrids; j++ {
				subgrid := make([]byte, szSubgrid*szSubgrid)
				for m := 0; m < szSubgrid; m++ {
					for n := 0; n < szSubgrid; n++ {
						idx := i*numSubgrids*szSubgrid*szSubgrid + j*szSubgrid + m*numSubgrids*szSubgrid + n
						subgrid[m*szSubgrid+n] = grid[idx]
					}
				}
				newSubgrid, ok := patterns[string(subgrid)]
				if !ok {
					panic(errors.New("Couldn't find rule for pattern: " + string(subgrid)))
				}
				// fmt.Println("Looking for: ", string(subgrid), " Found: ", string(newSubgrid))
				for m := 0; m < szSubgrid+1; m++ {
					for n := 0; n < szSubgrid+1; n++ {
						idx := i*numSubgrids*(szSubgrid+1)*(szSubgrid+1) + j*(szSubgrid+1) + m*numSubgrids*(szSubgrid+1) + n
						gridNew[idx] = newSubgrid[m*(szSubgrid+1)+n]
					}
				}

			}
		}
		grid = gridNew
	}

	on := 0
	for i := 0; i < len(grid); i++ {
		if grid[i] == '#' {
			on++
		}
	}

	return on
}

func solvePart1(input map[string]string) int {
	return solve(input, 5)
}

func solvePart2(input map[string]string) int {
	return solve(input, 18)
}

func printGrid(grid string) {
	sz := helpers.IntSqrt(len(grid))
	for i := 0; i < sz; i++ {
		fmt.Println(grid[i*sz : (i+1)*sz])
	}
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
