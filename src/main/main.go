package main

import (
	"go_learn"
	"leetcode"
)

func main() {
	// list := []int{3, 2, 1, 0, 4}
	// str := [][3]byte{{'a', 'b', 'c'}, {'1', '0', '2'}}
	// str2 := [][]int{}
	// fmt.Println(list, str, str2)
	// var list2 []int
	// list2 = append(list2, 1)
	// fmt.Println(list2)
	// var c Hello
	// fmt.Println(unsafe.Sizeof(c))
	// fmt.Println(unsafe.Sizeof(struct{}{}))
	// fmt.Println(&str[0][0], &str[0][1], &str[0][2], &str[1][0], &str[1][1], &str[1][2])
	go_learn.Str_learn()
	leetcode.ListNode_Solve()

}

type Hello struct {
	a int32
	b int32
}
