package main

import "fmt"

func postOrder(root *Node) {
	var lastVisited *Node
	stack := &Stack{}
	node := root
	for !(stack.IsEmpty() && node == nil) {
		if node != nil {
			stack.Push(node)
			node = node.L
		} else {
			node = stack.Peek().(*Node)
			if node.R != nil && lastVisited != node.R {
				node = node.R
			} else {
				fmt.Printf("%s ", node)
				lastVisited = node
				stack.Pop()
				node = nil
			}
		}
	}
}
