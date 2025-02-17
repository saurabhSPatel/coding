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

func dfsTraversalItr(root *TreeNode) int {

	st := []*TreeNode{root}

	for len(st) > 0 {

	}
	return 0
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		node1 := queue1[0]
		node2 := queue2[0]
		if node1.Val != node2.Val {
			return false
		}
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if node1.Left != nil {
			queue1 = append(queue1, node1.Left)
		}
		if node2.Left != nil {
			queue2 = append(queue2, node2.Left)
		}
		if node1.Right != nil {
			queue1 = append(queue1, node1.Right)
		}

		if node2.Right != nil {
			queue2 = append(queue2, node2.Right)
		}

	}
	return true
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

func diameterOfBinaryTree(root *TreeNode) int {

	res := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh, rh := dfs(root.Left), dfs(root.Right)
		res = max(res, lh+rh)
		return 1 + max(lh, rh)
	}

	dfs(root)
	return res
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxHeight(root.Left), maxHeight(root.Right))

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {

	res := true
	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := dfs(root.Left)
		rh := dfs(root.Right)
		if abs(rh-lh) > 1 {
			res = false
		}
		return 1 + max(lh, rh)

	}
	dfs(root)
	return res

}
