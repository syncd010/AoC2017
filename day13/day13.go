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
func validate(input []string) (err error) {
	for _, line := range input {
		data := strings.Split(line, ":")
		if len(data) < 2 {
			return errors.New("Incomplete input: " + line)
		}
		_, err = strconv.Atoi(strings.TrimSpace(data[0]))
		if err != nil {
			return err
		}
		_, err = strconv.Atoi(strings.TrimSpace(data[1]))
		if err != nil {
			return err
		}
	}
	return nil
}

// Converts to []int
func convert(input []string) []int {
	depths := make([]int, len(input))
	widths := make([]int, len(input))

	for i, line := range input {
		data := strings.Split(line, ":")
		depths[i], _ = strconv.Atoi(strings.TrimSpace(data[0]))
		widths[i], _ = strconv.Atoi(strings.TrimSpace(data[1]))
	}

	firewall := make([]int, helpers.Max(depths...)+1)

	for i, depth := range depths {
		firewall[depth] = widths[i]
	}

	return firewall
}

// Returns the severity or -1 if not caught
func getSeverity(firewall []int, delay int) int {
	severity := 0
	caught := false
	for layer, width := range firewall {
		if width == 0 {
			continue
		}
		timePassed := delay + layer
		scannerPos := timePassed % ((width - 1) * 2)
		if scannerPos == 0 {
			// fmt.Println("Caught in layer: ", layer, " TimePassed: ", timePassed, " ScannerPos: ", scannerPos)
			severity += layer * width
			caught = true
		}
	}

	if !caught {
		return -1
	}
	return severity
}

func solvePart1(firewall []int) int {
	return getSeverity(firewall, 0)
}

func solvePart2(input []int) int {
	// Even though its not stated in the problem, make sure we have a max delay
	// to prevent infinite loops
	maxDelay := 100000000
	for delay := 0; delay < maxDelay; delay++ {
		severity := getSeverity(input, delay)
		if severity == -1 {
			return delay
		}
	}
	return -1
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	layers := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(layers))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(layers))
}
