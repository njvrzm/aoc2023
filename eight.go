package aoc2023

import (
	"bufio"
	"log"
	"regexp"
)

type Network struct {
	nodes        map[string]Node
	instructions string
	location     Node
}

func (n *Network) PartOne() Result {
	return NumberResult{n.CountSteps(n.nodes["AAA"], func(node Node) bool { return node.label == "ZZZ" })}
}

func (n *Network) PartTwo() Result {
	return NumberResult{n.FindAll()}
}

func (n *Network) FindAll() int {
	endWithA := make([]Node, 0)
	for _, node := range n.nodes {
		if node.label[2] == 'A' {
			endWithA = append(endWithA, node)
		}
	}
	times := make([]int, len(endWithA))
	for i, node := range endWithA {
		times[i] = n.CountSteps(node, func(node Node) bool { return node.label[2] == 'Z' })
	}
	return LCM(times...)
}
func (n *Network) CountSteps(location Node, done func(Node) bool) int {
	il := len(n.instructions)

	for i := 0; ; i++ {
		if done(location) {
			return i
		}
		switch n.instructions[i%il] {
		case 'R':
			location = n.nodes[location.right]
		case 'L':
			location = n.nodes[location.left]
		default:
			log.Fatalf("What do I do with %c", n.instructions[i%il])
		}
	}
}
func (n *Network) Load(scanner *bufio.Scanner) {
	scanner.Scan()
	n.nodes = make(map[string]Node)
	n.instructions = scanner.Text()
	ReadBlank(scanner)
	for scanner.Scan() {
		node := ReadNode(scanner)
		n.nodes[node.label] = node
	}
	lc := make(map[string]int)
	rc := make(map[string]int)
	for _, node := range n.nodes {
		lc[node.left] += 1
		rc[node.right] += 1
	}
}
func (n *Network) GoTo(node Node) {
	n.location = node
}

type Node struct {
	label string
	left  string
	right string
}

var reNode = regexp.MustCompile(`[A-Z0-9]{3}`)

func ReadNode(scanner *bufio.Scanner) Node {
	parts := reNode.FindAllString(scanner.Text(), 3)
	return Node{label: parts[0], left: parts[1], right: parts[2]}
}
