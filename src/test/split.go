package test

import (
	"errors"
	"strings"
)

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep) //在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	// result = make([]string, 0, strings.Count(s, sep)+1)
	//提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。
	//使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}
