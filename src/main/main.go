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
	// leetcode.Insert_Sort([]int{6, 5, 4, 1, 3, 2})
	// arr := []int{6, 1, 4, 2, 3, 7}
	// fmt.Println(arr)
	// // leetcode.Quick_Sort2(arr, 0, 5)
	// // leetcode.Shell_Sort(arr)
	// // leetcode.QuickSort(arr)
	// // leetcode.CountSort(arr)
	// leetcode.SortArray(arr)
	// fmt.Println(arr)
	// str := "10"
	// str2 := 101
	// va, _ := strconv.ParseInt(str, 10, 16)
	// fmt.Println(va)
	// fmt.Printf("%T\n", va)
	// str3 := strconv.FormatInt(int64(str2), 2)
	// fmt.Printf("%v%T\n", str3, str3)
	// var a byte = '1'
	// b := 'A'
	// c := 'a'
	// d := int(a) - 10
	// fmt.Println(a, b, c)
	// fmt.Println(d)
	// s := fmt.Sprintf("%d", 10)
	// fmt.Println(s)
	// old := "heello"
	// new := strings.Replace(old, "e", "11e", -1)
	// new1 := strings.Replace(old, "e", "11e", -2)
	// new2 := strings.Replace(old, "e", "11e", 1)
	// fmt.Println(new, new1, new2)
	// // leetcode.CombinationSum([]int{1, 2, 3, 6, 7}, 8)
	// list1 := []int{}
	// list2 := make([]int, 5)
	// var list3 []int
	// fmt.Printf("%#v,%#v, %#v", list1, list2, list3)
	// fmt.Printf("%#v,%#v ,%#v ", len(list1), len(list2), len(list3))

	// a := reflect.DeepEqual([]int{1, 2}, []int{1, 2})
	// fmt.Println(a)

	// nums := [][]int{{1, 2}, {2, 3}, {2, 3}, {3, 2}, {3, 4}}
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if reflect.DeepEqual(nums[i], nums[j]) {
	// 			fmt.Println("true")

	// 		} else {
	// 			fmt.Println("false")
	// 		}

	// 	}
	// }
	// reversePairs([]int{2, 4, 3, 5, 1})
	// leetcode.MergeSort1([]int{7, 6, 5, 2, 3, 8, 4}, 0, 6)
	// merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	nums := []int{2, 0, 2, 1, 1, 0}
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p1++
			p0++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	/*copy+sort*/
	// nums1=append(nums1[:m],nums2...)
	// sort.Ints(nums1)
	/*双指针:从前往后*/
	// res:=make([]int,0,m+n)
	// res:=[]int{}
	// p,q:=0,0
	// for{
	//     if p==m{
	//         res=append(res,nums2[q:]...)
	//         break
	//     }
	//     if q==n{
	//         res=append(res,nums1[p:]...)
	//         break
	//     }
	//     if nums1[p]<=nums2[q]{
	//         res=append(res,nums1[p])
	//         p++
	//     }else {
	//         res=append(res,nums2[q])
	//         q++
	//     }
	// }
	// copy(nums1,res)
	/*双指针：从后往前*/
	l, r, res := m-1, n-1, m+n-1
	for {
		if l < 0 {
			nums1 = append(nums2[:r], nums1[res:]...)
			break
		}
		if r < 0 {
			nums1 = append(nums1[:l], nums1[res:]...)
			break
		}
		if nums1[l] >= nums1[r] {
			nums1[res] = nums1[l]
			res--
			l--
		} else {
			nums1[res] = nums2[r]
			res--
			r--
		}
	}

}

func reversePairs(nums []int) int {
	n := len(nums)
	temp := []int{}
	var res int
	for i := 0; i < n; i++ {
		num := 0
		if i > 0 && nums[i] >= nums[i-1] && len(temp) > 0 {
			temp = append(temp, temp[len(temp)-1])
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				num += 1
			}
		}
		if num == 0 {
			continue
		} else {
			temp = append(temp, num)
		}

	}
	for i := 0; i < len(temp); i++ {
		res += temp[i]
	}
	return res
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
