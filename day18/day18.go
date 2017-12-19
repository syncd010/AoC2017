package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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

func getValue(value string, registers map[string]int) int {
	x, err := strconv.Atoi(value)
	if err != nil {
		x = registers[value]
	}
	return x
}

func solvePart1(input []string) int {
	ip := 0
	registers := make(map[string]int)
	lastSound := 0
	var x, y int
	for ip < len(input) {
		inst := strings.Split(input[ip], " ")

		switch inst[0] {
		case "snd":
			lastSound = getValue(inst[1], registers)
		case "set":
			registers[inst[1]] = getValue(inst[2], registers)
		case "add":
			registers[inst[1]] += getValue(inst[2], registers)
		case "mul":
			registers[inst[1]] *= getValue(inst[2], registers)
		case "mod":
			registers[inst[1]] %= getValue(inst[2], registers)
		case "rcv":
			x = getValue(inst[1], registers)
			if x != 0 {
				return lastSound
			}
		case "jgz":
			x = getValue(inst[1], registers)
			y = getValue(inst[2], registers)
			if x > 0 {
				ip += y
				continue
			}
		}
		ip++
	}
	return lastSound
}

func execPart2(input []string, registers map[string]int, snd chan int, rcv chan int) int {
	ip := 0
	var x, y int
	count := 0

	for ip < len(input) {
		inst := strings.Split(input[ip], " ")

		switch inst[0] {
		case "snd":
			count++
			snd <- getValue(inst[1], registers)
		case "set":
			registers[inst[1]] = getValue(inst[2], registers)
		case "add":
			registers[inst[1]] += getValue(inst[2], registers)
		case "mul":
			registers[inst[1]] *= getValue(inst[2], registers)
		case "mod":
			registers[inst[1]] %= getValue(inst[2], registers)
		case "rcv":
			// Receive with timeout
			select {
			case registers[inst[1]] = <-rcv:
			case <-time.After(time.Second * 2):
				return count
			}
		case "jgz":
			x = getValue(inst[1], registers)
			y = getValue(inst[2], registers)
			if x > 0 {
				ip += y
				continue
			}
		}
		ip++
	}
	return count
}

func solvePart2(input []string) int {
	registers0 := make(map[string]int)
	registers0["p"] = 0
	registers1 := make(map[string]int)
	registers1["p"] = 1

	c01 := make(chan int, 1000)
	c10 := make(chan int, 1000)
	go execPart2(input, registers0, c01, c10)
	return execPart2(input, registers1, c10, c01)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
