package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("====主函数入口====")
	// array := []int{2, 3, 1, 4, 0, 6, 7, 5}
	/*数组取地址
	sliceasdas := []int{1, 2, 3}
	fmt.Printf("%p  %p  %p\n", sliceasdas, &sliceasdas[0], &sliceasdas[1])
	fmt.Printf("%p  %p  %p\n", &sliceasdas, &sliceasdas[0], &sliceasdas[1])
	fmt.Printf("%p  %p  %p  %p\n", &sliceasdas, &sliceasdas[0], &sliceasdas[1], *&sliceasdas)
	*/
	// list1 := []int{1, 2, 3, 4, 5, 6}
	// rotate(list1, 3)
	// pushed := []int{1, 2, 3, 4, 5}
	// poped := []int{4, 5, 3, 2, 1}
	// validateStackSequences(pushed, poped)
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	//           [1, 1, 4,  2, 1,  1,  0, 0]

}
func dailyTemperatures(temperatures []int) {
	n := len(temperatures)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}
	fmt.Println(res)
}
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0
	length := len(pushed)
	for i := 0; i < length; i++ {
		stack = append(stack, pushed[i])
		for index < len(popped) && popped[index] == stack[len(stack)-1] && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	fmt.Println(len(stack) == 0)
	return len(stack) == 0
}
func rotate(nums []int, k int) {
	fmt.Printf("%#v---%T--%v--%v--%p\n", nums, nums, &nums, reflect.TypeOf(nums), nums)
	n := len(nums)
	left := nums[:n-k]
	right := nums[n-k:]
	fmt.Printf("%#v---%T--%v--%v--%p\n", left, left, &left, reflect.TypeOf(left), left)
	fmt.Printf("%#v---%T--%v--%v--%p\n", right, right, &right, reflect.TypeOf(right), right)
	nums = append(right, left...)
	// fmt.Println("right", right)
	// all2one[0] = 100000

	fmt.Printf("%#v---%T--%v--%v--%p", nums, nums, &nums, reflect.TypeOf(nums), nums)

}
func calculate(s string) int {
	n := len(s)
	ans, num, sign := 0, 0, 1
	st := []int{}
	for i := 0; i < n; i++ {
		switch s[i] {
		case '+':
			ans += sign * num
			num = 0
			sign = 1
		case '-':
			ans += sign * num
			num = 0
			sign = -1
		case '(':
			st = append(st, ans)
			st = append(st, sign) // 对符号的处理
			ans = 0
			sign = 1
		case ')':
			ans += sign * num
			num = 0
			ans *= st[len(st)-1]
			st = st[:len(st)-1]
			ans += st[len(st)-1]
			st = st[:len(st)-1]
		case ' ':
			break
		default: // 数字
			num = 10*num + int(s[i]-'0')
		}
	}
	ans += sign * num
	fmt.Printf("s:%s ans: %d", s, ans)
	return ans
}
