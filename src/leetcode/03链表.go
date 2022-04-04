package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNode_Solve() {
	head := &ListNode{Val: 1, Next: nil}
	newhead := &ListNode{Next: head}
	fmt.Println(newhead.Next.Val)
}
