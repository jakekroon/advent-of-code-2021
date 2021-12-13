package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Graph struct {
	nodes map[string]*Node
}

type Node struct {
	label   string
	small   bool
	visited int
	next    []*Node
}

func (g *Graph) addNode(label string) *Node {
	if _, ok := g.nodes[label]; !ok {
		g.nodes[label] = &Node{label: label, small: unicode.IsLower(rune(label[0]))}
	}
	return g.nodes[label]
}

func (g *Graph) addEdge(src, dest string) {
	srcNode := g.addNode(src)
	destNode := g.addNode(dest)
	srcNode.next = append(srcNode.next, destNode)
	destNode.next = append(destNode.next, srcNode)
}

func (caves *Graph) countPaths(src *Node, dest *Node, count int, twice bool) int {
	if src == dest {
		return count + 1
	}

	for _, nextNode := range src.next {
		visitable := !nextNode.small || nextNode.visited == 0 || (nextNode.label != "start" && !twice)
		if visitable {
			nextNode.visited += 1
			count = caves.countPaths(nextNode, dest, count, twice ||
				(nextNode.small && nextNode.visited == 2))
			nextNode.visited -= 1
		}
	}
	return count
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	caveSystem := Graph{nodes: map[string]*Node{}}

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "-")
		caveSystem.addEdge(points[0], points[1])
	}

	caveSystem.nodes["start"].visited = 1
	totalPaths := caveSystem.countPaths(caveSystem.nodes["start"], caveSystem.nodes["end"], 0, false)

	fmt.Printf("Total paths through caves %v\n", totalPaths)
}
