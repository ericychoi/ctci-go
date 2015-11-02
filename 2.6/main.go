package main

import (
	"fmt"
)

//TODO: compare the answer
type Node struct {
	Next    *Node
	Content string
}

func main() {
	a := Node{Content: "A"}
	b := Node{Content: "B"}
	a.Next = &b
	b.Next = &a
	fmt.Printf("%+v %+v\n\n", a, b)

	fmt.Printf("%+v\n", findCircular(&a))
}

func findCircular(head *Node) *Node {
	var seen map[*Node]bool
	seen = make(map[*Node]bool)
	seen[head] = true
	for n := head; n.Next != nil; n = n.Next {
		next := n.Next
		fmt.Printf("seen: %+v\nn: %+v\nnext: %+v\n\n", seen, n, next)
		if seen[next] {
			return n.Next
		}
		seen[next] = true
	}
	return nil
}
