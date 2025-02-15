package common

import "fmt"

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func graphDfs(node *GraphNode, visited map[*GraphNode]bool) {
	if node == nil || visited[node] {
		return
	}
	fmt.Println(node.Val)
	visited[node] = true
	for _, neibhour := range node.Neighbors {
		graphDfs(neibhour, visited)
	}

}

func graphBfs(start *GraphNode) {
	if start == nil {
		return
	}

	visited := map[*GraphNode]bool{start: true}
	queue := []*GraphNode{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Val)
		for _, neighbor := range node.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
