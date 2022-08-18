package main

import (
	"fmt"
	"go_learn"
	"sort"
)

func main() {
	go_learn.Str_Test()
	// // go_learn.Str_Test()
	// // leetcode.Tree_Solve()
	// leetcode.ListNode_Solve()
	// test := [][]int{{1, 2}, {1, 1}, {2, 4}, {1, 5}}
	// sort.Slice(test, func(a, b int) bool {
	// 	return test[a][0] < test[b][0]
	// })
	// fmt.Println(test)
	// // go_learn.Select_solve()

}
func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	hash := make(map[int]int)
	count := 0
	x_list := []int{}
	y_list := []int{}
	for _, v := range nums {
		if c, ok := hash[v]; c > 0 && ok {
			hash[v] += 1
		} else {
			hash[v] = 1
			count += 1
		}
	}
	for x, y := range hash {
		x_list = append(x_list, x)
		y_list = append(y_list, y)
	}
	for i := 0; i < count-k; i++ {
		for j := 0; j < count-1; j++ {
			if y_list[j] > y_list[j+1] {
				y_list[j], y_list[j+1] = y_list[j+1], y_list[j]
				x_list[j], x_list[j+1] = x_list[j+1], x_list[j]
			}
		}
	}
	fmt.Println(x_list[count-k:])
	return x_list[count-k:]

}
