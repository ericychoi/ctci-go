package main

import "fmt"

func preOrder(root *Node) {
	stack := &Stack{}
	node := root
	for !(stack.IsEmpty() && node == nil) {
		if node != nil {
			fmt.Printf("%s ", node)
			stack.Push(node)
			node = node.L
		} else {
			node = stack.Pop().(*Node)
			node = node.R
		}
	}
}
