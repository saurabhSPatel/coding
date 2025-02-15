package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	l, r := dummy, head
	for n > 0 {
		r = r.Next
		n--
	}
	fmt.Println(r.Val, l)
	for r != nil {
		l = l.Next
		r = r.Next
	}
	fmt.Println(l.Val, l)
	l.Next = l.Next.Next

	return dummy.Next
}
