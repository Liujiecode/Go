package leetcode

import (
	"fmt"
)

/*1.冒泡排序
a.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
b.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
c.针对所有的元素重复以上的步骤，除了最后一个。
d.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。*/
func Bubble_Sort(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ { //n个数，排序n-1趟就可以
		for j := 0; j < n-i-1; j++ { //比较j与j+1，所以j+1<n，i趟后，可以j+1<n-i
			if array[j] > array[j+1] { //执行一趟，最大的数到最后
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("冒泡排序：", array)
	return array
}

//优化
func Bubble_Sort2(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ { //-i是因为一趟排好一个数，不需要再次遍历到n-1
			if array[j] >= array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true //某一趟，所有数都没变化，直接退出
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("冒泡排序优化：", array)
	return array
}

/*2.快速排序
a.从数列中挑出一个元素，称为 “基准”（pivot）;
b.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
c.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；*/
func Quick_Sort(array []int, l, r int) {
	if l >= r {
		fmt.Println("result:", array[l])
		return
	}
	i := parttion(array, l, r)
	Quick_Sort(array, 1, i-1)
	Quick_Sort(array, i+1, r)
}
func parttion(array []int, l, r int) int { //array := []int{2, 3, 1, 4, 0, 6, 7}
	i, j := l, r //array := []int{2, 3, 5,1, 4, 0, 6, 7}
	for i < j {
		for i < j && array[j] >= array[l] { //从右向左与2比较，>2，比较前一个数
			j -= 1
		}
		for i < j && array[i] <= array[l] { //从左向右与2比较，<2，比较后一个数
			i += 1
		}
		array[i], array[j] = array[j], array[i] //交换[1],[4]
		// fmt.Println("for内部i,j:", i, j)
	}
	array[l], array[i] = array[i], array[l] //交换[0]与[2]
	// fmt.Println("for外部l,i:", l, i)
	return i
}

//优化
func Quick_Sort2(array []int, l, r int) {
	for l < r {
		fmt.Println("快速排序优化：", array)
		i := parttion(array, l, r)
		if i-1 < r-1 {
			Quick_Sort2(array, l, i-1)
			l = i + 1
		} else {
			Quick_Sort2(array, i+1, r)
			r = i - 1
		}

	}
}

/*go写法，递归*/
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	base := data[0] //第一个元素当基准
	l, r := 0, len(data)-1
	for i := 1; i <= r; {
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
	}
	QuickSort(data[:l]) //data[l]就是基准的位置
	QuickSort(data[l+1:])
}

/*3.插入排序（插入排序在元素个数较少时和记录基本有序时效率是最高的）
a.插入排序第一个位置固定，故从第二个位置开始遍历
b.记录当前位置为Current,,从大到小遍历当前位置到0的值，倘若，前一个数大于后一个数，将前一个数赋值给后一个（因为最前面的数赋值给current）
因为从大到小遍历，j需要先运算再自减。遍历结束，循环退出。j会比0还小，为-1（当i的值最小使），将current的值给j，前面就排序完毕。*/
func Insert_Sort(array []int) []int { //array := []int{2, 3, 1, 4, 0, 6, 7}
	// 将数组的第一个元素当作已经排好序的，从第二个元素，即 i 从 1 开始遍历数组
	for i := 1; i < len(array); i++ {
		// 外围循环开始，把当前 i 指向的值用 current 保存
		current := array[i]
		// 指针 j 内循环，和 current 值比较，若 j 所指向的值比 current 值大，则该数右移一位
		j := i - 1
		for ; j >= 0 && array[j] > current; j-- {
			array[j+1] = array[j]
			// fmt.Printf("j+1:%d  j:%d\n", j+1, j)
		}
		// 内循环结束，j+1 所指向的位置就是 current 值插入的位置
		array[j+1] = current
		// fmt.Println("j+1", j+1)
	}
	fmt.Println("插入排序：", array)
	return array
}

/*4.插入排序变种——希尔排序
a.选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
b.按增量序列个数 k，对序列进行 k 趟排序；
c.每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。*/
func Shell_Sort(array []int) []int { //array := []int{2, 3, 1, 4, 0, 6, 7, 5}
	increment := len(array)
	for increment > 1 {
		increment = increment / 2
		for i := 0; i < increment; i++ {
			for j := i + increment; j < len(array); j += increment {
				temp := array[j]
				k := j - increment
				// for ; k >= 0; k -= increment {
				// 	if array[k] > temp {
				// 		array[k+increment] = array[k]
				// 		continue
				// 	}
				// 	break
				// }
				for ; k >= 0 && array[k] > temp; k -= increment {
					array[k+increment] = array[k]
				}
				array[k+increment] = temp
			}
		}
	}
	fmt.Println("希尔排序：", array)
	return array
}

/*5.简单选择排序(交换移动数据次数相当少，这样也就节省了排序时间,但不稳定)
a.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
b.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
c.重复第二步，直到所有元素均排序完毕。*/
/*主要思路就是我们每一趟在 n-i+1 个记录中选取关键字最小的记录作为有序序列的第 i 个记录。*/
func EasyChose_Sort(array []int) []int {
	// min := 0
	n := len(array)
	for i := 0; i < n-1; i++ { //需经过n-1次比较
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if min != i {
			array[min], array[i] = array[i], array[min]
		}
	}
	fmt.Println("简单选择排序:", array)
	return array
}

/*6.归并排序
a.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
b.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
c.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
d.重复步骤 3 直到某一指针达到序列尾；
e.将另一序列剩下的所有元素直接复制到合并序列尾。*/
// func reversePairs(nums []int) int {
//     n:=len(nums)
//     temp:=[]int{}
//     var res int
//     for i:=0;i<n;i++{
//         num:=0
//         if i>0 && nums[i]>=nums[i-1] && len(temp)>0{
//             temp=append(temp,temp[len(temp)-1])
//             continue
//         }
//         for j:=i+1;j<n;j++{
//             if nums[i]>nums[j]{
//                 num+=1
//             }
//         }
//         if num>0 && len(temp)==0{
//         temp=append(temp,num)
//         }else if len(temp)>0{
//         temp=append(temp,num)
//         }

//     }
//     for i:=0;i<len(temp);i++{
//         res+=temp[i]
//     }
//     return res
// }

func MergeSort1(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := MergeSort1(nums, start, mid) + MergeSort1(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}
func SortArray(nums []int) []int {
	mergeSort2(nums, 0, len(nums)-1)
	return nums
}
func mergeSort2(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)>>1
		mergeSort2(arr, left, mid)
		mergeSort2(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

//归并
func merge(arr []int, left, mid, right int) {
	//第一步，定义一个新的临时数组
	temparr := make([]int, right-left+1)
	// temparr:=[]int{}
	temp1, temp2 := left, mid+1
	index := 0
	//对应第二步，比较每个指针指向的值，小的存入大集合
	for temp1 <= mid && temp2 <= right {
		if arr[temp1] <= arr[temp2] {
			temparr[index] = arr[temp1]
			index += 1
			temp1++
		} else {
			temparr[index] = arr[temp2]
			index++
			temp2++
		}
	}
	//对应第三步，将某一小集合的剩余元素存到大集合中
	if temp1 <= mid {
		temparr = append(temparr[index:index+mid-temp1+1], arr[temp1:temp1+mid-temp1+1]...)
		// temparr[index : index+mid-temp1+1] = arr[temp1 : temp1+mid-temp1+1]
	}
	if temp2 <= right {
		temparr = append(temparr[index:index+right-temp2+1], arr[temp2:temp2+right-temp2+1]...)
		// temparr[index : index+right-temp2+1] = arr[temp2 : temp2+right-temp2+1]
	}
	//将大集合的元素复制回原数组
	arr = append(arr[left:left+right-left+1], temparr[0:right-left+1]...)
}

/*7.计数排序*/
func CountSort(nums []int) []int {
	n := len(nums)
	if n < 1 {
		return nums
	}
	//求出最大最小值
	max, min := nums[0], nums[0]
	for _, x := range nums {
		if max < x {
			max = x
		}
		if min > x {
			min = x
		}
	}
	//设置 presum 数组长度,然后求出我们的前缀和数组，
	//这里我们可以把求次数数组和前缀和数组用一个数组处理
	presum := make([]int, max-min+1)
	for _, x := range nums {
		presum[x-min] += 1
	}
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + presum[i]
	}
	//临时数组
	temp := make([]int, n)
	//遍历数组，开始排序,注意偏移量
	for i := n - 1; i >= 0; i-- {
		//查找 presum 字典，然后将其放到临时数组，注意偏移度
		index := presum[nums[i]-min] - 1
		temp[index] = nums[i]
		//相应位置减一
		presum[nums[i]-min] -= 1
	}
	//copy回原数组
	copy(nums, temp)
	return nums
}

/*8.堆排序*/
