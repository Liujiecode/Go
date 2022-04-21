package main

import "fmt"

func test() {
	// var n int
	// fmt.Scan(&n)
	// hash := make(map[int]string, n)
	// h := make([]int, n)
	// s := make([]string, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&h[i])
	// }
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&s[i])
	// }
	// for i := 0; i < n; i++ {
	// 	hash[h[i]] = s[i]
	// }
	// fmt.Println(n, h, s)
	n := 7
	h := []int{3, 2, 2, 2, 1, 2, 1}
	s := []string{"b", "cd", "bdc", "bd", "a", "a", "aa"}
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n-i-1; j++ {
			if h[j] > h[j+1] {
				h[j], h[j+1] = h[j+1], h[j]
				s[j], s[j+1] = s[j+1], s[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	for c := 0; c < n-1; c++ {

		for i := 1; i < n; i++ {
			// current := s[i]
			j := i - 1
			for ; j >= 0 && h[i] == h[j]; j-- {
				k := 0
				for ; k < len(s[i]) && k < len(s[j]); k++ {
					if s[j][k] > s[i][k] {
						s[j], s[i] = s[i], s[j]
						break

					}
				}
				// if k == len(s[i]) && k < len(s[j]) {
				// 	s[j], s[i] = s[i], s[j]
				// }
				// s[j+1] = current
			}

		}
	}
	fmt.Println(s)
}
