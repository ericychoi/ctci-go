package main

import "fmt"

func main() {
	root := buildTree()
	lists := createDLL(root)
	for k, v := range lists {
		fmt.Printf("%d -> %s\n", k, v)
	}
}

func buildTree() *Node {
	a := &Node{value: "A"}
	b := &Node{value: "B"}
	c := &Node{value: "C"}
	d := &Node{value: "D"}
	e := &Node{value: "E"}
	f := &Node{value: "F"}

	a.L = b
	a.R = c
	c.L = d
	c.R = f
	d.R = e

	return a
}

type LNode struct {
	Value *Node
	Next  *LNode
}

func (l *LNode) String() string {
	var str string
	var i *LNode
	for i = l; i.Next != nil; i = i.Next {
		str += i.Value.String()
	}
	str += i.Value.String()
	return str
}

func createDLL(n *Node) map[int]*LNode {
	lists := make(map[int]*LNode)
	DFSD(n, 0, lists)
	return lists
}

func DFSD(n *Node, d int, l map[int]*LNode) {
	visit(n, d, l)
	if n.L != nil {
		DFSD(n.L, d+1, l)
	}
	if n.R != nil {
		DFSD(n.R, d+1, l)
	}
}

func visit(n *Node, d int, l map[int]*LNode) {
	ln := &LNode{Value: n}
	if l[d] == nil {
		l[d] = ln
	} else {
		var ptr *LNode
		for ptr = l[d]; ptr.Next != nil; ptr = ptr.Next {
		}
		ptr.Next = ln
	}
}
