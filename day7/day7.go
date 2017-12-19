package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/syncd010/AoC2017/helpers"
)

// Validates the input
func validate(input []string) error {
	// Lazy, just trust the input...
	return nil
}

// Tree structure
type Tree struct {
	name        string
	weight      int
	totalWeigth int
	children    []string
	parent      string
}

// Converts the input to an appropriate structure
func convert(input []string) (treeMap map[string]*Tree, root string) {
	treeMap = make(map[string]*Tree)

	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		node := newNode(line)
		treeMap[node.name] = &node
	}

	// Fill parents
	for parent, node := range treeMap {
		// This is for the next part
		root = parent
		for _, child := range node.children {
			treeMap[child].parent = parent
		}
	}

	// Starting from any node, find the root
	for treeMap[root].parent != "" {
		root = treeMap[root].parent
	}

	// Fill totalWeigths
	fillTotalWeight(root, treeMap)

	return treeMap, root
}

func fillTotalWeight(nodeName string, tree map[string]*Tree) {
	node := tree[nodeName]
	weight := node.weight
	for _, child := range node.children {
		fillTotalWeight(child, tree)
		weight += tree[child].totalWeigth
	}
	node.totalWeigth = weight
}

func newNode(input string) Tree {
	re := regexp.MustCompile("[A-Za-z0-9]+")
	stripped := re.FindAllString(input, -1)

	weight := 0
	if len(stripped) > 1 {
		var err error
		weight, err = strconv.Atoi(stripped[1])
		helpers.Check(err, "Invalid number format")
	}

	var children []string
	if len(stripped) > 2 {
		children = stripped[2:]
	}
	t := Tree{name: stripped[0], weight: weight, children: children}

	return t
}

func solvePart1(tree map[string]*Tree, root string) string {
	return root
}

func getUnbalancedNode(tree map[string]*Tree, nodeName string) (unbalancedNode string, balancedWeight int) {
	unbalancedChildren := getUnbalancedChildren(tree, nodeName)
	if len(unbalancedChildren) == 0 {
		// Unbalance not on our children
		return "", 0
	}

	for _, child := range unbalancedChildren {
		unbalancedNode, balancedWeight = getUnbalancedNode(tree, child)
		if unbalancedNode != "" {
			return unbalancedNode, balancedWeight
		}
	}
	// We have unbalanced children but they all have balanced children, so one of them is the unbalance
	// Just get the first one
	unbalancedNode = unbalancedChildren[0]
	balancedWeight = tree[unbalancedNode].weight
	node := tree[nodeName]
	for _, child := range node.children {
		if child != unbalancedNode {
			balancedWeight -= (tree[unbalancedNode].totalWeigth - tree[child].totalWeigth)
			break
		}
	}
	return unbalancedNode, balancedWeight
}

func getUnbalancedChildren(tree map[string]*Tree, nodeName string) (unbalancedChildren []string) {
	node := tree[nodeName]
	// Auxiliary map from weigths to node names
	auxMap := make(map[int][]string)
	for _, child := range node.children {
		k := tree[child].totalWeigth
		auxMap[k] = append(auxMap[k], child)
	}

	for _, children := range auxMap {
		if len(children) == 1 {
			unbalancedChildren = append(unbalancedChildren, children[0])
		}
	}
	return unbalancedChildren
}

func solvePart2(tree map[string]*Tree, nodeName string) int {
	_, balancedWeight := getUnbalancedNode(tree, nodeName)
	return balancedWeight
}

func main() {
	input := helpers.ReadInput(os.Args[1:]...)
	helpers.Check(validate(input), "Please provide a valid input")

	tree, root := convert(input)
	//	printTree(tree, root, 0)

	fmt.Printf("Fist part of the quiz is: %v\n", solvePart1(tree, root))
	fmt.Printf("Second part of the quiz is: %v\n", solvePart2(tree, root))

}

func printTree(tree map[string]*Tree, nodeName string, ident int) {
	node := tree[nodeName]
	for i := 0; i < ident; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("%v \t w:%v \t tw:%v\n", node.name, node.weight, node.totalWeigth)
	for _, child := range node.children {
		printTree(tree, child, ident+1)
	}
}
