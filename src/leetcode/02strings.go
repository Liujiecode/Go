package leetcode

import (
	"fmt"
)

/*单指针*/
func SortColors(nums []int) []int {
	n := len(nums)
	p0 := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0 += 1
		}
	}
	for j := p0; j < n; j++ {
		if nums[j] == 1 {
			nums[j], nums[p0] = nums[p0], nums[j]
			p0 += 1
		}
	}
	return nums
}

/*双指针*/
func SortColors_B(nums []int) []int {
	// n := len(nums)
	p0, p1 := 0, 0
	for index, v := range nums {
		if v == 1 {
			nums[index], nums[p1] = nums[p1], nums[index]
			p1 += 1
		} else if v == 0 {
			nums[index], nums[p0] = nums[p0], nums[index]
			if p0 < p1 {
				nums[p1], nums[index] = nums[index], nums[p1]
			}
			p0 += 1
			p1 += 1
		}

	}
	return nums
}

func S_V(v int) string {
	var s_num, double_num, only_num int
	var res []int
	if v == 1 {
		return "a"
	} else if v == 2 {
		return "ba"
	} else if v == 3 {
		return "baa"
	} else {
		for i := 4; i < 100; i++ {
			if i*(i-1)/2 > v {
				s_num = i
				break
			}
		}
		only_num = s_num*(s_num-1)/2 - v
		for i := 0; i < only_num; i++ {
			c := 97 + i + s_num - only_num
			res = append(res, c)
		}
		for i := 0; i < double_num; i++ {
			c := 97 + i
			res = append(res, c, c)
		}
		fmt.Print(res)
		return "ffff"
		// return res
	}
}

/*39.组合总合*/
func CombinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start int, temp []int, sum int)

	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTmp := make([]int, len(temp))
				copy(newTmp, temp)
				res = append(res, newTmp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
