package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Tree_Solve() {
	list1 := []int{1, 2, 3, 4, 5, 6}
	stack := make([]TreeNode, 6)
	for i := 0; i < len(list1); i++ {
		stack[i].Val = list1[i]
	}
	for i := 0; i < len(list1)/2; i++ {
		if i*2+1 < len(list1) && i*2+2 < len(list1) {
			stack[i].Left = &stack[i*2+1]
			stack[i].Right = &stack[i*2+2]
		} else if i*2+1 < len(list1) && i*2+2 >= len(list1) {
			stack[i].Left = &stack[i*2+1]

		}

	}
	fmt.Println(stack)
}
