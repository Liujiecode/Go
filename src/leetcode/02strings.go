package leetcode

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
