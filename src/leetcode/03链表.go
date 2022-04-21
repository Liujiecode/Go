package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func POP(head *ListNode, index int) *ListNode {
	newhead := head
	for i := 0; i < index; i++ {
		newhead = newhead.Next
	}
	newhead.Next = newhead.Next.Next
	return head
}
func Push(head *ListNode, target int) *ListNode {
	newhead := head
	for newhead != nil {
		newhead = newhead.Next
	}
	newhead.Next.Val = target
	return head
}

func ListNode_Solve() {
	head := ListNode{Val: 1}
	Push(&head, 2)
	Push(&head, 3)
	fmt.Println(head.Next)
	// list := []int{1, 2, 3, 4, 5}
	// head := ListNode{Val: list[0]}
	// for i := 1; i < len(list); i++ {
	// 	head.Next.Val = list[i]
	// }
	// fmt.Println(head.Next)

}
