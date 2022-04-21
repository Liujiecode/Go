package main

import (
	"fmt"
	"sort"
)

func main() {
	// go_learn.Str_Test()
	// leetcode.Tree_Solve()
	// go_learn.Select_solve()
	num := []int{0, 1}
	//无限循环
	// for i := 0; i < len(num); i++ {
	// 	num = append(num, i+1)
	// }
	//for range是先拷贝再操作！
	for i := range num {
		num = append(num, i+1)
	}
	fmt.Println(num)

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
