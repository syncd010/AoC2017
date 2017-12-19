package main

import (
	"fmt"
	"math"
	"os"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

func invertSlice(buffer []int, startPos int, length int, bufferAux []int) []int {
	// Invert
	for i := 0; i < length; i++ {
		pos := (startPos + i) % len(buffer)
		bufferAux[length-i-1] = buffer[pos]
	}
	// Copy back
	for i := 0; i < length; i++ {
		pos := (startPos + i) % len(buffer)
		buffer[pos] = bufferAux[i]
	}
	return buffer
}

func createBuffers() (buffer []int, bufferAux []int) {
	// Create the list
	buffer = make([]int, 256)
	for i := range buffer {
		buffer[i] = i
	}
	// Create the auxiliary buffer
	bufferAux = make([]int, len(buffer))
	return buffer, bufferAux
}

// Get a knot hash. Format should be "%02x" or "%08b"
func getKnotHash(input string, format string) string {
	lengths := make([]int, len(input))

	for i, c := range input {
		lengths[i] = int(c)
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	buffer, bufferAux := createBuffers()

	skip := 0
	currPos := 0
	for i := 0; i < 64; i++ {
		for _, l := range lengths {
			buffer = invertSlice(buffer, currPos, l, bufferAux)
			currPos = (currPos + l + skip) % len(buffer)
			skip++
		}
	}
	// Xor
	sz := int(math.Sqrt(float64(len(buffer))))
	denseHash := make([]int, sz)
	for i := 0; i < sz; i++ {
		denseHash[i] = buffer[i*sz]
		for j := 1; j < sz; j++ {
			denseHash[i] ^= buffer[i*sz+j]
		}
	}

	knotHash := ""
	for i := 0; i < sz; i++ {
		knotHash += fmt.Sprintf(format, denseHash[i])
	}

	return knotHash
}

// SIZE of the grid
const SIZE = 128

func getHashes(input string) []string {
	hashes := make([]string, SIZE)
	for i := 0; i < SIZE; i++ {
		hashes[i] = getKnotHash(input+fmt.Sprintf("-%v", i), "%08b")
	}
	return hashes
}

func solvePart1(input string) int {
	hashes := getHashes(input)
	res := 0
	for i := 0; i < SIZE; i++ {
		for _, c := range hashes[i] {
			if c == '1' {
				res++
			}
		}
	}
	return res
}

func markRegion(regions [][]int, hashes []string, i int, j int, id int) {
	regions[i][j] = id

	dirs := []struct {
		y, x int
	}{
		{y: i - 1, x: j},
		{y: i + 1, x: j},
		{y: i, x: j - 1},
		{y: i, x: j + 1},
	}

	for _, dir := range dirs {
		if (dir.x >= 0) && (dir.x < SIZE) &&
			(dir.y >= 0) && (dir.y < SIZE) &&
			(regions[dir.y][dir.x] == 0) &&
			(hashes[dir.y][dir.x] == '1') {
			markRegion(regions, hashes, dir.y, dir.x, id)
		}
	}
}

func solvePart2(input string) int {
	hashes := getHashes(input)
	regions := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		regions[i] = make([]int, SIZE)
	}

	maxRegion := 0
	for i := 0; i < SIZE; i++ {
		for j, c := range hashes[i] {
			if (c == '1') && (regions[i][j] == 0) {
				maxRegion++
				markRegion(regions, hashes, i, j, maxRegion)
			}
		}
	}

	return maxRegion
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(input[0]))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(input[0]))

}
