package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Rule in the Turing machine
type Rule struct {
	id    string
	value []byte
	move  []int
	next  []string
}

// Validates the input
func validate(input []string) error {
	// Accept
	return nil
}

func newRule(id string) Rule {
	rule := Rule{id: id}
	rule.value = make([]byte, 2)
	rule.move = make([]int, 2)
	rule.next = make([]string, 2)
	return rule
}

// Converts to an appropriate format
func convert(input []string) (start string, rules map[string]Rule, steps int) {
	var currentRule Rule
	var value, currentIdx int

	dirMap := map[string]int{"right": 1, "left": -1}

	rules = make(map[string]Rule)
	for i := 0; i < len(input); i++ {
		line := strings.Trim(input[i], ".:")

		if strings.Contains(line, "Begin in state") {
			start = line[strings.LastIndex(line, " ")+1:]
		}

		if strings.Contains(line, "Perform a diagnostic") {
			aux := strings.Split(line, " ")
			steps = helpers.Atoi(aux[len(aux)-2])
		}

		if strings.Contains(line, "In state") {
			id := line[strings.LastIndex(line, " ")+1:]
			currentRule = newRule(id)
			rules[id] = currentRule
		}

		if strings.Contains(line, "If the current value is") {
			currentIdx = helpers.Atoi(line[strings.LastIndex(line, " ")+1:])
		}

		if strings.Contains(line, "Write the value") {
			value = helpers.Atoi(line[strings.LastIndex(line, " ")+1:])
			currentRule.value[currentIdx] = byte(value)
		}

		if strings.Contains(line, "Move one slot") {
			dir := line[strings.LastIndex(line, " ")+1:]
			currentRule.move[currentIdx] = dirMap[dir]
		}

		if strings.Contains(line, "Continue with state") {
			currentRule.next[currentIdx] = line[strings.LastIndex(line, " ")+1:]
		}
	}

	return start, rules, steps
}

func printRules(rules map[string]Rule) {
	for k, v := range rules {
		fmt.Println("State ", k)
		for i := 0; i <= 1; i++ {
			fmt.Println("\t", "If the current value is:", i)
			fmt.Println("\t\t", "Write the value", v.value[i])
			fmt.Println("\t\t", "Move one slot to", v.move[i])
			fmt.Println("\t\t", "Continue with state", v.next[i])
		}
	}
}

func turing(start string, rules map[string]Rule, steps int) (tape []byte) {
	tape = make([]byte, steps)

	if len(rules) == 0 {
		return tape
	}

	slot := 0
	currValue := tape[slot]
	rule := rules[start]

	for step := 0; step < steps; step++ {
		currValue = tape[slot]
		tape[slot] = rule.value[currValue]
		slot = (slot + rule.move[currValue]) % steps
		if slot < 0 {
			slot = steps - 1
		}
		rule = rules[rule.next[currValue]]
	}

	return tape
}

func solvePart1(start string, rules map[string]Rule, steps int) int {
	tape := turing(start, rules, steps)

	checksum := 0
	for _, b := range tape {
		checksum += int(b)
	}
	return checksum
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	start, rules, steps := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(start, rules, steps))
}
