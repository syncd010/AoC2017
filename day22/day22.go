package main

import (
	"fmt"
	"os"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

// Converts to an appropria format
func convert(input []string) [][]byte {
	w := len(input)
	res := make([][]byte, w)
	for i := 0; i < w; i++ {
		res[i] = []byte(input[i])
	}
	return res
}

func make2DByteArray(w, h int, content byte) [][]byte {
	res := make([][]byte, h)
	for i := 0; i < h; i++ {
		res[i] = make([]byte, w)
		for j := range res[i] {
			res[i][j] = content
		}
	}
	return res
}

// Checks if the x,y in inside the grid.
// If not, allocates a new grid, copies the old one to the new and returns
func checkBounds(grid [][]byte, x, y int) (newGrid [][]byte, newX, newY int) {
	if (grid != nil) &&
		(x >= 0) && (y >= 0) &&
		(x < len(grid)) && (y < len(grid[0])) {
		return grid, x, y
	}

	// Make a new grid, copy the old one to the center
	szIni := len(grid)
	szGrid := 2*szIni + szIni%2
	newGrid = make2DByteArray(szGrid, szGrid, '.')

	for i, row := range grid {
		for j, c := range row {
			newGrid[(szGrid-szIni)/2+i][(szGrid-szIni)/2+j] = byte(c)
		}
	}

	newX, newY = (szGrid-szIni)/2+x, (szGrid-szIni)/2+y
	return newGrid, newX, newY
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func move(x, y, dir int) (newX, newY int) {
	switch dir {
	case UP:
		y--
	case RIGHT:
		x++
	case DOWN:
		y++
	case LEFT:
		x--
	}
	return x, y
}

func simulate(ini [][]byte, steps int, stepFn func([][]byte, int, int, int) (int, int)) ([][]byte, int) {
	grid, _, _ := checkBounds([][]byte(ini), len(ini), len(ini))

	x, y := len(grid)/2, len(grid)/2
	dir := UP
	infected, infections := 0, 0
	for i := 0; i < steps; i++ {
		grid, x, y = checkBounds(grid, x, y)
		dir, infected = stepFn(grid, x, y, dir)
		infections += infected
		x, y = move(x, y, dir)
	}
	return grid, infections
}

func stepPart1(grid [][]byte, x, y int, dir int) (newDir int, infected int) {
	switch grid[y][x] {
	case '#':
		grid[y][x] = '.'
		newDir = (dir + 1) % 4
	case '.':
		grid[y][x] = '#'
		newDir = (dir + 3) % 4
		infected++
	}
	return
}

func stepPart2(grid [][]byte, x, y int, dir int) (newDir int, infected int) {
	switch grid[y][x] {
	case '#':
		grid[y][x] = 'F'
		newDir = (dir + 1) % 4
	case 'F':
		grid[y][x] = '.'
		newDir = (dir + 2) % 4
	case 'W':
		grid[y][x] = '#'
		newDir = dir
		infected++
	case '.':
		grid[y][x] = 'W'
		newDir = (dir + 3) % 4
	}
	return
}

func solvePart1(input [][]byte) int {
	_, infected := simulate(input, 10000, stepPart1)
	return infected
}

func solvePart2(input [][]byte) int {
	_, infected := simulate(input, 10000000, stepPart2)
	return infected
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
