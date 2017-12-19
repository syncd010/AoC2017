package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

// Converts to an appropria format
func convert(input []string) []string {
	return strings.Split(input[0], ",")
}

func findIdx(what byte, where []byte) int {
	for i, c := range where {
		if c == what {
			return i
		}
	}
	return -1
}

var aux []byte

func dance(input []string, order []byte) []byte {
	if aux == nil || len(aux) != len(order) {
		aux = make([]byte, len(order))
	}
	for _, step := range input {
		switch step[0] {
		case 's':
			n, _ := strconv.Atoi(step[1:])
			copy(aux[:n], order[len(order)-n:])
			copy(order[n:], order[:len(order)-n])
			copy(order[:n], aux[:n])
		case 'x':
			who := strings.Split(step[1:], "/")
			a, _ := strconv.Atoi(who[0])
			b, _ := strconv.Atoi(who[1])
			order[a], order[b] = order[b], order[a]
		case 'p':
			who := strings.Split(step[1:], "/")
			a := findIdx(who[0][0], order)
			b := findIdx(who[1][0], order)
			order[a], order[b] = order[b], order[a]
		}
	}
	return order
}

func solvePart1(input []string) string {
	order := []byte("abcdefghijklmnop")
	return string(dance(input, order))
}

func solvePart2(input []string) string {
	order := []byte("abcdefghijklmnop")
	memo := make(map[string]string)

	iter := 1000000000
	for i := 0; i < iter; i++ {
		memoKey := string(order)
		if nextOrder, ok := memo[memoKey]; ok {
			// Cycle, advance
			i = int(iter/i) * i
			order = []byte(nextOrder)
			continue
		}
		order = dance(input, order)
		memo[memoKey] = string(order)
	}
	return string(order)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
