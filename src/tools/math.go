package tools

import "fmt"

func Abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}

}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

/*[]int 一维整形数组去重*/
func Int_Rm_duplicate(list []int) []int {
	x := make([]int, 0)
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	fmt.Println("去重数组后：", x)
	return x
}

/*[]int 一维整形数组去重*/
func Str_Rm_duplicate(list []string) []string {
	x := make([]string, 0)
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	fmt.Println("去重数组后：", x)
	return x
}
