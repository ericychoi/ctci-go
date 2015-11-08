package main

import "fmt"

func main() {
	root := buildTree()
	inOrder(root)
}

func buildTree() *Node {
	a := &Node{value: "A"}
	b := &Node{value: "B"}
	c := &Node{value: "C"}
	d := &Node{value: "D"}
	e := &Node{value: "E"}
	f := &Node{value: "F"}
	g := &Node{value: "G"}
	h := &Node{value: "H"}
	i := &Node{value: "I"}

	f.L = b
	b.L = a
	b.R = d
	d.L = c
	d.R = e
	f.R = g
	g.R = i
	i.L = h

	return f
}

func inOrder(n *Node) {
	var queue []*Node
	c := n
	for len(queue) != 0 || c != nil {
		if c != nil {
			queue = append([]*Node{c}, queue...)
			c = c.L
		} else {
			c = queue[0]
			queue = queue[1:]
			fmt.Printf("%s ", c)
			c = c.R
		}
	}
	fmt.Println("")
}

type Node struct {
	L     *Node
	R     *Node
	value string
}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.value)
}