package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

// Converts to an appropria format
func convert(input []string) int {
	n, err := strconv.Atoi(input[0])
	helpers.Check(err, "Invalid number")
	return n
}

// Solution with Arrays
func spinArray(input int, iter int) []int {
	buf := make([]int, iter)
	currPos := 0
	for i := 1; i < iter; i++ {
		currPos = ((input%i)+currPos)%i + 1
		for j := i; j > currPos; j-- {
			buf[j] = buf[j-1]
		}
		buf[currPos] = i
	}
	return buf
}

func findIdx(what int, where []int) int {
	for i := 0; i < len(where); i++ {
		if where[i] == what {
			return i
		}
	}
	return -1
}

// Straight solution
func solvePart1Array(input int) int {
	buf := spinArray(input, 2018)
	return buf[(findIdx(2017, buf)+1)%len(buf)]
}

// Calcular just the initial position for each element and return that
// Fast because we don't adjust the elements when a new one is inserted
func calcInitialPositions(input int, iter int) []int {
	buf := make([]int, iter)
	finalPos := 0
	for i := 1; i < iter; i++ {
		finalPos = ((input%i)+finalPos)%i + 1
		buf[i] = finalPos
	}
	return buf
}

// Alternative solution, kind of complicated
func solvePart1ArrayAlt(input int) int {
	iter := 2018
	buf := calcInitialPositions(input, iter)

	// Find previous occupant of position 2018
	whichPos := buf[iter-1]
	insertedBefore := 0
	for i := iter - 2; i >= 0; i-- {
		if buf[i]+insertedBefore == whichPos {
			return i
		}
		if buf[i] < whichPos-insertedBefore {
			insertedBefore++
		}
	}
	return -1
}

// Fast solution for part 2
func solvePart2Array(input int) int {
	iter := 50000000
	buf := calcInitialPositions(input, iter)

	// Find the last that occupied position 1
	for i := iter - 1; i > 0; i-- {
		if buf[i] == 1 {
			return i
		}
	}
	return -1
}

// Solution with lists

// Next element in a circular list
func circularNext(l *list.List, e *list.Element) *list.Element {
	if e.Next() == nil {
		e = l.Front()
	} else {
		e = e.Next()
	}
	return e
}

func spinList(input int, iter int) *list.List {
	buf := list.New()
	e := buf.PushFront(0)
	for i := 1; i < iter; i++ {
		for j := 0; j < input; j++ {
			e = circularNext(buf, e)
		}
		e = buf.InsertAfter(i, e)
	}
	return buf
}

// Straightforward solution
func solvePart1List(input int) int {
	buf := spinList(input, 2018)

	for e := buf.Front(); e != nil; e = e.Next() {
		if e.Value == 2017 {
			e = circularNext(buf, e)
			return e.Value.(int)
		}
	}
	return 0
}

// Straightforward yet slow. Not sure if this is actually correct, as i never waited for it to finish.
// Should take a couple of hundred seconds...
func solvePart2List(input int) int {
	iter := 50000000
	buf := spinList(input, iter)

	return buf.Front().Next().Value.(int)
}

func solvePart1(input int) int {
	return solvePart1Array(input)
}

func solvePart2(input int) int {
	return solvePart2Array(input)
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	convertedInput := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(convertedInput))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(convertedInput))

}
