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
	muls := 0
	var x, y int
	for ip < len(input) {
		inst := strings.Split(input[ip], " ")

		switch inst[0] {
		case "set":
			registers[inst[1]] = getValue(inst[2], registers)
		case "sub":
			registers[inst[1]] -= getValue(inst[2], registers)
		case "mul":
			registers[inst[1]] *= getValue(inst[2], registers)
			muls++
		case "jnz":
			x = getValue(inst[1], registers)
			y = getValue(inst[2], registers)
			if x != 0 {
				ip += y
				continue
			}
		}
		ip++
	}
	return muls
}

func solvePart2(input []string) int {
	var b, c, d, e, h int

	b = 99*100 + 100000
	c = b + 17000

	for ; b <= c; b += 17 {
	loop3:
		for d = 2; d != b; d++ {
		loop2:
			for e = 2; e != b; e++ {
				if d*e > b { // Optimization added
					break loop2 // Innermost loop, unnecessary but consistent with the rest
				}
				if d*e == b {
					h++
					break loop3 // Optimization added
				}
			}
		}
	}

	return h
}

func inputTranslated() int {
	var a, b, c, d, e, f, g, h int

	a = 1
	b = 99      // set b 99
	c = b       // set c b
	if a != 0 { // jnz a 2
		b = b*100 + 100000 // mul b 100, sub b -100000
		c = b + 17000      // set c b, sub c -17000
	} // jnz 1 5

loop3:
	f = 1 // set f 1
	d = 2 // set d 2
loop2:
	e = 2 // set e 2
loop1:
	g = d*e - b // set g d, mul g e, sub g b
	if g == 0 { // jnz g 2
		f = 0 // set f 0
	}
	e++         // sub e -1
	g = e - b   // set g e, sub g b
	if g != 0 { // jnz g -8
		goto loop1
	}
	d++         // sub d -1
	g = d - b   // set g d, sub g b
	if g != 0 { // jnz g -13
		goto loop2
	}
	if f == 0 { // jnz f 2
		h++ // sub h -1
	}
	g = b - c   // set g b, sub g c
	if g != 0 { // jnz g 2
		b += 17    // sub b -17
		goto loop3 // jnz 1 -23
	}
	//jnz 1 3

	return h
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
