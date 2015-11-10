package main

import (
	"fmt"
	"strings"
)

type Path struct {
	nodes []*Node
}

func (p *Path) Add(n *Node) {
	p.nodes = append(p.nodes, n)
}

func (p *Path) String() string {
	var nodeStr []string
	for i := len(p.nodes) - 1; i >= 0; i-- {
		nodeStr = append(nodeStr, p.nodes[i].String())
	}
	return strings.Join(nodeStr, "->")
}

// 4.9
func FindPath(n *Node, sum int) []*Path {
	var paths []*Path

	fmt.Printf("starting in FindPath: %s, %d\n", n, sum)

	diff := sum - n.Val
	if diff == 0 {
		path := &Path{
			nodes: []*Node{n},
		}
		return []*Path{path}
	}

	if diff < 0 {
		return nil
	}

	var LPaths, RPaths []*Path

	if n.L != nil {
		LPaths = FindPath(n.L, diff)
		for _, p := range LPaths {
			p.Add(n)
		}
		if LPaths != nil && len(LPaths) > 0 {
			paths = append(paths, LPaths...)
		}
	}

	if n.R != nil {
		RPaths = FindPath(n.R, diff)
		for _, p := range RPaths {
			p.Add(n)
		}
		if RPaths != nil && len(RPaths) > 0 {
			paths = append(paths, RPaths...)
		}
	}

	fmt.Printf("in FindPath: %s, %d, %s\n", n, sum, paths)
	return paths
}

func FindPathRoot(root *Node, sum int) []*Path {
	var paths []*Path
	if root == nil {
		return paths
	}
	paths = FindPath(root, sum)
	fmt.Printf("in FindPathRoot: %s, %d, %s\n", root, sum, paths)

	LPaths := FindPathRoot(root.L, sum)
	fmt.Printf("in FindPathRoot: LPaths %s\n", LPaths)

	paths = append(paths, LPaths...)

	RPaths := FindPathRoot(root.R, sum)
	fmt.Printf("in FindPathRoot: RPaths %s\n", RPaths)

	paths = append(paths, RPaths...)

	return paths
}
