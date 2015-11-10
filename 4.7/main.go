package main

import "fmt"

func main() {
	root, d, g := buildTree()
	common := findCommon(root, d, g)
	fmt.Printf("found: %s\n", common)

	common = findCommon(root, g, root.L.R.R)
	fmt.Printf("found: %s\n", common)
}

func buildTree() (*Node, *Node, *Node) {
	a := &Node{value: "A"}
	b := &Node{value: "B"}
	c := &Node{value: "C"}
	d := &Node{value: "D"}
	e := &Node{value: "E"}
	g := &Node{value: "G"}
	h := &Node{value: "H"}
	i := &Node{value: "I"}

	a.L = b
	a.R = c
	b.L = d
	b.R = e
	e.L = g
	e.R = i
	c.L = h

	return a, d, g
}

func DFS(root, n *Node) bool {
	if root == nil {
		return false
	}
	if root == n {
		return true
	}

	if root.L != nil && DFS(root.L, n) {
		return true
	}

	if root.R != nil && DFS(root.R, n) {
		return true
	}

	return false
}

func findParent(root, n, parent *Node) *Node {
	if root == nil {
		return nil
	}
	if root == n {
		return parent
	}
	var p *Node

	if root.L != nil {
		p = findParent(root.L, n, root)
		if p != nil {
			return p
		}
	}

	if root.R != nil {
		p = findParent(root.R, n, root)
		if p != nil {
			return p
		}
	}

	return nil
}

func findCommon(root, i, j *Node) *Node {
	var parent *Node

	// check the starting condition first
	if DFS(i, j) {
		return i
	}

	n := i
	for root != parent {
		parent = findParent(root, n, nil)
		fmt.Printf("found parent %s root %s n %s\n", parent, root, n)

		if DFS(parent, j) {
			fmt.Printf("found parent in j %s\n", j)
			return parent
		}
		n = parent
	}

	return root
}
