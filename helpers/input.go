package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Check if the error is not null and panic with the message if so
func Check(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

// Atoi is Atoi without returning errors. Any error causes a panic
func Atoi(str string) int {
	res, err := strconv.Atoi(str)
	Check(err, "Invalid number")
	return res
}

// AAtoi is Array Atoi
func AAtoi(str ...string) []int {
	res := make([]int, len(str))
	var err error
	for i, s := range str {
		res[i], err = strconv.Atoi(s)
		Check(err, "Invalid number")
	}
	return res
}

// ReadInput reads the file passed in the args or from the standard input
// and returns its content, with empty lines stripped and all lines trimmed
// If it can't read the input, it panics
func ReadInput(args ...string) []string {
	var scanner *bufio.Scanner
	if len(args) == 0 {
		fmt.Printf("Input:\n")
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(args[0])
		Check(err, "Couldn't open file: "+args[0])
		scanner = bufio.NewScanner(f)
	}

	var res []string
	for scanner.Scan() {
		// s := strings.Trim(scanner.Text(), " \n\t")
		s := scanner.Text()
		if len(s) > 0 {
			res = append(res, s)
		}
	}
	Check(scanner.Err(), "Error reading input")
	return res
}
