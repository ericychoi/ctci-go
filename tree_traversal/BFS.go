package main

func BFS(root *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		node.Visit()
		if node.L != nil && !node.L.IsVisited() {
			queue = append(queue, node.L)
		}
		if node.R != nil && !node.R.IsVisited() {
			queue = append(queue, node.R)
		}
	}
}
