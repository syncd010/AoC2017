package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Lazy, just trust the input...
	return nil
}

// Converts the input to an appropriate structure
func convert(input []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range input {
		elems := strings.Fields(line)
		if len(elems) < 2 {
			continue
		}
		for _, e := range elems[2:] {
			graph[elems[0]] = append(graph[elems[0]], strings.Split(strings.Trim(e, " ,"), ",")...)
		}
	}

	return graph
}

func getReachable(graph map[string][]string, from string) map[string]bool {
	visited := make(map[string]bool)

	reachable := graph[from]
	// Do a Breadth First search, keeping tab of the visited nodes
	for len(reachable) > 0 {
		var nextReachable []string
		for _, node := range reachable {
			if visited[node] {
				continue
			}
			visited[node] = true
			nextReachable = append(nextReachable, graph[node]...)
		}
		reachable = nextReachable
	}

	return visited
}

func solvePart1(graph map[string][]string) int {
	return len(getReachable(graph, "0"))
}

func solvePart2(graph map[string][]string) int {
	// We assume that if `b in graph[a]` then `a in graph[b]`
	connectedGraphs := 0

	visited := make(map[string]bool)
	for k := range graph {
		if visited[k] {
			continue
		}
		for v := range getReachable(graph, k) {
			visited[v] = true
		}
		connectedGraphs++
	}
	return connectedGraphs
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	graph := convert(input)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(graph))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(graph))
}
