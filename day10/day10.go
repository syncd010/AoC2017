package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	in := input[0]
	if len(in) < 1 {
		return errors.New("Input must not be empty")
	}

	for _, c := range in {
		if c != ' ' && c != ',' && (c < '0' || c > '9') {
			return errors.New("Input must be only numbers")
		}
	}
	return nil
}

func convertPart1(input []string) []int {
	numbersStr := strings.Split(input[0], ",")
	numbers := make([]int, len(numbersStr))

	for i, nStr := range numbersStr {
		numbers[i], _ = strconv.Atoi(nStr)
	}
	return numbers
}

func convertPart2(input []string) []int {
	numbers := make([]int, len(input[0]))

	for i, c := range input[0] {
		numbers[i] = int(c)
	}
	return numbers
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

// First part, sum
func solvePart1(lengths []int) int {
	buffer, bufferAux := createBuffers()

	skip := 0
	currPos := 0
	for _, l := range lengths {
		buffer = invertSlice(buffer, currPos, l, bufferAux)
		currPos = (currPos + l + skip) % len(buffer)
		skip++
	}
	return buffer[0] * buffer[1]
}

// Second part
func solvePart2(lengths []int) string {
	buffer, bufferAux := createBuffers()

	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

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
		knotHash += fmt.Sprintf("%02x", denseHash[i])
	}

	return knotHash
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertPart1(input)))

	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertPart2(input)))

}
