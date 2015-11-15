package main

import "fmt"

type Node struct {
	L       *Node
	R       *Node
	value   string
	visited bool
}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.value)
}

func (n *Node) IsVisited() bool {
	return n.visited
}

func (n *Node) Visit() {
	fmt.Printf("%s ", n)
	n.visited = true
}
