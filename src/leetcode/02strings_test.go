package leetcode

import (
	"reflect"
	"testing"
)

//01.颜色分类
func TestSortColors(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}
	tests := []test{
		{input: []int{2, 0, 2, 1, 0, 1}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{0, 0, 2, 1, 1, 1}, want: []int{0, 0, 1, 1, 1, 2}},
		{input: []int{2, 0, 1, 1, 0, 2}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{0, 0, 2, 1, 1, 2}, want: []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tc := range tests {
		got := SortColors(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("期望：%v 结果：%v", tc.want, got)
		}
	}
	// got := SortColors([]int{2, 0, 2, 1, 0, 1})
	// want := []int{0, 0, 1, 1, 2, 2}
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("期望：%v 结果：%v", want,got)
	// }
}
func TestSortColors_B(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}
	tests := []test{
		{input: []int{2, 0, 2, 1, 0, 1}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{0, 0, 2, 1, 1, 1}, want: []int{0, 0, 1, 1, 1, 2}},
		{input: []int{2, 0, 1, 1, 0, 2}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{0, 0, 2, 1, 1, 2}, want: []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tc := range tests {
		got := SortColors_B(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("期望：%v 结果：%v", tc.want, got)
		}
	}
	// got := SortColors([]int{2, 0, 2, 1, 0, 1})
	// want := []int{0, 0, 1, 1, 2, 2}
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("期望：%v 结果：%v", want,got)
	// }
}
