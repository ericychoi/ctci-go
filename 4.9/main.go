package main

import "fmt"

func main() {
	root := buildTree()
	paths := FindPathRoot(root, 5)
	for _, p := range paths {
		fmt.Printf("%s\n", p)
	}
}

func buildTree() *Node {
	a := &Node{value: "A", Val: 1}
	b := &Node{value: "B", Val: 2}
	c := &Node{value: "C", Val: 1}
	d := &Node{value: "D", Val: 0}
	e := &Node{value: "E", Val: 3}
	f := &Node{value: "F", Val: 2}
	g := &Node{value: "G", Val: 3}
	h := &Node{value: "H", Val: 1}
	i := &Node{value: "I", Val: 2}
	j := &Node{value: "J", Val: 4}

	a.L = b
	a.R = c
	b.L = d
	c.L = e
	c.R = f
	d.R = g
	g.L = h
	g.R = i
	i.R = j

	return a
}

type Node struct {
	L       *Node
	R       *Node
	value   string
	Val     int
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
