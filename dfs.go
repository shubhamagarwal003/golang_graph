package graph

import (
	"fmt"
)

func dfs(node *Node, visit map[*Node]bool) {
	visit[node] = true
	fmt.Println(node.Value)
	for _, e := range node.Edges {
		if !visit[e.Dest] {
			dfs(e.Dest, visit)
		}
	}
}

func DFS(nodes []*Node) {
	visit := make(map[*Node]bool)
	for _, n := range nodes {
		visit[n] = false
	}
	for _, n := range nodes {
		if !visit[n] {
			dfs(n, visit)
		}
	}
}
