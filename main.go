package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1] = []int{res[len(res)-1][0], max(intervals[i][1], res[len(res)-1][1])}
		} else {
			res = append(res, intervals[i])
		}
	}

	// Replace this placeholder return statement with your code
	return res
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minAbsD := math.MaxInt
	var res [][]int
	for i := 0; i < len(arr)-1; i++ {
		minAbsD = min(minAbsD, arr[i+1]-arr[i])
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minAbsD {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res

}

func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}

func main() {
	// node1 := &TreeNode{Val: 1}
	// node2 := &TreeNode{Val: 2}
	// node3 := &TreeNode{Val: 3}
	// node4 := &TreeNode{Val: 4}
	// node5 := &TreeNode{Val: 5}
	// node6 := &TreeNode{Val: 6}
	// node7 := &TreeNode{Val: 7}

	// // Manually linking nodes to create the tree
	// node1.Left = node2
	// node1.Right = node3

	// node2.Left = node4
	// node2.Right = node5

	// node3.Left = node6
	// node3.Right = node7

	// Printing tree structure for verification
	// fmt.Println(dfsTraversal(node1))

	node1 := &GraphNode{Val: 1}
	node2 := &GraphNode{Val: 2}
	node3 := &GraphNode{Val: 3}
	node4 := &GraphNode{Val: 4}
	node5 := &GraphNode{Val: 5}

	// Establish the neighbors (edges) between nodes
	node1.Neighbors = []*GraphNode{node2, node3}
	node2.Neighbors = []*GraphNode{node1, node4}
	node3.Neighbors = []*GraphNode{node1, node5}
	node4.Neighbors = []*GraphNode{node2}
	node5.Neighbors = []*GraphNode{node3}
	// visited := make(map[*GraphNode]bool)
	// graphDfs(node1, visited)
	// graphBfs(node1)

	// fmt.Println(spiralOrder([][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// },
	// ))

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val:  2,
	// 		Next: nil,
	// 	},
	// }
	head = removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
