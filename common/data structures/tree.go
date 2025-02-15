package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsTraversal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfsTraversal(root.Left),
		dfsTraversal(root.Right)) + 1
}

func bfsTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.Val)
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

}
